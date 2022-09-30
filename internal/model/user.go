package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint64    `valid:"-" gorm:"primaryKey"`
	FirstName string    `valid:"required" gorm:"varchar(60)"`
	LastName  string    `valid:"required" gorm:"varchar(60)"`
	CPF       string    `valid:"required" gorm:"varchar(15)"`
	CompanyID uint64    `valid:"required" gorm:"column:company_id"`
	Company   Company   `valid:"required" gorm:"foreignKey:company_id;references:id"`
	DDD       string    `valid:"-" gorm:"varchar(2)"`
	Phone     string    `valid:"-" gorm:"varchar(9)"`
	Birth     time.Time `valid:"-"`
	AddressID uint64    `valid:"required" gorm:"column:address_id"`
	Address   Address   `valid:"required" gorm:"foreignKey:address_id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Status    string    `gorm:"varchar(20)"`
	CreatedAt time.Time `valid:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `valid:"-" gorm:"autoUpdateTime:milli"`
} // @name User

func (user *User) Validate() error {
	return validator.New().Struct(user)
}

func (User) TableName() string {
	return "user"
}
