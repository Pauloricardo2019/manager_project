package company

import (
	"gerenciador/internal/model"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate_OK(t *testing.T) {
	faker := faker.New()

	addressMock := model.Address{
		ID:           1,
		Street:       faker.Address().StreetName(),
		Complement:   faker.Address().SecondaryAddress(),
		Number:       faker.UInt32(),
		Neighborhood: faker.Address().CountryAbbr(),
		City:         faker.Address().Country(),
		State:        faker.Address().State(),
		ZipCode:      faker.Address().PostCode(),
	}

	companyMock := &model.Company{
		Owner:     faker.Company().Name(),
		CNPJ:      faker.RandomStringWithLength(13),
		Login:     faker.RandomStringWithLength(10),
		Password:  "T@est123",
		AddressID: 1,
		Address:   addressMock,
	}

	createMock := func(companyMock *model.Company) error {
		companyMock.ID = 1
		return nil
	}

	statusCode, err := Create(
		companyMock,
		createMock,
	)
	require.NoError(t, err)

	assert.Equal(t, statusCode, 201)
	assert.Equal(t, uint64(1), companyMock.ID)
	assert.Equal(t, "Status_Active", companyMock.Status)

}
