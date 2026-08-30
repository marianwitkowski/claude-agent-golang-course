// Narzędzie postep — atomowe operacje na postep/student.json.
//
// Helper agenta go-tutor. Uczeń go nie uruchamia.
//
// Każda operacja modyfikująca:
//  1. Wczytuje student.json (lub tworzy domyślny przy `init`).
//  2. Migruje schemat (v0 → v1 → v2), jeśli stary.
//  3. Backupuje obecny plik do postep/backups/student.{TS}.json.
//  4. Modyfikuje strukturę w pamięci.
//  5. Zapisuje do postep/student.json.tmp.
//  6. Waliduje (ponowne parsowanie).
//  7. Atomowy os.Rename .tmp → student.json.
//
// Komendy:
//
//	init --imie X --cel Y --tempo Z [--system S --go-cmd C --go-version V --shell SH --edytor E]
//	read [--field <sciezka.kropkowa>]
//	set --field <sciezka> --value <wartosc>
//	add-lekcja --id X.Y --trudnosc 1-5
//	add-cwiczenie --lekcja X.Y --poziom warmup|main|star
//	add-mocna-strona "tekst"
//	add-do-powtorki --temat T --lekcja X.Y
//	remove-do-powtorki --temat T
//	update-srodowisko [--system S] [--go-cmd C] [--go-version V] [--shell SH] [--edytor E]
//	add-notatka "tekst"
//	end-session
//	recovery
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const wersjaSchematu = 2

const (
	maxMocnychStron = 7
	maxNotatek      = 20
)

// ===== Model stanu =====

type Srodowisko struct {
	System    string `json:"system"`
	GoCmd     string `json:"go_cmd"`
	GoVersion string `json:"go_version"`
	Shell     string `json:"shell"`
	Edytor    string `json:"edytor"`
}

type Lekcja struct {
	ID       string `json:"id"`
	Data     string `json:"data"`
	Trudnosc int    `json:"trudnosc_subiektywna"`
}

type Cwiczenie struct {
	Lekcja string `json:"lekcja"`
	Poziom string `json:"poziom"`
	Data   string `json:"data"`
}

type Powtorka struct {
	Temat          string `json:"temat"`
	Lekcja         string `json:"lekcja"`
	DataZauwazenia string `json:"data_zauwazenia"`
}

type Stan struct {
	WersjaSchematu     int         `json:"schema_version"`
	Imie               string      `json:"imie"`
	Cel                string      `json:"cel"`
	TempoGodzTydz      string      `json:"tempo_godz_tydz"`
	Rozpoczeto         string      `json:"rozpoczeto"`
	OstatniaSesja      string      `json:"ostatnia_sesja"`
	LiczbaSesji        int         `json:"liczba_sesji"`
	AktualnaLekcja     string      `json:"aktualna_lekcja"`
	Srodowisko         Srodowisko  `json:"srodowisko"`
	UkonczoneLekcje    []Lekcja    `json:"ukonczone_lekcje"`
	UkonczoneCwiczenia []Cwiczenie `json:"ukonczone_cwiczenia"`
	MocneStrony        []string    `json:"mocne_strony"`
	DoPowtorki         []Powtorka  `json:"do_powtorki"`
	NotatkiTutora      []string    `json:"notatki_tutora"`

	// Pola zapisane przez nowszą wersję narzędzia, której ten plik nie zna.
	// Przechowywane w surowej postaci i zapisywane z powrotem — bez tego
	// starsze narzędzie kasowałoby stan, którego nie rozumie.
	nieznane map[string]json.RawMessage
}

// znanePola muszą odpowiadać tagom json powyżej — używane do odsiania
// tego, co trafia do `nieznane`.
var znanePola = []string{
	"schema_version", "imie", "cel", "tempo_godz_tydz", "rozpoczeto",
	"ostatnia_sesja", "liczba_sesji", "aktualna_lekcja", "srodowisko",
	"ukonczone_lekcje", "ukonczone_cwiczenia", "mocne_strony",
	"do_powtorki", "notatki_tutora",
}

func nowyStan() Stan {
	return Stan{
		WersjaSchematu:     wersjaSchematu,
		LiczbaSesji:        1,
		AktualnaLekcja:     "1.1",
		UkonczoneLekcje:    []Lekcja{},
		UkonczoneCwiczenia: []Cwiczenie{},
		MocneStrony:        []string{},
		DoPowtorki:         []Powtorka{},
		NotatkiTutora:      []string{},
	}
}

