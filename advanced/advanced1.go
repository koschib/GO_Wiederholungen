package advanced

import "fmt"

func Advanced1() {

	fmt.Println("Ich bin Advanced 1:")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Im verzeichnins GOPath werden alle GO Dateien gespeichert.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Binärdateien werden nach go install unter GOPATH/bin gespeichert.")
	fmt.Println("Binärdateien aus Packages werden in External Libraries gespeichert.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("In GO geben wir erst den Variablennamen an, dann wird der Datentyp deklariert.")

	var a int = 10
	fmt.Println("Ich bin eine Variable mit dem Namen a und der Wert 10.", a)

}
