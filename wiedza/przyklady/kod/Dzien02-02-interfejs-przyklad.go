package main

import "fmt"

// Tworzenie interfejsu Speaker z zadeklarowaniem funkcji Speak()
type Speaker interface {
	Speak() string
}

// Human - struktura reprezentuja człowieka z podanym imieniem
type Human struct {
	Name string
}

// Implementacja interfejsu Speaker (funkcji Speak()) dla struktury Human
func (h Human) Speak() string {
	return fmt.Sprintf("Siemka, jestem %s", h.Name)
}

// Dog - struktura reprezentuja psa z podanym imieniem
type Dog struct {
	Name string
}

// Implementacja interfejsu Speaker (funkcji Speak()) dla struktury Dog
func (d Dog) Speak() string {
	return fmt.Sprintf("%s szczeka: hau! hau!", d.Name)
}

func main() {
	// przykładowe obiekty/struktury
	human := Human{Name: "Jan"}
	dog := Dog{Name: "Reksio"}

	speakers := []Speaker{human, dog}
	// iterowanie po elementach
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}

}
