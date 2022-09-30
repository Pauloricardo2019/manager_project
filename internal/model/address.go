package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Address struct {
	ID           uint64 `valid:"-" gorm:"primaryKey"`
	Street       string `valid:"required" gorm:"varchar(50)"`
	Complement   string `valid:"required" gorm:"varchar(20)"`
	Number       uint32 `valid:"required" gorm:"int"`
	Neighborhood string `valid:"required" gorm:"varchar(30)"`
	City         string `valid:"required" gorm:"varchar(50)"`
	State        string `valid:"required" gorm:"varchar(50)"`
	ZipCode      string `valid:"required" gorm:"varchar(12)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
} // @name Address

func (a *Address) Validate() error {
	return validator.New().Struct(a)
}

func (a *Address) TableName() string {
	return "address"
}
