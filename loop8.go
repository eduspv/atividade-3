package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("digite um numero: ")
	fmt.Scan(&n1)
	for x := 1; x >= 1; x++ {
		if n1%x == 0 {
			fmt.Println(x)

		}

	}
}
