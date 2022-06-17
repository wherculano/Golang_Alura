package main

import "fmt"

type CheckingAccount struct{
	owner string
	agencyNumber int
	accountNumber int
	balance float64
}


func (c *CheckingAccount) Withdraw(value float64)(string, bool ){
	allowedWithdrawal := value > 0 && value <= c.balance 
	if allowedWithdrawal{
		c.balance -= value
		return "Withdrawal successful.", true
	}else {
		return "Insufficient balance or invalid value.", false
	}
}


func (c *CheckingAccount) Deposit(value float64)(string, float64){
	if value > 0 {
		c.balance += value
		return "Deposit successful.", c.balance
	}else {
		return "Invalid value.", c.balance
	}
}


func (c *CheckingAccount) Transfer(value float64, destinationAccount *CheckingAccount)string{
	_, withdrawalSuccessful := c.Withdraw(value)
	
	if withdrawalSuccessful {
		destinationAccount.Deposit(value)
		return "Transference successful"
	}else {
		return "Something went wrong"
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

	message, status_withdraw :=  wagner.Withdraw(-10)
	fmt.Println("Wagner ->", message, "Status:", status_withdraw)

	status_deposit, new_balance := matheus.Deposit(1000)
	fmt.Println("Matheus ->", status_deposit, "Balance:", new_balance)

	fmt.Println(danielle.Transfer(852, &wagner))
	fmt.Println(wagner)
	fmt.Println(*danielle)
}