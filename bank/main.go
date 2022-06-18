package main

import (
	"Alura/bank/accounts"
	"Alura/bank/clients"
	"fmt"
)


func PayBankSlip(account verifyAccount, value float64){
	account.Withdraw(value)
}


type verifyAccount interface{
	Withdraw(value float64)(string, bool)
}


func main() {
	wagner := accounts.CheckingAccount{
		Owner: clients.Owner{
			Name:       "Wagner Herculano",
			CPF:        "123.456.789-10",
			Occupation: "Developer",
		},
		AgencyNumber:  012,
		AccountNumber: 3456,
	}
	wagner.Deposit(10720)

	matheus := clients.Owner{"Matheus", "987.654.321-01", "Student"}
	matheusAccount := accounts.CheckingAccount{
		Owner:         matheus,
		AgencyNumber:  123,
		AccountNumber: 4567,
	}
	matheusAccount.Deposit(10531.0)

	var danielleAccount *accounts.CheckingAccount
	danielleAccount = new(accounts.CheckingAccount)
	danielle := clients.Owner{"Danielle", "741.258.963-31", "housewife"}
	danielleAccount.Owner = danielle
	danielleAccount.AgencyNumber = 321
	danielleAccount.AccountNumber = 7890
	danielleAccount.Deposit(4852)
	// when comparing values using this way above (x == y) we need to use pointers (*x == y*)
	// otherwise the address will be compared (&x == &y)

	fmt.Println(wagner)
	fmt.Println(matheusAccount)
	fmt.Println(*danielleAccount)

	message, status_withdraw := wagner.Withdraw(-10)
	fmt.Println("Wagner ->", message, "Status:", status_withdraw)

	status_deposit, new_balance := matheusAccount.Deposit(1000)
	fmt.Println("Matheus ->", status_deposit, "Balance:", new_balance)

	fmt.Println(danielleAccount.Transfer(852, &wagner))
	fmt.Println(wagner)
	fmt.Println(*danielleAccount)
	fmt.Println(danielleAccount.GetBalance())

	joao := clients.Owner{Name: "João", CPF: "321.654.987-00", Occupation: "bricklayer"}
	joaoAccount := accounts.SavingAccount{Owner: joao, NumberAccount: 753, NumberAgency: 159}
	joaoAccount.Deposit(1000)
	fmt.Println(joaoAccount)
	fmt.Println(joaoAccount.Withdraw(987))
	fmt.Println("Saldo Joao:", joaoAccount.GetBalance())

	PayBankSlip(&joaoAccount, 13)
	fmt.Println(joaoAccount.GetBalance())

	fmt.Println("Saldo Danielle:", danielleAccount.GetBalance())
	PayBankSlip(danielleAccount, 100)
	fmt.Println(danielleAccount.GetBalance())

	fmt.Println("Saldo Wagner:", wagner.GetBalance())
	PayBankSlip(&wagner, 100)
	fmt.Println(wagner.GetBalance())

}
