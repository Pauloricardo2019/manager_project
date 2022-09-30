package user

import (
	tokenRepo "gerenciador/adapter/database/sql/token"
	"gerenciador/adapter/database/sql/user"
	"gerenciador/internal/error_map"
	"gerenciador/internal/model"
)

func Login(login *model.LoginRequest, createToken tokenRepo.CreateTokenFn, getUserByEmail user.GetByEmailFn) (bool, *model.Token, error) {

	condition, userFind, err := getUserByEmail(login.Email)
	if err != nil {
		return condition, nil, error_map.WrapError(error_map.ErrValidateLogin, "User not found")
	}

	login.HashedPassword = encodePassword(login.Password)
	login.Password = ""

	if userFind.HashedPassword != login.HashedPassword {
		return false, nil, error_map.WrapError(error_map.ErrValidateLogin, "User not found")
	}

	tokenCreated, err := createToken(token)
	if err != nil {
		return false, nil, error_map.WrapError(error_map.ErrValidateLogin, "Cannot create Token")
	}

	return true, tokenCreated, nil
}
