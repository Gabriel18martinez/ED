package main

import (
	"fmt"
	"projetinho/utils"
	"projetinho/utils/outroUtils"
)



func main() {
	fmt.Println("Ola, mundo!")

	fmt.Println(utils.Somar(5,6))
	fmt.Println(utils.Subtrair(4.3, 6.2))
	fmt.Println(utils.Multiplicar(4.3, 6.2))
	fmt.Println(outroUtils.Dividir(4.3, 6.2))
}

