// https://leetcode.com/problems/simple-bank-system/

package simple_bank_system

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
	}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !b.checkAccountValidaty(account1) || !b.checkAccountValidaty(account2) {
		return false
	}

	account1Balance := b.balance[account1-1]
	if account1Balance < money {
		return false
	}

	b.balance[account2-1] += money
	b.balance[account1-1] -= money

	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	if !b.checkAccountValidaty(account) {
		return false
	}

	b.balance[account-1] += money
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if !b.checkAccountValidaty(account) {
		return false
	}

	balance := b.balance[account-1]
	if balance < money {
		return false
	}

	b.balance[account-1] -= money
	return true
}

func (b *Bank) checkAccountValidaty(account int) bool {
	return !(account < 1 || account > len(b.balance))
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
