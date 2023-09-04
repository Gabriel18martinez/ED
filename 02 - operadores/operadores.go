package main

import "fmt"

func main() {
	var x, y, z = 4, 10, -2

	// aritméticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	//atribuição

	z++ // z = z + 1
	z-- // z = z - 1
	z += 2
	z -= 2
	z *= 2
	z /= 2
	z %= 2

	// comparação

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
	fmt.Println(x < y)
	fmt.Println(x <= y)

	// Lógicos

	fmt.Println(true && false) // AND
	fmt.Println(true )
}