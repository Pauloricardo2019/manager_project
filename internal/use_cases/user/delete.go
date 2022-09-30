package user

import (
	userRepo "gerenciador/adapter/database/sql/user"
)

func Delete(
	userId uint64,
	deleteFn userRepo.DeleteFn,
) (int, error) {

	err := deleteFn(userId)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
