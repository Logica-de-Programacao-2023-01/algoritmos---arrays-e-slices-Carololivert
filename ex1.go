package main

import "fmt"

func main() {
	var x, c string
	fmt.Print("Qual a primeira palavra?")
	fmt.Scan(&x)
	fmt.Print("Qual a segunda palavra?")
	fmt.Scan(&c)

	s1 := x + "," + c

	fmt.Println("resultado :", s1)

}
