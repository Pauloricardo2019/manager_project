package company

import (
	"gerenciador/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetByID_OK(t *testing.T) {

	var id uint64 = 1

	getByIDMock := func(id uint64) (bool, *model.Company, error) {
		return true, &model.Company{
			ID: id,
		}, nil
	}

	statusCode, Company, err := GetByID(
		id,
		getByIDMock,
	)
	require.NoError(t, err)

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, Company.ID, id)

}
