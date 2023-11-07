// https://leetcode.com/problems/simple-bank-system/

package simple_bank_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleBankSystem(t *testing.T) {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	assert.Equal(t, true, bank.Withdraw(3, 10)) // return true, account 3 has a balance of $20, so it is valid to withdraw $10.
	// Account 3 has $20 - $10 = $10.
	assert.Equal(t, true, bank.Transfer(5, 1, 20)) // return true, account 5 has a balance of $30, so it is valid to transfer $20.
	// Account 5 has $30 - $20 = $10, and account 1 has $10 + $20 = $30.
	assert.Equal(t, true, bank.Deposit(5, 20)) // return true, it is valid to deposit $20 to account 5.
	// Account 5 has $10 + $20 = $30.
	assert.Equal(t, false, bank.Transfer(3, 4, 15)) // return false, the current balance of account 3 is $10,
	// so it is invalid to transfer $15 from it.
	assert.Equal(t, false, bank.Withdraw(10, 50)) // return false, it is invalid because account 10 does not exist.
}
