package main

import "fmt"

func main() {
	var n1, n2 int
	var class = ""

	fmt.Scan(&n1)
	for i := 2; i <= 5; i++ {
		fmt.Scan(&n2)

		if n2 > n1 {
			if class == "" {
				class = "C"
			} else if class == "D" {
				class = "N"
				break
			}
		}

		if n1 < n2 {
			if class == ""{
				class = "D"
			} else if class == "C" {
				class = "N"
				break
			}
		}

		n1 = n2
	}

	fmt.Println(class)
}