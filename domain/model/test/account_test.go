package test

import (
	"testing"

	"github.com/joaoptgaino/imersao-codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "987"
	name := "JoaoBank"
	bank, err := model.NewBank(code, name)

	accountNumber := "act_0001"
	ownerName := "Joao Gaino"

	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.Bank.ID, bank.ID)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
