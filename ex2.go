package main

import "fmt"

func main() {
	var x string
	fmt.Print("Qual a palavra?")
	fmt.Scan(&x)

	s := x
	fmt.Println("Número de caracteres: =", len(s))

}
