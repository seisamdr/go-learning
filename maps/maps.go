package main

import "fmt"

func main() {
	var chickens map[string]int
	chickens = map[string]int{}
	chickens["jan"] = 50
	chickens["feb"] = 30
	fmt.Println(chickens["jan"])
	fmt.Println(chickens["feb"])
}