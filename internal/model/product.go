package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Product struct {
	ID        uint16    `valid:"-" gorm:"primaryKey"`
	Name      string    `valid:"required" gorm:"varchar(60)"`
	Price     string    `valid:"required" gorm:"varchar(60)"`
	Quantity  uint32    `valid:"required" `
	Status    string    `valid:"-" gorm:"varchar(60)"`
	CreatedAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:"-"`
} // @name Product

func (p *Product) Validate() error {
	return validator.New().Struct(p)
}

func (Product) TableName() string {
	return "product"
}
