package test

import (
	"testing"

	"github.com/joaoptgaino/imersao-codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewBank(t *testing.T) {
	code := "987"
	name := "JoãoBank"
	bank, err := model.NewBank(code, name)
	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(bank.ID))
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, err = model.NewBank("", name)
	require.NotNil(t, err)
}
