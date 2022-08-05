package advanced

import "fmt"

func Advanced3() {

	fmt.Println("Ich bin Advanced3")
	fmt.Println("-----------------------------------------------------")

	// Variable alter durch scan erfassen

	var alter int
	fmt.Print("Bitte geben Sie ihr Alter ein: ")
	fmt.Scan(&alter)

	// Variable alter ausgeben
	fmt.Println("Ihr Alter ist: ", alter)
	fmt.Println("-----------------------------------------------------")

	// If else um alter zu prüfen und Fehlermeldung auszugeben

	if alter < 0 {
		fmt.Println("Bitte geben Sie eine positive Zahl ein!")
		Advanced3()
	} else if alter >= 13 && alter <= 17 {
		fmt.Println("PG-13")
	} else if alter >= 18 {
		fmt.Println("R-Rated")
	} else {
		fmt.Println("PG-bewertet")
	}

}
