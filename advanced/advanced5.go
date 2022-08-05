package advanced

import (
	_ "bufio"
	"fmt"
	_ "os"
	_ "strconv"
)

func Bmi() {

	fmt.Println("BMI Rechner")
	fmt.Println("-----------------------------------------------------")

	var weight float64
	var height float64
	var bmi float64

	fmt.Println("Gewicht in kg: ")
	fmt.Scanf("%f", &weight)
	fmt.Println("Ihre eingabe:", weight)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Größe in m: ")
	fmt.Scanf("%f", &height)
	fmt.Println("Ihre eingabe:", height)
	fmt.Println("-----------------------------------------------------")

	bmi = weight / (height * height)

	// BMI ausgeben
	if bmi < 18.5 {
		fmt.Println("Sie sind untergewichtig.")
	} else if bmi >= 18.5 && bmi <= 25 {
		fmt.Println("Sie sind normalgewichtig.")
	} else if bmi > 25 && bmi <= 30 {
		fmt.Println("Sie sind übergewichtig.")
	} else if bmi > 30 && bmi <= 35 {
		fmt.Println("Sie sind obegewichtig.")
	} else if bmi > 35 && bmi <= 40 {
		fmt.Println("Sie sind extreme obegewichtig.")
	} else {
		fmt.Println("Sie sind ernsthaft obegewichtig.")
	}

}
