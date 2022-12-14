package user

import (
	"fmt"
	"gerenciador/adapter/database/sql/user"
	userRepo "gerenciador/adapter/database/sql/user"
	"gerenciador/internal/error_map"
	"gerenciador/internal/model"
	emailverifier "github.com/AfterShip/email-verifier"
)

func Update(
	userID uint64,
	user *model.User,
	update userRepo.UpdateFn,
	getUserByEmail user.GetByEmailFn,
) error {

	if user.ID != userID {
		return error_map.WrapError(error_map.ErrValidationUser, "error to update user. You're not authorized to manage this user")
	}

	verifier := emailverifier.NewVerifier()
	res, err := verifier.Verify(user.Email)

	if err != nil {
		return error_map.WrapError(error_map.ErrValidationUser, " verify email address failed")
	}

	if !res.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return error_map.WrapError(error_map.ErrValidationUser, " invalid email")
	}

	_, existingEmail, err := getUserByEmail(user.Email)
	if err != nil {
		return err
	}

	if existingEmail.ID != user.ID && existingEmail.Email == user.Email {
		return error_map.WrapError(error_map.ErrValidationUser, " email already exists")
	}

	err = update(user)
	if err != nil {
		return err
	}

	return nil

}
