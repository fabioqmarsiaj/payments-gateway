package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Ammount = 900
	assert.Nil(t, transaction.IsValid())

}

func TestTransactions_IsNotValidWithAmmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Ammount = 1001
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())

}

func TestTransactions_IsNotValidWithAmmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Ammount = 0
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the ammount must be greater than 1", err.Error())

}
