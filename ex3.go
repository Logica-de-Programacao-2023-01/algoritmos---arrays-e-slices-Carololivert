package main

import "fmt"

func main() {
	var x, c string
	var y int

	fmt.Print("Qual a palavra?:")
	fmt.Scan(&x)
	fmt.Print("Qual a caractere:")
	fmt.Scan(&x)

	for _, i := range x {
		if string(i) == c {
			y++

		}
	}
	fmt.Printf("O caractere %s ocorre %d vezes na string %s.\n", c, y, x)
}