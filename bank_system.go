package main

type SavingsAccount struct {
	balence int
}

type CheckingAccount struct {
	balence int
}

type InvestmentAccount struct {
	balence int
}

func (acc *SavingsAccount) MonthlyInterest() int { return acc.balence * 5 / 1200 }

func (acc *CheckingAccount) MonthlyInterest() int { return acc.balence / 1200 }

func (acc *InvestmentAccount) MonthlyInterest() int { return acc.balence * 2 / 1200 }

func (acc *SavingsAccount) CheckBalance() int { return acc.balence }

func (acc *CheckingAccount) CheckBalance() int { return acc.balence }

func (acc *InvestmentAccount) CheckBalance() int { return acc.balence }

func (acc *SavingsAccount) Withdraw(balence int) string {
	if balence <= 0 {
		return "Amount cannot be negative"
	}
	if acc.balence < balence {
		return "Account balance is not enough"
	} else {
		acc.balence = acc.balence - balence
		return "Success"
	}
}

func (acc *SavingsAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	} else {
		acc.balence = acc.balence + amount
		return "Success"
	}
}

func (acc *CheckingAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	} else {
		acc.balence = acc.balence + amount
		return "Success"
	}
}

func (acc *InvestmentAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	} else {
		acc.balence = acc.balence + amount
		return "Success"
	}
}

func (acc *CheckingAccount) Withdraw(balence int) string {
	if balence <= 0 {
		return "Amount cannot be negative"
	}
	if acc.balence < balence {
		return "Account balance is not enough"
	} else {
		acc.balence = acc.balence - balence
		return "Success"
	}
}

func (acc *InvestmentAccount) Withdraw(balence int) string {
	if balence <= 0 {
		return "Amount cannot be negative"
	}
	if acc.balence < balence {
		return "Account balance is not enough"
	} else {
		acc.balence = acc.balence - balence
		return "Success"
	}
}

func (acc *SavingsAccount) Transfer(receiver Account, amount int) string {

	_, ok := receiver.(*SavingsAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}
	_, ok = receiver.(*CheckingAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}
	_, ok = receiver.(*InvestmentAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}

	return "Invalid receiver account"
}

func (acc *CheckingAccount) Transfer(receiver Account, amount int) string {

	_, ok := receiver.(*SavingsAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}
	_, ok = receiver.(*CheckingAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}
	_, ok = receiver.(*InvestmentAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}

	return "Invalid receiver account"
}

func (acc *InvestmentAccount) Transfer(receiver Account, amount int) string {

	_, ok := receiver.(*SavingsAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}
	_, ok = receiver.(*CheckingAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}
	_, ok = receiver.(*InvestmentAccount)
	if ok {
		if output := acc.Withdraw(amount); output == "Success" {
			if out := receiver.Deposit(amount); out == "Success" {
				return "Success"
			}
		}
		return acc.Withdraw(amount)
	}

	return "Invalid receiver account"
}

func NewSavingsAccount() *SavingsAccount {
	return &SavingsAccount{balence: 0}
}

func NewCheckingAccount() *CheckingAccount {
	return &CheckingAccount{balence: 0}
}

func NewInvestmentAccount() *InvestmentAccount {
	return &InvestmentAccount{balence: 0}
}