// pustePisteZamiastNil pilnuje, żeby puste listy zapisywały się jako [],
// a nie null — plik ma wyglądać tak samo niezależnie od historii zmian.
func (s *Stan) pusteListyZamiastNil() {
	if s.UkonczoneLekcje == nil {
		s.UkonczoneLekcje = []Lekcja{}
	}
	if s.UkonczoneCwiczenia == nil {
		s.UkonczoneCwiczenia = []Cwiczenie{}
	}
	if s.MocneStrony == nil {
		s.MocneStrony = []string{}
	}
	if s.DoPowtorki == nil {
		s.DoPowtorki = []Powtorka{}
	}
	if s.NotatkiTutora == nil {
		s.NotatkiTutora = []string{}
	}
}

func (s *Stan) UnmarshalJSON(dane []byte) error {
	type alias Stan // alias bez metod — inaczej rekurencja
	var a alias
	if err := json.Unmarshal(dane, &a); err != nil {
		return err
	}
	*s = Stan(a)

	var wszystkie map[string]json.RawMessage
	if err := json.Unmarshal(dane, &wszystkie); err != nil {
		return err
	}
	for _, k := range znanePola {
		delete(wszystkie, k)
	}
	if len(wszystkie) > 0 {
		s.nieznane = wszystkie
	}
	return nil
}

func (s Stan) MarshalJSON() ([]byte, error) {
	type alias Stan
	znane, err := marshalBezHTML(alias(s))
	if err != nil {
		return nil, err
	}
	if len(s.nieznane) == 0 {
		return znane, nil
	}

	klucze := make([]string, 0, len(s.nieznane))
	for k := range s.nieznane {
		klucze = append(klucze, k)
	}
	sort.Strings(klucze)

	var buf bytes.Buffer
	buf.Write(znane[:len(znane)-1]) // wszystko bez zamykającego }
	for _, k := range klucze {
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		kj, err := marshalBezHTML(k)
		if err != nil {
			return nil, err
		}
		buf.Write(kj)
		buf.WriteByte(':')
		buf.Write(s.nieznane[k])
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// marshalBezHTML serializuje bez zamiany < > & na sekwencje \u — plik ma
// być czytelny dla człowieka, a nie wklejany do HTML-a.
func marshalBezHTML(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}

// ===== Ścieżki =====

type Sciezki struct {
	Katalog string // katalog główny projektu
	Student string
	Backupy string
	Tmp     string
}

func sciezkiDla(katalog string) Sciezki {
	return Sciezki{
		Katalog: katalog,
		Student: filepath.Join(katalog, "postep", "student.json"),
		Backupy: filepath.Join(katalog, "postep", "backups"),
		Tmp:     filepath.Join(katalog, "postep", "student.json.tmp"),
	}
}

// znajdzKatalogGlowny szuka w górę katalogu zawierającego postep/ i wiedza/.
func znajdzKatalogGlowny() (string, error) {
	kat, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if jestKatalogiem(filepath.Join(kat, "postep")) && jestKatalogiem(filepath.Join(kat, "wiedza")) {
			return kat, nil
		}
		rodzic := filepath.Dir(kat)
		if rodzic == kat {
			return "", errors.New("nie znalazłem katalogu głównego projektu (oczekiwane: postep/ + wiedza/); podaj -root")
		}
		kat = rodzic
	}
}

func jestKatalogiem(sciezka string) bool {
	info, err := os.Stat(sciezka)
	return err == nil && info.IsDir()
}

// ===== Odczyt, zapis, backup =====

func dzisiaj() string {
	return time.Now().Format("2006-01-02")
}

// znacznikCzasu ma mikrosekundy, żeby backupy były unikalne nawet przy
// szybkich sekwencjach (add-lekcja + set + end-session w tej samej sekundzie).
func znacznikCzasu() string {
	t := time.Now()
	return fmt.Sprintf("%s-%06d", t.Format("2006-01-02-15-04-05"), t.Nanosecond()/1000)
}

func wczytajSurowo(sc Sciezki) ([]byte, error) {
	dane, err := os.ReadFile(sc.Student)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%s nie istnieje. Najpierw `init`", sc.Student)
	}
	if err != nil {
		return nil, err
	}
	return dane, nil
}

