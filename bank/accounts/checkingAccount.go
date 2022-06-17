package accounts

import "Alura/bank/clients"


type CheckingAccount struct{
	Owner clients.Owner
	AgencyNumber int
	AccountNumber int
	Balance float64
}


func (c *CheckingAccount) Withdraw(value float64)(string, bool ){
	allowedWithdrawal := value > 0 && value <= c.Balance 
	if allowedWithdrawal{
		c.Balance -= value
		return "Withdrawal successful.", true
	}else {
		return "Insufficient Balance or invalid value.", false
	}
}


func (c *CheckingAccount) Deposit(value float64)(string, float64){
	if value > 0 {
		c.Balance += value
		return "Deposit successful.", c.Balance
	}else {
		return "Invalid value.", c.Balance
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