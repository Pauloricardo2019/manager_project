package user

import (
	userRepo "gerenciador/adapter/database/sql/user"
	"gerenciador/internal/model"
)

func GetByID(
	userID uint64,
	getByID userRepo.GetByIDFn,
) (int, *model.User, error) {

	found, user, err := getByID(userID)
	if err != nil {
		return 500, nil, err
	}

	if !found {
		return 404, nil, err
	}

	return 200, user, nil
}
