package main

import (
	"fmt"
)

func main(){
	// Tipe Data Numerik Non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	// Tipe Data Numerik Desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// Tipe Data bool
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// Tipe Data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	// Nilai nil & Zero Value
		// Zero value dari string adalah "" (string kosong).
		// Zero value dari bool adalah false.
		// Zero value dari tipe numerik non-desimal adalah 0.
		// Zero value dari tipe numerik desimal adalah 0.0.

	
}