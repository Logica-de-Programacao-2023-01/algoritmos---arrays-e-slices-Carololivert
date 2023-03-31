package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Qual o nome?")
	fmt.Scan(&x)

	nome := x

	fmt.Println(strings.ToUpper(nome))

}
