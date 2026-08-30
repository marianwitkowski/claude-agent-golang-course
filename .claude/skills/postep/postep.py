#!/usr/bin/env python3
"""Atomowe operacje na postep/student.json — helper dla agenta go-tutor.

Każda operacja modyfikująca:
  1. Wczytuje student.json (lub tworzy domyślny przy `init`).
  2. Migruje schemat (v0 → v1 → v2) jeśli stary.
  3. Backupuje obecny plik do postep/backups/student.{ISO}.json.
  4. Modyfikuje obiekt w pamięci.
  5. Zapisuje do postep/student.json.tmp.
  6. Waliduje (re-parse JSON).
  7. Atomowy mv .tmp → student.json (Path.replace).

Komendy:
  init --imie X --cel Y --tempo Z [--system S --go-cmd P --go-version V]
  read [--field <dotted.path>]
  set --field <name> --value <val>
  add-lekcja --id X.Y --trudnosc 1-5
  add-cwiczenie --lekcja X.Y --poziom warmup|main|star
  add-mocna-strona "tekst"
  add-do-powtorki --temat T --lekcja X.Y
  remove-do-powtorki --temat T
  update-srodowisko [--system S] [--go-cmd P] [--go-version V] [--shell SH] [--edytor E]
  add-notatka "tekst"
  end-session
  recovery
"""

from __future__ import annotations

import argparse
import json
import shutil
import sys
from datetime import date, datetime
from pathlib import Path

SCHEMA_VERSION = 2

DEFAULT_STUDENT = {
    "schema_version": SCHEMA_VERSION,
    "imie": "",
    "cel": "",
    "tempo_godz_tydz": "",
    "rozpoczeto": "",
    "ostatnia_sesja": "",
    "liczba_sesji": 1,
    "aktualna_lekcja": "1.1",
    "srodowisko": {
        "system": "",
        "go_cmd": "",
        "go_version": "",
        "shell": "",
        "edytor": "",
    },
    "ukonczone_lekcje": [],
    "ukonczone_cwiczenia": [],
    "mocne_strony": [],
    "do_powtorki": [],
    "notatki_tutora": [],
}


def find_root() -> Path:
    """Znajdź katalog główny projektu (zawiera postep/ i wiedza/)."""
    here = Path(__file__).resolve()
    for parent in [*here.parents, Path.cwd()]:
        if (parent / "postep").is_dir() and (parent / "wiedza").is_dir():
            return parent
    raise SystemExit("BŁĄD: nie znalazłem katalogu głównego projektu (oczekiwane: postep/ + wiedza/)")


ROOT = find_root()
STUDENT = ROOT / "postep" / "student.json"
BACKUPS = ROOT / "postep" / "backups"
TMP = ROOT / "postep" / "student.json.tmp"


def today() -> str:
    return date.today().isoformat()


def ts() -> str:
    # Mikrosekundy zapewniają unikalność backupów nawet przy szybkich sekwencjach
    # (np. add-lekcja + set + end-session w jednej sekundzie).
    return datetime.now().strftime("%Y-%m-%d-%H-%M-%S-%f")


def read_state() -> dict:
    if not STUDENT.exists():
        raise SystemExit(f"BŁĄD: {STUDENT} nie istnieje. Najpierw `init`.")
    try:
        return json.loads(STUDENT.read_text(encoding="utf-8"))
    except json.JSONDecodeError as e:
        raise SystemExit(
            f"BŁĄD: {STUDENT} nie jest poprawnym JSON-em ({e}). "
            f"Uruchom `postep.py recovery` aby przywrócić z backupu."
        )


def migrate(state: dict) -> dict:
    """Aktualizuj stary schemat do bieżącego (SCHEMA_VERSION)."""
    sv = state.get("schema_version", 0)
    if sv > SCHEMA_VERSION:
        raise SystemExit(
            f"BŁĄD: schema_version={sv} jest nowsza niż obsługiwana ({SCHEMA_VERSION}). "
            f"Zaktualizuj postep.py."
        )
    if sv < 1:
        state["schema_version"] = 1
    if sv < 2:
        state.setdefault("srodowisko", {
            "system": "", "go_cmd": "", "go_version": "", "shell": "", "edytor": "",
        })
        state["schema_version"] = 2
    return state


