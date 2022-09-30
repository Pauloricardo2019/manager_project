package company

import (
	"gerenciador/internal/constants"
	"gerenciador/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdate_OK(t *testing.T) {

	var id uint64 = 2

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

	CompanyMock := &model.Company{
		ID:        1,
		Owner:     "Joao",
		AddressID: 1,
		Address:   addressMock,
	}

	updateMock := func(CompanyMock *model.Company) error {
		CompanyMock.Status = constants.StatusPending
		return nil
	}

	statusCode, err := Update(
		id,
		CompanyMock,
		updateMock,
	)
	require.NoError(t, err)

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, CompanyMock.Status, "Status_Pending")

}
