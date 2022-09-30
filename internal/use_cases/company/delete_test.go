package company

import (
	"gerenciador/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelete_OK(t *testing.T) {

	var id uint64 = 2

	deleteMock := func(id uint64) error {
		return nil
	}

	getByID := func(id uint64) (bool, *model.Company, error) {
		return true, &model.Company{
			ID: id,
		}, nil
	}

	statusCode, err := Delete(
		id,
		getByID,
		deleteMock,
	)
	require.NoError(t, err)

	assert.Equal(t, statusCode, 200)

}