def backup(state_path: Path = STUDENT) -> Path | None:
    """Zapisz kopię obecnego pliku do postep/backups/. Zwraca ścieżkę backupu."""
    if not state_path.exists():
        return None
    BACKUPS.mkdir(parents=True, exist_ok=True)
    dest = BACKUPS / f"student.{ts()}.json"
    shutil.copy2(state_path, dest)
    return dest


def write_atomic(state: dict) -> None:
    """Zapisz state do student.json atomowo (write tmp → validate → replace)."""
    payload = json.dumps(state, indent=2, ensure_ascii=False) + "\n"
    TMP.write_text(payload, encoding="utf-8")
    # Re-parse jako walidacja
    json.loads(TMP.read_text(encoding="utf-8"))
    TMP.replace(STUDENT)


def save(state: dict) -> None:
    """Pełny zapis: backup + atomowy write."""
    backup()
    write_atomic(state)


def _get_path(state: dict, dotted: str):
    """Wczytaj wartość spod ścieżki kropkowej (np. 'srodowisko.system')."""
    obj = state
    for part in dotted.split("."):
        obj = obj[part]
    return obj


def _set_path(state: dict, dotted: str, value) -> None:
    """Ustaw wartość pod ścieżką kropkową."""
    parts = dotted.split(".")
    parent = state
    for p in parts[:-1]:
        parent = parent[p]
    parent[parts[-1]] = value


# ===== Komendy =====

def cmd_init(args):
    if STUDENT.exists():
        raise SystemExit(
            f"BŁĄD: {STUDENT} już istnieje. Aby zmienić pola, użyj `set` / `update-srodowisko`. "
            f"Aby zacząć od nowa, użyj skill `reset-kursu`."
        )
    state = json.loads(json.dumps(DEFAULT_STUDENT))  # deep copy
    state["imie"] = args.imie
    state["cel"] = args.cel
    state["tempo_godz_tydz"] = args.tempo
    state["rozpoczeto"] = today()
    state["ostatnia_sesja"] = today()
    # Pełen snapshot środowiska od razu (nie wymaga osobnego update-srodowisko)
    if args.system:
        state["srodowisko"]["system"] = args.system
    if args.go_cmd:
        state["srodowisko"]["go_cmd"] = args.go_cmd
    if args.go_version:
        state["srodowisko"]["go_version"] = args.go_version
    if args.shell:
        state["srodowisko"]["shell"] = args.shell
    if args.edytor:
        state["srodowisko"]["edytor"] = args.edytor
    # init: nie ma czego backupować, ale i tak atomowy write
    write_atomic(state)
    print(f"OK: utworzono {STUDENT.relative_to(ROOT)} (schema v{SCHEMA_VERSION})")


def cmd_read(args):
    state = read_state()
    if args.field:
        try:
            value = _get_path(state, args.field)
        except KeyError as e:
            raise SystemExit(f"BŁĄD: pole '{args.field}' nie istnieje ({e})")
        print(json.dumps(value, ensure_ascii=False, indent=2))
    else:
        print(json.dumps(state, ensure_ascii=False, indent=2))


def cmd_set(args):
    state = migrate(read_state())
    try:
        _set_path(state, args.field, args.value)
    except KeyError as e:
        raise SystemExit(f"BŁĄD: pole '{args.field}' nie istnieje ({e})")
    save(state)
    print(f"OK: {args.field} = {args.value!r}")


