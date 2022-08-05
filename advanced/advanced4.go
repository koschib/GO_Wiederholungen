package advanced

import "fmt"

func Advanced4() {

	fmt.Println("Ich bin Advanced 4")

	// 4.1
	fmt.Println("Das ist 4.1")
	var i int
	fmt.Print("Bitte geben Sie eine Zahl zwischen 1 und 3 ein: ")
	fmt.Scan(&i)

	switch i {
	case 1:
		fmt.Println("Eins.")
	case 2:
		fmt.Println("Zwei.")
	case 3:
		fmt.Println("Drei.")
	default:
		fmt.Println("Falsche Eingabe!")
	}
	zwei()

}
func zwei() {
	fmt.Println("Das ist 4.2 Unterscheidung Wochentag oder Wochenende")
	var wochentag string
	fmt.Print("Bitte geben Sie einen Wochentag ein ( Beispiel Montag): ")
	fmt.Scan(&wochentag)

	// switch Weekday Montag bis Freitag Ausgabe Es ist ein Wochentag
	switch wochentag {
	case "Montag":
		fmt.Println("Es ist ein Wochentag.")
	case "Dienstag":
		fmt.Println("Es ist ein Wochentag.")
	case "Mittwoch":
		fmt.Println("Es ist ein Wochentag.")
	case "Donnerstag":
		fmt.Println("Es ist ein Wochentag.")
	case "Freitag":
		fmt.Println("Es ist ein Wochentag.")
	case "Samstag":
		fmt.Println("Es ist ein Wochenende.")
	case "Sonntag":
		fmt.Println("Es ist ein Wochenende.")
	default:
		fmt.Println("Falsche Eingabe!")
		zwei()
	}
	drei()

}
func drei() {
	fmt.Println("Das ist 4.3 Unterscheidung Vormittag/Nachmittag ")

	var jetzt int
	fmt.Print("Bitte geben Sie die aktuelle Stunde an (0-24): ")
	fmt.Scan(&jetzt)

	// switch 0-12 Vormittags, 13-24 Nachmittags
	switch jetzt {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		fmt.Println("Es ist Vormittags.")
	case 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		fmt.Println("Es ist Nachmittags.")

	}
	vier()
}
func vier() {
	fmt.Println("Das ist 4.4 Boolean ")

	// unterscheidet ob die Eingabe “boolean”, “int” oder nichts von beiden ist und dies jeweils in einem kurzen Satz ausgibt.

	var eingabe string
	fmt.Print("Bitte geben Sie einen Wert ein: ")
	fmt.Scan(&eingabe)

	// prüfen ob eingabe INT wert oder Boolean wert ist

}
