package main

import "fmt"

type CheckingAccount struct{
	owner string
	agencyNumber int
	accountNumber int
	balance float64
}


func (c *CheckingAccount) Withdraw(value float64)string {
	allowedWithdrawal := value > 0 && value <= c.balance 
	if allowedWithdrawal{
		c.balance -= value
		return "Withdrawal successful."
	}else {
		return "Insufficient balance or invalid value."
	}
}


func main(){
	wagner := CheckingAccount{
		owner: "Wagner Herculano", 
		agencyNumber: 012, 
		accountNumber: 3456,
		balance: 10720,
	}

	matheus := CheckingAccount{"Matheus", 123, 4567, 10531.0}

	var danielle *CheckingAccount
	danielle = new(CheckingAccount)
	danielle.owner = "Danielle"
	danielle.agencyNumber = 321
	danielle.accountNumber = 7890
	danielle.balance = 4852
	// when comparing values using this way above (x == y) we need to use pointers (*x == y*)
	// otherwise the address will be compared (&x == &y)

	fmt.Println(wagner)
	fmt.Println(matheus)
	fmt.Println(*danielle)

	fmt.Println(wagner.Withdraw(-10))
}