def cmd_add_lekcja(args):
    state = migrate(read_state())
    if not (1 <= args.trudnosc <= 5):
        raise SystemExit("BŁĄD: trudnosc musi być 1-5")

    # Upsert po id — chroni przed duplikatami przy retry/timeout.
    # Jeśli uczeń powtarza lekcję, trudność subiektywna może się zmienić — bierzemy najświeższą.
    for entry in state["ukonczone_lekcje"]:
        if entry.get("id") == args.id:
            entry["data"] = today()
            entry["trudnosc_subiektywna"] = args.trudnosc
            state["ostatnia_sesja"] = today()
            save(state)
            print(f"OK: zaktualizowano lekcję {args.id} (trudność {args.trudnosc})")
            return

    state["ukonczone_lekcje"].append({
        "id": args.id,
        "data": today(),
        "trudnosc_subiektywna": args.trudnosc,
    })
    state["ostatnia_sesja"] = today()
    save(state)
    print(f"OK: dopisano lekcję {args.id} (trudność {args.trudnosc})")


def cmd_add_cwiczenie(args):
    state = migrate(read_state())

    # Upsert po (lekcja, poziom) — idempotentny przy retry.
    for entry in state["ukonczone_cwiczenia"]:
        if entry.get("lekcja") == args.lekcja and entry.get("poziom") == args.poziom:
            entry["data"] = today()
            save(state)
            print(f"OK: zaktualizowano ćwiczenie {args.lekcja}/{args.poziom}")
            return

    state["ukonczone_cwiczenia"].append({
        "lekcja": args.lekcja,
        "poziom": args.poziom,
        "data": today(),
    })
    save(state)
    print(f"OK: dopisano ćwiczenie {args.lekcja}/{args.poziom}")


def cmd_add_mocna_strona(args):
    state = migrate(read_state())
    if args.text in state["mocne_strony"]:
        print(f"INFO: '{args.text}' już jest na liście — pomijam")
        return
    state["mocne_strony"].append(args.text)
    # Trzymaj max 7 najnowszych
    state["mocne_strony"] = state["mocne_strony"][-7:]
    save(state)
    print(f"OK: dopisano mocną stronę: {args.text!r}")


def cmd_add_do_powtorki(args):
    state = migrate(read_state())
    # Nie dubluj jeśli ten sam temat już jest
    for d in state["do_powtorki"]:
        if d.get("temat") == args.temat:
            print(f"INFO: temat '{args.temat}' już w do_powtorki — pomijam")
            return
    state["do_powtorki"].append({
        "temat": args.temat,
        "lekcja": args.lekcja,
        "data_zauwazenia": today(),
    })
    save(state)
    print(f"OK: dopisano do_powtorki: {args.temat}")


def cmd_remove_do_powtorki(args):
    state = migrate(read_state())
    before = len(state["do_powtorki"])
    state["do_powtorki"] = [d for d in state["do_powtorki"] if d.get("temat") != args.temat]
    removed = before - len(state["do_powtorki"])
    save(state)
    print(f"OK: usunięto {removed} wpisów do_powtorki o temacie {args.temat!r}")


def cmd_update_srodowisko(args):
    state = migrate(read_state())
    updated = {}
    for field in ("system", "go_cmd", "go_version", "shell", "edytor"):
        v = getattr(args, field)
        if v is not None:
            state["srodowisko"][field] = v
            updated[field] = v
    if not updated:
        raise SystemExit("BŁĄD: nic do zaktualizowania — podaj co najmniej jedno pole")
    save(state)
    print(f"OK: zaktualizowano środowisko: {updated}")


def cmd_add_notatka(args):
    state = migrate(read_state())
    state["notatki_tutora"].append(args.text)
    # Trzymaj max 20 najnowszych
    state["notatki_tutora"] = state["notatki_tutora"][-20:]
    save(state)
    print(f"OK: dopisano notatkę tutora")


def cmd_end_session(args):
    state = migrate(read_state())
    state["ostatnia_sesja"] = today()
    state["liczba_sesji"] = state.get("liczba_sesji", 0) + 1
    save(state)
    print(f"OK: zakończono sesję #{state['liczba_sesji']}")