func wczytajStan(sc Sciezki) (Stan, error) {
	dane, err := wczytajSurowo(sc)
	if err != nil {
		return Stan{}, err
	}
	var s Stan
	if err := json.Unmarshal(dane, &s); err != nil {
		return Stan{}, fmt.Errorf("%s nie jest poprawnym JSON-em (%v). Uruchom `recovery`, aby przywrócić z backupu", sc.Student, err)
	}
	return s, nil
}

// migruj podnosi stary schemat do bieżącego. Brakujące `srodowisko`
// wypełnia się samo wartościami zerowymi przy parsowaniu.
func migruj(s *Stan) error {
	if s.WersjaSchematu > wersjaSchematu {
		return fmt.Errorf("schema_version=%d jest nowsza niż obsługiwana (%d). Zaktualizuj narzędzie postep", s.WersjaSchematu, wersjaSchematu)
	}
	s.WersjaSchematu = wersjaSchematu
	return nil
}

func backup(sc Sciezki) error {
	dane, err := os.ReadFile(sc.Student)
	if errors.Is(err, os.ErrNotExist) {
		return nil // nie ma czego backupować
	}
	if err != nil {
		return err
	}
	if err := os.MkdirAll(sc.Backupy, 0o755); err != nil {
		return err
	}
	cel := filepath.Join(sc.Backupy, fmt.Sprintf("student.%s.json", znacznikCzasu()))
	return os.WriteFile(cel, dane, 0o644)
}

func zapiszAtomowo(sc Sciezki, s Stan) error {
	s.pusteListyZamiastNil()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(s); err != nil {
		return err
	}
	if err := os.WriteFile(sc.Tmp, buf.Bytes(), 0o644); err != nil {
		return err
	}
	// Walidacja: plik na dysku musi się parsować, zanim podmienimy oryginał.
	kontrola, err := os.ReadFile(sc.Tmp)
	if err != nil {
		return err
	}
	var cokolwiek any
	if err := json.Unmarshal(kontrola, &cokolwiek); err != nil {
		return fmt.Errorf("zapisany plik nie parsuje się (%v) — oryginał nietknięty", err)
	}
	return os.Rename(sc.Tmp, sc.Student)
}

// zapisz = backup + atomowy zapis.
func zapisz(sc Sciezki, s Stan) error {
	if err := backup(sc); err != nil {
		return err
	}
	return zapiszAtomowo(sc, s)
}

// ===== Ścieżki kropkowe (read/set) =====

func pobierzSciezke(dane map[string]any, kropkowa string) (any, error) {
	var biezacy any = dane
	for _, czesc := range strings.Split(kropkowa, ".") {
		mapa, ok := biezacy.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("pole '%s' nie istnieje (fragment '%s' nie jest obiektem)", kropkowa, czesc)
		}
		wartosc, ok := mapa[czesc]
		if !ok {
			return nil, fmt.Errorf("pole '%s' nie istnieje (brak klucza '%s')", kropkowa, czesc)
		}
		biezacy = wartosc
	}
	return biezacy, nil
}

func ustawSciezke(dane map[string]any, kropkowa string, wartosc string) error {
	czesci := strings.Split(kropkowa, ".")
	rodzic := dane
	for _, czesc := range czesci[:len(czesci)-1] {
		nastepny, ok := rodzic[czesc].(map[string]any)
		if !ok {
			return fmt.Errorf("pole '%s' nie istnieje (fragment '%s' nie jest obiektem)", kropkowa, czesc)
		}
		rodzic = nastepny
	}
	ostatni := czesci[len(czesci)-1]
	if _, ok := rodzic[ostatni]; !ok {
		return fmt.Errorf("pole '%s' nie istnieje (brak klucza '%s')", kropkowa, ostatni)
	}
	rodzic[ostatni] = wartosc
	return nil
}

// ===== Komendy =====

