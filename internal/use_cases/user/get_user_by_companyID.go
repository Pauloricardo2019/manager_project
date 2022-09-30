package user

import (
	companyRepo "gerenciador/adapter/database/sql/company"
	userRepo "gerenciador/adapter/database/sql/user"
	"gerenciador/internal/model"
)

func GetUsersByCompanyID(
	companyID uint64,
	getCompanyID companyRepo.GetByIDFn,
	getAllUsersByCompanyID userRepo.GetUsersByCompanyIDFn,
) (int, []model.User, error) {

	found, company, err := getCompanyID(companyID)
	if err != nil {
		return 500, nil, err
	}

	if !found {
		return 404, nil, err
	}

	users, err := getAllUsersByCompanyID(company.ID)
	if err != nil {
		return 500, nil, err
	}

	return 200, users, nil
}
