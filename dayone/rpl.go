package main

import "fmt"

// func main() {

// 	fmt.Printf("SMKN 46 JAKARTA")
// 	fmt.Printf("===============")
// 	// loweringText("Maju Terus Pantang Mundur")

// 	// fmt.Printf("Hasilnya Adalah " , nilai + nilai2)
// }

// func loweringText(textInput string){
// 	fmt.Println(strings.ToLower(textInput))
// }

func main() {
	var nilai1 = 10
	var nilai2 = 20
	// var total = nilai1 + nilai2
	// fmt.Println("Hasilnya adalaj ==>", total)
	fmt.Println("Hasil tambahnya adalah ==>", addMethod(float32(nilai1), float32(nilai2)))
	fmt.Println("Hasil kurangnya adalah ==>", subtractMethod(float32(nilai1), float32(nilai2)))
	fmt.Println("Hasil kalinya adalah ==>", multiplayMethod(float32(nilai1), float32(nilai2)))
	fmt.Println("Hasil baginya adalah ==>", divideMethod(float32(nilai1), float32(nilai2)))

	listNilai := [...]float64{7.0, 8.5, 9.1}

	for _, lislistNilaiItem := range listNilai {
		fmt.Println("Hasil Array =>", lislistNilaiItem)
	}
}

func addMethod(nilai1 float32, nilai2 float32) float32 {
	var hitung = nilai1 + nilai2
	return hitung
}
func subtractMethod(nilai1 float32, nilai2 float32) float32 {
	var hitung = nilai1 - nilai2
	return hitung
}
func multiplayMethod(nilai1 float32, nilai2 float32) float32 {
	var hitung = nilai1 * nilai2
	return hitung
}
func divideMethod(nilai1 float32, nilai2 float32) float32 {
	var hitung = nilai1 / nilai2
	return hitung
}
