package main

import "fmt"

func main() {
	var maior int
	var valor int

	for {
		fmt.Print("Digite um número (ou 0 para sair): ")
		fmt.Scan(&valor)

		if valor == 0 {
			break
		}

		if valor > maior {
			maior = valor
		}
	}
	if maior == 0 {
		fmt.Println("nenhum numero foi digitado")

	} else {
		fmt.Printf("o maior numero é %d\n", maior)
	}
}
