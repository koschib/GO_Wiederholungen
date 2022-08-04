package advanced

import "fmt"

func Advanced2() {
	println("Ich bin Advanced Aufgabe 2:")
	println("-----------------------------------------------------")

	// eingabe per Scan holen vom User
	var eingabe int
	fmt.Println("Bitte geben Sie eine Zahl ein zwischen 1 und 100:")
	fmt.Scanf("%d", &eingabe)
	fmt.Println("Ihre Zahl ist:", eingabe)
	println("-----------------------------------------------------")

	if eingabe > 60 {
		println("Die Person ist alt geworden.")
	} else if eingabe > 30 {
		println("Die Person ist weiser geworden.")
	} else if eingabe > 20 {
		println("Die Person ist erwachsen geworden.")
	} else if eingabe > 10 {
		println("Die Person ist noch jung.")
	} else {
		println("Wächst noch auf.")
	}
}
