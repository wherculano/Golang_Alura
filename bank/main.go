package main

import "fmt"

type ContaCorrente struct{
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func main(){
	wagner := ContaCorrente{
		titular: "Wagner Herculano", 
		numeroAgencia: 012, 
		numeroConta: 3456,
		saldo: 10720,
	}

	matheus := ContaCorrente{"Matheus", 123, 4567, 10531.0}

	fmt.Println(wagner)
	fmt.Println(matheus)
}