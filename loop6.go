package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("forneca o numero inteiro")
	fmt.Scan(&n1)
	for i := 1; i <= 10; i++ {
		fmt.Println("a sua tabuada fica", n1*i)

	}
}
