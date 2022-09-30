package user

import (
	"gerenciador/adapter/database/sql/user"
	"gerenciador/internal/model"
)

func Create(
	user *model.User,
	createUser user.CreateFn,
) (int, error) {

	err := user.Validate()
	if err != nil {
		return 400, err
	}

	err = createUser(user)
	if err != nil {
		return 500, err
	}

	return 201, nil
}
