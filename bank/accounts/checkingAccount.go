package accounts

import "Alura/bank/clients"


type CheckingAccount struct{
	Owner clients.Owner
	AgencyNumber int
	AccountNumber int
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


func (c *CheckingAccount) GetBalance()float64{
	return c.balance
}
