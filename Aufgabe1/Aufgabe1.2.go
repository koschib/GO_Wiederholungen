// package gibt den Namen des Pakets an in dem die Funktion liegt
package Aufgabe1

// Mit Import können andere Pakete importiert werden.
import (
	"fmt"
	"math/rand"
)

// mit func können Funktionen definiert werden. Wir können mehrere Funktionen
// in einer GO Datei definieren. Diese werden dann in der main Methode aufgerufen.
// Die Funktionen werden mit dem Schlüsselwort func beginnen und mit dem Namen der Funktion
// definiert. Die Funktionen können Parameter haben. Die Parameter werden mit dem Schlüsselwort
// () geklammert.
// Die Funktionen können auch Return Werte haben. Diese werden mit dem Schlüsselwort
// return geklammert.
// Die Funktionen können auch einen Block enthalten. Dieser Block wird mit dem Schlüsselwort
// {} geklammert.

func Eigenschaften() {

	// randomNumber ist eine Variable vom Typ int, die eine Zufallszahl zwischen 0 und XX
	randomNumber := rand.Intn(49)

	// Mit println können wir das Ergebnis ausgeben auf der Konsole.
	fmt.Println("Ich bin Aufgabe 1.2: Zufallszahl")
	println("Ich bin eine Zufallszahl", randomNumber)
	println("Im Code von Aufgabe2/aufgabe2.go sind die Erläuterungen sichtbar")

}

// In der Main Methode können wir den Programmfluss steuern.
// In diesem Beispiel wird die Funktion println aufgerufen.
func main() {
	println("Hallo zusammen")

}
