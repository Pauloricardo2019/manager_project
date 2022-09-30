package company

import (
	companyRepo "gerenciador/adapter/database/sql/company"
	"gerenciador/internal/model"
)

func Update(
	id uint64,
	company *model.Company,
	update companyRepo.UpdateFn,
) (int, error) {
	company.ID = id

	err := company.Validate()
	if err != nil {
		return 400, err
	}

	err = update(company)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
