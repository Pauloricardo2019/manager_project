package company

import (
	companyRepo "gerenciador/adapter/database/sql/company"
	"gerenciador/internal/model"
)

func GetByID(
	id uint64,
	getByID companyRepo.GetByIDFn,
) (int, *model.Company, error) {

	found, Company, err := getByID(id)
	if err != nil {
		return 500, nil, err
	}

	if !found {
		return 404, nil, err
	}

	return 200, Company, nil
}
