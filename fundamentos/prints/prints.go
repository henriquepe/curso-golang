package main

import "fmt"

// Comando fmt.Print imprime e não quebra linha
// Comando fmt.Println imprime e quebra linha

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")
	fmt.Println(" Nova ")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é" + x) <-- Erro
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	// %d - inteiro
	// %f - ponto flutuante
	// %t - booleano
	// %s - string
	// %v - qualquer tipo
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}