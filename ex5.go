package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	var n int
	fmt.Print("Qual o nome?")
	fmt.Scan(&x)
	fmt.Print("Qual o numero?")
	fmt.Scan(&n)
	if n > len(x) {
		n = len(x)
	}
	x = strings.ToUpper(x[:n]) + x[n:]
	fmt.Println("A palavra com as primerias letras em maiusculo:", x)

}