func cmdInit(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("init", flag.ExitOnError)
	imie := fs.String("imie", "", "imię ucznia (wymagane)")
	cel := fs.String("cel", "", "praca | narzędzia | hobby | szkoła (wymagane)")
	tempo := fs.String("tempo", "", "np. '<2', '2-5', '5-10', '10+' (wymagane)")
	system := fs.String("system", "", "macOS | Linux | Windows")
	goCmd := fs.String("go-cmd", "", "komenda Go (zwykle po prostu \"go\")")
	goVersion := fs.String("go-version", "", "wersja Go bez prefiksu, np. 1.25.3")
	shell := fs.String("shell", "", "np. zsh, bash, PowerShell")
	edytor := fs.String("edytor", "", "np. VS Code, GoLand")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	for nazwa, wartosc := range map[string]string{"imie": *imie, "cel": *cel, "tempo": *tempo} {
		if wartosc == "" {
			return fmt.Errorf("brak wymaganego argumentu --%s", nazwa)
		}
	}
	if _, err := os.Stat(sc.Student); err == nil {
		return fmt.Errorf("%s już istnieje. Aby zmienić pola, użyj `set` / `update-srodowisko`; aby zacząć od nowa — skill `reset-kursu`", sc.Student)
	}

	s := nowyStan()
	s.Imie = *imie
	s.Cel = *cel
	s.TempoGodzTydz = *tempo
	s.Rozpoczeto = dzisiaj()
	s.OstatniaSesja = dzisiaj()
	s.Srodowisko = Srodowisko{
		System:    *system,
		GoCmd:     *goCmd,
		GoVersion: *goVersion,
		Shell:     *shell,
		Edytor:    *edytor,
	}

	if err := os.MkdirAll(filepath.Dir(sc.Student), 0o755); err != nil {
		return err
	}
	// init nie ma czego backupować, ale zapis i tak jest atomowy.
	if err := zapiszAtomowo(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: utworzono postep/student.json (schema v%d)\n", wersjaSchematu)
	return nil
}

func cmdRead(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("read", flag.ExitOnError)
	pole := fs.String("field", "", "ścieżka kropkowa, np. srodowisko.go_cmd")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}

	surowe, err := wczytajSurowo(sc)
	if err != nil {
		return err
	}
	var dane map[string]any
	if err := json.Unmarshal(surowe, &dane); err != nil {
		return fmt.Errorf("%s nie jest poprawnym JSON-em (%v). Uruchom `recovery`", sc.Student, err)
	}

	// Bez --field oddajemy plik bajt w bajt: zachowuje kolejność pól
	// ze schematu, której mapa by nie utrzymała (Go sortuje klucze map).
	if *pole == "" {
		os.Stdout.Write(surowe)
		return nil
	}

	wartosc, err := pobierzSciezke(dane, *pole)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(wartosc); err != nil {
		return err
	}
	fmt.Print(buf.String())
	return nil
}

func cmdSet(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("set", flag.ExitOnError)
	pole := fs.String("field", "", "np. aktualna_lekcja albo srodowisko.edytor (wymagane)")
	wartosc := fs.String("value", "", "nowa wartość (wymagane)")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	if *pole == "" {
		return errors.New("brak wymaganego argumentu --field")
	}

	surowe, err := wczytajSurowo(sc)
	if err != nil {
		return err
	}
	var dane map[string]any
	if err := json.Unmarshal(surowe, &dane); err != nil {
		return fmt.Errorf("%s nie jest poprawnym JSON-em (%v). Uruchom `recovery`", sc.Student, err)
	}
	if err := ustawSciezke(dane, *pole, *wartosc); err != nil {
		return err
	}

	// Powrót przez strukturę: przywraca kolejność pól i pilnuje typów.
	zmienione, err := marshalBezHTML(dane)
	if err != nil {
		return err
	}
	var s Stan
	if err := json.Unmarshal(zmienione, &s); err != nil {
		return fmt.Errorf("pole '%s' nie przyjmuje wartości tekstowej (%v) — `set` działa na polach tekstowych", *pole, err)
	}
	if err := migruj(&s); err != nil {
		return err
	}
	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: %s = %q\n", *pole, *wartosc)
	return nil
}

