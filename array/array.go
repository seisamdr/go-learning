package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "manggo", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println("Isi element pertama \t", fruits[0])
}