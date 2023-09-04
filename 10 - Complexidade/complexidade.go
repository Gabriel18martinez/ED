package main

import "fmt"

func busca(matriz [][]int, n int, k int) bool {
	i, j int
	i = 0

	for i < n {
		j = 0
		for j < n {
			if (matriz [i][j] == x) {
				return true // acho
			}
			j++
		}
		i++
	}

	return false // nao achou
}