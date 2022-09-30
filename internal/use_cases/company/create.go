package company

import (
	"crypto/sha256"
	"fmt"
	companyRepo "gerenciador/adapter/database/sql/company"
	"gerenciador/internal/constants"
	"gerenciador/internal/error_map"
	"gerenciador/internal/model"
	"regexp"
)

func encodePassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum)
}

func validatePassword(password string) error {

	if password == "" {
		return error_map.WrapError(error_map.ErrValidationCompany, "password cannot to be empty")
	}

	tests := []string{".{7,}", "([a-z]{1,})", "([A-Z]{1,})", "([0-9]{1,})", "([!@#$&*]{1,})"}
	for index, test := range tests {
		t, err := regexp.MatchString(test, password)
		if err != nil {
			return error_map.WrapError(error_map.ErrValidationCompany, err.Error())
		}
		if !t {
			if index == 0 {
				return error_map.WrapError(error_map.ErrValidationCompany, "this password must be 7 characters long")
			}
			if index == 1 {
				return error_map.WrapError(error_map.ErrValidationCompany, "this password must be at least 1 letter")
			}
			if index == 2 {
				return error_map.WrapError(error_map.ErrValidationCompany, "this password must be at least 1 uppercase letter")
			}
			if index == 3 {
				return error_map.WrapError(error_map.ErrValidationCompany, "this password must be at least 1 number")
			}
			if index == 4 {
				return error_map.WrapError(error_map.ErrValidationCompany, "this password must be at least 1 special character")
			}
		}
	}

	return nil
}

func Create(
	company *model.Company,
	create companyRepo.CreateFn,
) (int, error) {
	err := company.Validate()
	if err != nil {
		return 400, err
	}

	err = validatePassword(company.Password)
	if err != nil {
		return 400, err
	}

	hashedPassword := encodePassword(company.Password)
	company.Password = ""
	company.HashedPassword = hashedPassword

	company.Status = constants.StatusActive

	err = create(company)
	if err != nil {
		return 500, err
	}

	return 201, nil
}
