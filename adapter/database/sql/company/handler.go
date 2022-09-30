package company

import (
	"errors"
	"gerenciador/adapter/database/sql"
	"gerenciador/internal/model"
	"gorm.io/gorm"
)

type CreateFn func(Company *model.Company) error

func Create(Company *model.Company) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	return db.
		Create(Company).
		Error

}

type ListFn func() ([]model.Company, error)

func List() ([]model.Company, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return nil, err
	}

	Companys := make([]model.Company, 0)

	if err = db.
		Find(&Companys).
		Error; err != nil {
		return nil, err
	}

	return Companys, nil
}

type GetByIDFn func(id uint64) (bool, *model.Company, error)

func GetByID(id uint64) (bool, *model.Company, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}

	Company := &model.Company{}

	if err = db.Where(&model.Company{
		ID: id,
	}).
		First(Company).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, Company, nil
		}
		return false, nil, err
	}

	return true, Company, nil

}

type GetByLoginFn func(login string, hashedPassword string) (bool, *model.Company, error)

func GetByLogin(login string, hashedPassword string) (bool, *model.Company, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}

	company := &model.Company{}

	if err = db.
		Where(&model.Company{
			Login:          login,
			HashedPassword: hashedPassword,
		}).
		First(company).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, company, nil
		}
		return false, nil, err
	}
	return true, company, nil
}

type UpdateFn func(Company *model.Company) error

func Update(Company *model.Company) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	if err = db.
		Save(Company).
		Error; err != nil {
		return err
	}

	return nil
}

type DeleteFn func(id uint64) error

func Delete(id uint64) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	if err = db.
		Delete(&model.Company{}, "ID = ?", id).
		Error; err != nil {
		return err
	}

	return nil
}
