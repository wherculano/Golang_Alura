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

	var danielle *ContaCorrente
	danielle = new(ContaCorrente)
	danielle.titular = "Danielle"
	danielle.numeroAgencia = 321
	danielle.numeroConta = 7890
	danielle.saldo = 4852
	// when comparing values using this way above (x == y) we need to use pointers (*x == y*)
	// otherwise the address will be compared (&x == &y)

	fmt.Println(wagner)
	fmt.Println(matheus)
	fmt.Println(*danielle)
}