package main

import "fmt"

func main() {
	var soma, contador, valor int

	for {
		fmt.Print("Digite um número (ou 0 para sair): ")
		fmt.Scan(&valor)

		if valor == 0 {
			break
		}

		soma += valor
		contador++
	}

	if contador == 0 {
		fmt.Println("Nenhum número foi digitado")
	} else {
		media := float64(soma) / float64(contador)
		fmt.Printf("A média dos %d números digitados é %.2f\n", contador, media)
	}
}
