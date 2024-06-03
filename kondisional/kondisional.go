package main

import "fmt"

func main() {
	var point = 3
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. nilai anda adalah %d\n", point)
	}
}