def cmd_recovery(args):
    """Przywróć student.json z najnowszego działającego backupu."""
    BACKUPS.mkdir(parents=True, exist_ok=True)
    backups = sorted(BACKUPS.glob("student.*.json"), reverse=True)
    if not backups:
        raise SystemExit("BŁĄD: brak backupów do przywrócenia (postep/backups/ jest puste)")

    for b in backups:
        try:
            data = json.loads(b.read_text(encoding="utf-8"))
        except json.JSONDecodeError:
            continue
        # Mamy działający backup
        if STUDENT.exists():
            broken = STUDENT.with_name(f"student.broken.{ts()}.json")
            STUDENT.rename(broken)
            print(f"INFO: stary plik przeniesiony do {broken.name}")
        shutil.copy2(b, STUDENT)
        print(f"OK: przywrócono z {b.name}")
        print(f"     imię: {data.get('imie', '?')}")
        print(f"     aktualna lekcja: {data.get('aktualna_lekcja', '?')}")
        print(f"     ukończonych lekcji: {len(data.get('ukonczone_lekcje', []))}")
        return

    raise SystemExit("BŁĄD: żaden z backupów nie parsuje się jako JSON")


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Atomic update of postep/student.json",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog=__doc__,
    )
    sub = parser.add_subparsers(dest="cmd", required=True)

    p = sub.add_parser("init", help="Stwórz nowy student.json")
    p.add_argument("--imie", required=True)
    p.add_argument("--cel", required=True)
    p.add_argument("--tempo", required=True, help="np. '<2', '2-5', '5-10', '10+'")
    p.add_argument("--system", help="macOS | Linux | Windows")
    p.add_argument("--go-cmd", dest="go_cmd", help="np. go (zwykle po prostu \"go\")")
    p.add_argument("--go-version", dest="go_version", help="wersja Go, np. 1.25.0")
    p.add_argument("--shell", help="np. zsh, bash, PowerShell")
    p.add_argument("--edytor", help="np. VS Code, PyCharm, Sublime")

    p = sub.add_parser("read", help="Pokaż stan (cały lub jedno pole)")
    p.add_argument("--field", help="ścieżka kropkowa, np. srodowisko.system")

    p = sub.add_parser("set", help="Ustaw wartość pojedynczego pola")
    p.add_argument("--field", required=True, help="np. aktualna_lekcja lub srodowisko.go_cmd")
    p.add_argument("--value", required=True)

    p = sub.add_parser("add-lekcja", help="Dopisz ukończoną lekcję")
    p.add_argument("--id", required=True, help="np. 4.1")
    p.add_argument("--trudnosc", type=int, required=True, choices=[1, 2, 3, 4, 5])

    p = sub.add_parser("add-cwiczenie", help="Dopisz ukończone ćwiczenie")
    p.add_argument("--lekcja", required=True)
    p.add_argument("--poziom", required=True, choices=["warmup", "main", "star"])

    p = sub.add_parser("add-mocna-strona", help="Dopisz mocną stronę ucznia")
    p.add_argument("text")

    p = sub.add_parser("add-do-powtorki", help="Dopisz temat do powtórki")
    p.add_argument("--temat", required=True)
    p.add_argument("--lekcja", required=True)

    p = sub.add_parser("remove-do-powtorki", help="Usuń temat z listy do powtórki (po zaliczeniu)")
    p.add_argument("--temat", required=True)

    p = sub.add_parser("update-srodowisko", help="Zaktualizuj pola środowiska (OS, komenda Go itp.)")
    p.add_argument("--system")
    p.add_argument("--go-cmd", dest="go_cmd")
    p.add_argument("--go-version", dest="go_version")
    p.add_argument("--shell")
    p.add_argument("--edytor")

    p = sub.add_parser("add-notatka", help="Dopisz notatkę tutora (prywatna)")
    p.add_argument("text")

    p = sub.add_parser("end-session", help="Zamknij sesję (ostatnia_sesja=dziś, liczba_sesji+=1)")

    p = sub.add_parser("recovery", help="Przywróć student.json z najnowszego działającego backupu")

    return parser


def main():
    parser = build_parser()
    args = parser.parse_args()
    cmd_name = args.cmd.replace("-", "_")
    cmd_func = globals()[f"cmd_{cmd_name}"]
    cmd_func(args)


if __name__ == "__main__":
    main()
