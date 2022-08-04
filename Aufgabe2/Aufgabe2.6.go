package Aufgabe2

import "fmt"

func Reku() {

	fmt.Println("Ich bin Aufgabe 2.6: rekursive Funktion")
	fmt.Println("-----------------------------------------------------")

	fmt.Println("Die rekursive Funktion ist eine Funktion, die sich selbst aufruft.")
	fmt.Println("-----------------------------------------------------")

	rekursiveFunktion(1, 8)
}

func rekursiveFunktion(x int, y int) int {

	if x != y {
		fmt.Println("Stimmt noch nicht, noch eine Runde!")
		return rekursiveFunktion(x+1, y)
	} else {
		fmt.Println("Jetzt passt das. Wir sind fertig!")
		return x

	}
}
