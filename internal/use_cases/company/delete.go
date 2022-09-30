package company

import companyRepo "gerenciador/adapter/database/sql/company"

func Delete(
	id uint64,
	getByID companyRepo.GetByIDFn,
	delete companyRepo.DeleteFn,
) (int, error) {

	found, _, err := getByID(id)
	if err != nil {
		return 500, err
	}

	if !found {
		return 404, err
	}

	err = delete(id)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
