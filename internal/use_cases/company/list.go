package company

import (
	companyRepo "gerenciador/adapter/database/sql/company"
	"gerenciador/internal/model"
)

func List(
	list companyRepo.ListFn,
) (int, []model.Company, error) {

	Companys, err := list()
	if err != nil {
		return 500, nil, err
	}

	return 200, Companys, nil
}
