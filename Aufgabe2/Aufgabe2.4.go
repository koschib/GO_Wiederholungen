package Aufgabe2

func Arraylive() {

	println("Ich bin Aufgabe 2.4: Array Live")
	println("-----------------------------------------------------")

	var Array1 [10]int
	Array1[0] = 2
	for i := 1; i < 10; i++ {
		Array1[i] = Array1[i-1] * 2

		println(Array1[i])

		// Filtern Werte von Index 4 bis 9  als „slice1“!
		slice1 := Array1[4:10]
		println("Slice1: ", slice1)

		// Kapazität und Länge
		println("Kapazität:", cap(slice1))
		println("Länge:", len(slice1))
		println("-----------------------------------------------------")

		var sum int
		for i := 0; i < len(slice1); i++ {
			sum += slice1[i]
			println("Summe Inhalte Index: ", sum)

		}
		// Element aus Array abfragen
		println("Element 4 : Array1[4]: ", Array1[4])
		println("-----------------------------------------------------")

		// Element 2 in Array1 ändern
		Array1[2] = 100
		println("Element 2 in Array1 ändern: ", Array1[2])
		println("-----------------------------------------------------")

		//aus demselben int-Array „Array1“ ein Slice mit der Funktion make(). Der
		//Bereich sollte mindestens über 5 Index Werte gehen. Das neue Slice soll „slice2“ heißen.

		// make() aus Array1 erzeugen 5 Index Werte
		slice2 := make([]int, 20)
		println("Slice2: ", slice2)
		println("-----------------------------------------------------")

		// Slice2 aus Array1 erzeugen
		slice2 = Array1[2:7]

		//  Slice2 ausgeben
		println("Slice2: ", slice2)

		// Element aus Slice 2 abfragen
		println("Element 2 aus Slice2 : Slice2[2]: ", slice2[2])
		println("-----------------------------------------------------")

		// Slice2 erweitern
		slice2 = append(slice2, 1, 2, 3)
		println("Slice2: ", slice2)
	}
}
