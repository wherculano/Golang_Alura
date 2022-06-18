package accounts

import "Alura/bank/clients"


type SavingAccount struct {
	Owner clients.Owner
	NumberAccount, NumberAgency, Operation int
	balance float64
}

func (c *SavingAccount) Withdraw(value float64)(string, bool ){
	allowedWithdrawal := value > 0 && value <= c.balance 
	if allowedWithdrawal{
		c.balance -= value
		return "Withdrawal successful.", true
	}else {
		return "Insufficient balance or invalid value.", false
	}
}


func (c *SavingAccount) Deposit(value float64)(string, float64){
	if value > 0 {
		c.balance += value
		return "Deposit successful.", c.balance
	}else {
		return "Invalid value.", c.balance
	}
}

func (c *SavingAccount) GetBalance()float64{
	return c.balance
}