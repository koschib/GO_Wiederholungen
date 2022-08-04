package Aufgabe2

import "fmt"

func Funktionen() {

	fmt.Println("Ich bin Aufgabe 2.5: Funktionen")
	fmt.Println("-----------------------------------------------------")
	FunktionOhneRueckgabewert()

	// Funktionmitrueckgabewert + 322 rechnen
	fmt.Println("Funktion mit Rückgabewert:", FunktionMitRueckgabewert()+322)
	fmt.Println("-----------------------------------------------------")
	Theorie()

}

// Funktion mit Rückgabewert

func FunktionMitRueckgabewert() int {
	return 42
}

// funktion ohne Rückgabewert

func FunktionOhneRueckgabewert() {
	fmt.Println("Ich bin eine Funktion ohne Rückgabewert")
}

func Theorie() {

	fmt.Println("Funktionen mit Rückgabewert werden genutzt um Variablen zu speichern")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Funktionen ohne Rückgabewert werden genutzt um Code zu verketten")

}