func cmdAddLekcja(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("add-lekcja", flag.ExitOnError)
	id := fs.String("id", "", "np. 4.1 (wymagane)")
	trudnosc := fs.Int("trudnosc", 0, "1-5 (wymagane)")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	if *id == "" {
		return errors.New("brak wymaganego argumentu --id")
	}
	if *trudnosc < 1 || *trudnosc > 5 {
		return errors.New("trudnosc musi być 1-5")
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}

	// Upsert po id — chroni przed duplikatami przy powtórzeniu komendy.
	// Gdy uczeń powtarza lekcję, bierzemy najświeższą ocenę trudności.
	for i := range s.UkonczoneLekcje {
		if s.UkonczoneLekcje[i].ID == *id {
			s.UkonczoneLekcje[i].Data = dzisiaj()
			s.UkonczoneLekcje[i].Trudnosc = *trudnosc
			s.OstatniaSesja = dzisiaj()
			if err := zapisz(sc, s); err != nil {
				return err
			}
			fmt.Printf("OK: zaktualizowano lekcję %s (trudność %d)\n", *id, *trudnosc)
			return nil
		}
	}

	s.UkonczoneLekcje = append(s.UkonczoneLekcje, Lekcja{
		ID:       *id,
		Data:     dzisiaj(),
		Trudnosc: *trudnosc,
	})
	s.OstatniaSesja = dzisiaj()
	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: dopisano lekcję %s (trudność %d)\n", *id, *trudnosc)
	return nil
}

func cmdAddCwiczenie(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("add-cwiczenie", flag.ExitOnError)
	lekcja := fs.String("lekcja", "", "np. 4.1 (wymagane)")
	poziom := fs.String("poziom", "", "warmup | main | star (wymagane)")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	if *lekcja == "" {
		return errors.New("brak wymaganego argumentu --lekcja")
	}
	switch *poziom {
	case "warmup", "main", "star":
	default:
		return fmt.Errorf("poziom musi być warmup, main albo star (dostałem %q)", *poziom)
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}

	// Upsert po (lekcja, poziom) — idempotentny przy powtórzeniu.
	for i := range s.UkonczoneCwiczenia {
		if s.UkonczoneCwiczenia[i].Lekcja == *lekcja && s.UkonczoneCwiczenia[i].Poziom == *poziom {
			s.UkonczoneCwiczenia[i].Data = dzisiaj()
			if err := zapisz(sc, s); err != nil {
				return err
			}
			fmt.Printf("OK: zaktualizowano ćwiczenie %s/%s\n", *lekcja, *poziom)
			return nil
		}
	}

	s.UkonczoneCwiczenia = append(s.UkonczoneCwiczenia, Cwiczenie{
		Lekcja: *lekcja,
		Poziom: *poziom,
		Data:   dzisiaj(),
	})
	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: dopisano ćwiczenie %s/%s\n", *lekcja, *poziom)
	return nil
}

func cmdAddMocnaStrona(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("add-mocna-strona", flag.ExitOnError)
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	tekst := strings.TrimSpace(strings.Join(fs.Args(), " "))
	if tekst == "" {
		return errors.New("podaj tekst mocnej strony jako argument")
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}
	for _, istniejaca := range s.MocneStrony {
		if istniejaca == tekst {
			fmt.Printf("INFO: %q już jest na liście — pomijam\n", tekst)
			return nil
		}
	}

	s.MocneStrony = append(s.MocneStrony, tekst)
	s.MocneStrony = ostatnie(s.MocneStrony, maxMocnychStron)
	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: dopisano mocną stronę: %q\n", tekst)
	return nil
}

func cmdAddDoPowtorki(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("add-do-powtorki", flag.ExitOnError)
	temat := fs.String("temat", "", "czego dotyczy luka (wymagane)")
	lekcja := fs.String("lekcja", "", "np. 7.2 (wymagane)")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	if *temat == "" || *lekcja == "" {
		return errors.New("wymagane argumenty: --temat i --lekcja")
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}
	for _, p := range s.DoPowtorki {
		if p.Temat == *temat {
			fmt.Printf("INFO: temat %q już w do_powtorki — pomijam\n", *temat)
			return nil
		}
	}

	s.DoPowtorki = append(s.DoPowtorki, Powtorka{
		Temat:          *temat,
		Lekcja:         *lekcja,
		DataZauwazenia: dzisiaj(),
	})
	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: dopisano do_powtorki: %s\n", *temat)
	return nil
}

