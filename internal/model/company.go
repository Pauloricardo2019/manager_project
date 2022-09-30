package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Company struct {
	ID             uint64  `valid:"-" gorm:"primaryKey"`
	Owner          string  `valid:"required" gorm:"varchar(50)"`
	CNPJ           string  `valid:"required" gorm:"varchar(15)"`
	Login          string  `valid:"required" gorm:"varchar(30)"`
	HashedPassword string  `valid:"required" gorm:"varchar(64)"`
	Password       string  `valid:"required" gorm:"varchar(1)"`
	AddressID      uint64  `valid:"required" gorm:"column:address_id"`
	Address        Address `valid:"required" gorm:"foreignKey:address_id;references:id"`
	Status         string  `valid:"-" gorm:"varchar(20)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (c *Company) Validate() error {
	return validator.New().Struct(c)
}

func (Company) TableName() string {
	return "company"
}
