package main

import "fmt"

func main(){

	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)
	// 0 0 false "" <nil> -> Valores padrão (zeros) da linguagem GO
}