func cmdRemoveDoPowtorki(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("remove-do-powtorki", flag.ExitOnError)
	temat := fs.String("temat", "", "temat do usunięcia (wymagane)")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	if *temat == "" {
		return errors.New("brak wymaganego argumentu --temat")
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}

	pozostale := make([]Powtorka, 0, len(s.DoPowtorki))
	for _, p := range s.DoPowtorki {
		if p.Temat != *temat {
			pozostale = append(pozostale, p)
		}
	}
	usuniete := len(s.DoPowtorki) - len(pozostale)
	s.DoPowtorki = pozostale

	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: usunięto %d wpisów do_powtorki o temacie %q\n", usuniete, *temat)
	return nil
}

func cmdUpdateSrodowisko(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("update-srodowisko", flag.ExitOnError)
	system := fs.String("system", "", "macOS | Linux | Windows")
	goCmd := fs.String("go-cmd", "", "komenda Go")
	goVersion := fs.String("go-version", "", "wersja Go bez prefiksu, np. 1.25.3")
	shell := fs.String("shell", "", "np. zsh, bash, PowerShell")
	edytor := fs.String("edytor", "", "np. VS Code, GoLand")
	if err := fs.Parse(argumenty); err != nil {
		return err
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}

	// Aktualizujemy tylko pola faktycznie podane w wierszu poleceń —
	// pusty łańcuch bez flagi nie może wyzerować istniejącej wartości.
	zmienione := map[string]string{}
	fs.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "system":
			s.Srodowisko.System = *system
			zmienione["system"] = *system
		case "go-cmd":
			s.Srodowisko.GoCmd = *goCmd
			zmienione["go_cmd"] = *goCmd
		case "go-version":
			s.Srodowisko.GoVersion = *goVersion
			zmienione["go_version"] = *goVersion
		case "shell":
			s.Srodowisko.Shell = *shell
			zmienione["shell"] = *shell
		case "edytor":
			s.Srodowisko.Edytor = *edytor
			zmienione["edytor"] = *edytor
		}
	})
	if len(zmienione) == 0 {
		return errors.New("nic do zaktualizowania — podaj co najmniej jedno pole")
	}

	if err := zapisz(sc, s); err != nil {
		return err
	}

	nazwy := make([]string, 0, len(zmienione))
	for k := range zmienione {
		nazwy = append(nazwy, k)
	}
	sort.Strings(nazwy)
	opisy := make([]string, 0, len(nazwy))
	for _, k := range nazwy {
		opisy = append(opisy, fmt.Sprintf("%s=%q", k, zmienione[k]))
	}
	fmt.Printf("OK: zaktualizowano środowisko: %s\n", strings.Join(opisy, ", "))
	return nil
}

func cmdAddNotatka(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("add-notatka", flag.ExitOnError)
	if err := fs.Parse(argumenty); err != nil {
		return err
	}
	tekst := strings.TrimSpace(strings.Join(fs.Args(), " "))
	if tekst == "" {
		return errors.New("podaj treść notatki jako argument")
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}
	s.NotatkiTutora = ostatnie(append(s.NotatkiTutora, tekst), maxNotatek)
	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Println("OK: dopisano notatkę tutora")
	return nil
}

func cmdEndSession(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("end-session", flag.ExitOnError)
	if err := fs.Parse(argumenty); err != nil {
		return err
	}

	s, err := wczytajStan(sc)
	if err != nil {
		return err
	}
	if err := migruj(&s); err != nil {
		return err
	}
	s.OstatniaSesja = dzisiaj()
	s.LiczbaSesji++
	if err := zapisz(sc, s); err != nil {
		return err
	}
	fmt.Printf("OK: zakończono sesję #%d\n", s.LiczbaSesji)
	return nil
}

