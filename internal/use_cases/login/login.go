package login

import (
	"crypto/sha256"
	"fmt"
	companyRepo "gerenciador/adapter/database/sql/company"
	tokenRepo "gerenciador/adapter/database/sql/token"
	"gerenciador/internal/error_map"
	"gerenciador/internal/model"
	"github.com/google/uuid"
	"time"
)

func encodePassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum)
}

func generateUserToken(ID uint64) (*model.Token, error) {
	UserToken := &model.Token{
		Value:     uuid.New().String(),
		CompanyID: ID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	return UserToken, nil
}

func Login(
	loginRequest *model.LoginRequest,
	createToken tokenRepo.CreateTokenFn,
	getByLogin companyRepo.GetByLoginFn,
) (int, *model.Token, error) {

	hash := encodePassword(loginRequest.Password)

	found, company, err := getByLogin(loginRequest.Login, hash)
	if err != nil {
		return 500, nil, err
	}

	if !found {
		return 404, nil, err
	}

	token, err := generateUserToken(company.ID)
	if err != nil {
		return 400, nil, error_map.WrapError(error_map.ErrValidateLogin, "Cannot generate Token")
	}

	tokenCreated, err := createToken(token)
	if err != nil {
		return 500, nil, err
	}

	return 201, tokenCreated, nil
}
