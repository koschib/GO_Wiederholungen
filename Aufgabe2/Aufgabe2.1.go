package Aufgabe2

import "fmt"

func Variablen() {

	fmt.Println("Ich bin Aufgabe 2.1: Variablen")
	// Variablen deklarieren und initialisieren

	// 1. Variablen anlegen und initialisieren mit Datentyp und Wert
	var a int = 10
	fmt.Println("a =", a)

	var b int = 20
	fmt.Println("b =", b)

	// Initialisierung mit Rechnung
	var c int = a + b
	fmt.Println("c =", c)

	// ohne Datentyp
	var d = 5
	fmt.Println("d =", d)

}
