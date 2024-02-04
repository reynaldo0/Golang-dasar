package main

import (
    "fmt"
    "strings"
)

func main() {
    var nilai1 = 10
    var nilai2 = 20

    fmt.Println("SMKN46 Cipinang")
    fmt.Println("==================================")
    loweringText("Maju Terus Pantang Mundur")
    fmt.Println("")

    fmt.Println("Hasilnya penjumlahan adalah : ", tambah(float32(nilai1),float32(nilai2)))
    fmt.Println("Hasilnya kurang adalah      : ", kurang(float32(nilai1),float32(nilai2)))
    fmt.Println("Hasilnya bagi adalah        : ", bagi(float32(nilai1),float32(nilai2)))
    fmt.Println("Hasilnya kali adalah        : ", kali(float32(nilai1),float32(nilai2)))

    listNilai := [...]float64{7.0, 8.5, 9.1}


    listInt := [...]int64{1,2,3,4,5,6,7,8,9,10}

    for _, listNilaiItem := range listNilai {
        fmt.Println("Hasil array: => ", listNilaiItem )
    }

    fmt.Println("")
    for _ , listIntItem := range listInt {
        fmt.Println("Hasil array: => ", listIntItem )
    }
}

func loweringText(textInput string)  {
    fmt.Println(strings.ToLower(textInput))
}

func tambah(nilai1 float32, nilai2 float32) float32 {
    var hitung = nilai1 + nilai2
    return hitung
}
func kurang(nilai1 float32, nilai2 float32) float32 {
    var hitung = nilai1 - nilai2
    return hitung
}
func bagi(nilai1 float32, nilai2 float32) float32 {
    var hitung = nilai1 / nilai2
    return hitung
}
func kali(nilai1 float32, nilai2 float32) float32 {
    var hitung = nilai1 * nilai2
    return hitung
}