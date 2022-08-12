package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // Lembrar: um float literal é sempre do tipo float64
	y := 2
	// fmt.Println(x / y) -> Lembrar: não possível realizar cálculos entre valores de diferentes tipos
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(123))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123)) // Itoa é uma função que converte int para string

	// string para int
	num, _ := strconv.Atoi("123") // Atoi é uma função que converte string para int
	fmt.Println(num - 122)

	// string para boolean
	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}

	

}