package company

import (
	"gerenciador/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList_OK(t *testing.T) {

	listMock := func() ([]model.Company, error) {

		addressMock := model.Address{
			ID:           1,
			Street:       "Rua tal tal",
			Complement:   "Case",
			Number:       10,
			Neighborhood: "Vila",
			City:         "Nova Cidade",
			State:        "Acre",
			ZipCode:      "120000-000",
		}

		CompanyMock := model.Company{
			ID:        1,
			Owner:     "Joao",
			AddressID: 1,
			Address:   addressMock,
		}

		return []model.Company{
			CompanyMock,
		}, nil

	}
	statusCode, Companys, err := List(
		listMock,
	)

	require.NoError(t, err)

	assert.Equal(t, statusCode, 200)
	assert.True(t, len(Companys) > 0)

}