func cmdRecovery(sc Sciezki, argumenty []string) error {
	fs := flag.NewFlagSet("recovery", flag.ExitOnError)
	if err := fs.Parse(argumenty); err != nil {
		return err
	}

	wpisy, err := os.ReadDir(sc.Backupy)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	var nazwy []string
	for _, w := range wpisy {
		n := w.Name()
		if !w.IsDir() && strings.HasPrefix(n, "student.") && strings.HasSuffix(n, ".json") {
			nazwy = append(nazwy, n)
		}
	}
	if len(nazwy) == 0 {
		return errors.New("brak backupów do przywrócenia (postep/backups/ jest puste)")
	}
	// Znacznik czasu ma stałą szerokość, więc sortowanie po nazwie
	// malejąco daje najnowszy backup na początku.
	sort.Sort(sort.Reverse(sort.StringSlice(nazwy)))

	for _, nazwa := range nazwy {
		sciezka := filepath.Join(sc.Backupy, nazwa)
		dane, err := os.ReadFile(sciezka)
		if err != nil {
			continue
		}
		var s Stan
		if err := json.Unmarshal(dane, &s); err != nil {
			continue // uszkodzony backup — próbujemy starszego
		}

		if _, err := os.Stat(sc.Student); err == nil {
			uszkodzony := filepath.Join(filepath.Dir(sc.Student), fmt.Sprintf("student.broken.%s.json", znacznikCzasu()))
			if err := os.Rename(sc.Student, uszkodzony); err != nil {
				return err
			}
			fmt.Printf("INFO: stary plik przeniesiony do %s\n", filepath.Base(uszkodzony))
		}
		if err := os.WriteFile(sc.Student, dane, 0o644); err != nil {
			return err
		}
		fmt.Printf("OK: przywrócono z %s\n", nazwa)
		fmt.Printf("     imię: %s\n", lubZnak(s.Imie))
		fmt.Printf("     aktualna lekcja: %s\n", lubZnak(s.AktualnaLekcja))
		fmt.Printf("     ukończonych lekcji: %d\n", len(s.UkonczoneLekcje))
		return nil
	}

	return errors.New("żaden z backupów nie parsuje się jako JSON")
}

// ===== Drobiazgi =====

func ostatnie[T any](s []T, n int) []T {
	if len(s) <= n {
		return s
	}
	return s[len(s)-n:]
}

func lubZnak(s string) string {
	if s == "" {
		return "?"
	}
	return s
}

// ===== Wejście =====

var komendy = map[string]func(Sciezki, []string) error{
	"init":               cmdInit,
	"read":               cmdRead,
	"set":                cmdSet,
	"add-lekcja":         cmdAddLekcja,
	"add-cwiczenie":      cmdAddCwiczenie,
	"add-mocna-strona":   cmdAddMocnaStrona,
	"add-do-powtorki":    cmdAddDoPowtorki,
	"remove-do-powtorki": cmdRemoveDoPowtorki,
	"update-srodowisko":  cmdUpdateSrodowisko,
	"add-notatka":        cmdAddNotatka,
	"end-session":        cmdEndSession,
	"recovery":           cmdRecovery,
}

func uzycie() {
	fmt.Fprintln(os.Stderr, "użycie: postep [-root <katalog>] <komenda> [argumenty]")
	fmt.Fprintln(os.Stderr, "\nkomendy:")
	nazwy := make([]string, 0, len(komendy))
	for k := range komendy {
		nazwy = append(nazwy, k)
	}
	sort.Strings(nazwy)
	for _, n := range nazwy {
		fmt.Fprintf(os.Stderr, "  %s\n", n)
	}
	fmt.Fprintln(os.Stderr, "\nszczegóły argumentów: postep <komenda> -h")
}

func main() {
	globalne := flag.NewFlagSet("postep", flag.ExitOnError)
	root := globalne.String("root", "", "katalog główny projektu (domyślnie: szukany w górę od bieżącego)")
	globalne.Usage = uzycie
	if err := globalne.Parse(os.Args[1:]); err != nil {
		os.Exit(2)
	}

	argumenty := globalne.Args()
	if len(argumenty) == 0 {
		uzycie()
		os.Exit(2)
	}

	nazwaKomendy := argumenty[0]
	komenda, ok := komendy[nazwaKomendy]
	if !ok {
		fmt.Fprintf(os.Stderr, "BŁĄD: nieznana komenda %q\n\n", nazwaKomendy)
		uzycie()
		os.Exit(2)
	}

	katalog := *root
	if katalog == "" {
		znaleziony, err := znajdzKatalogGlowny()
		if err != nil {
			fmt.Fprintf(os.Stderr, "BŁĄD: %v\n", err)
			os.Exit(1)
		}
		katalog = znaleziony
	}

	if err := komenda(sciezkiDla(katalog), argumenty[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "BŁĄD: %v\n", err)
		os.Exit(1)
	}
}
