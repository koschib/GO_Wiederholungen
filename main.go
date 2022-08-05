package main

import (
	"fmt"
	"wiederholungen/Aufgabe1"
	"wiederholungen/Aufgabe2"
	_ "wiederholungen/advanced"
	advanced2 "wiederholungen/advanced"
)

func main() {

	fmt.Println("Willkommen, das sind meine Lösungen:")
	fmt.Println("-----------------------------------------------------")

	// Auswahl Funktionen durch scan loesungen1 oder loesungen2 oder beenden
	fmt.Println("Wählen Sie eine Lösung:")
	fmt.Println("1. Lösungen Aufgabe 1")
	fmt.Println("2. Lösungen Aufgabe 2")
	fmt.Println("3. GO Advanced 3")
	fmt.Println("4. BMI Berechnung")
	fmt.Println("5. Beenden")
	var eingabe int
	fmt.Scanf("%d", &eingabe)
	switch eingabe {
	case 1:
		loesungen1()
	case 2:
		loesungen2()
	case 3:
		advanced()
	case 4:
		bmiaufruf()

	case 5:
		fmt.Println("Auf Wiedersehen!")
	default:
		fmt.Println("Falsche Eingabe!")
	}

}

// Lösungen zu Aufgaben I und II

func loesungen1() {

	Aufgabe1.Datatypes()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe1.Eigenschaften()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe1.Rechner()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe1.Microservice()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

}

func loesungen2() {

	Aufgabe2.Variablen()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe2.Code()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe2.Array()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe2.Arraylive()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe2.Funktionen()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	Aufgabe2.Reku()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

}
func advanced() {

	advanced2.Advanced1()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	advanced2.Advanced2()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	advanced2.Advanced3()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

	advanced2.Advanced4()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")

}

func bmiaufruf() {
	advanced2.Bmi()

}