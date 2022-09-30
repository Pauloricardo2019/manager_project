package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Category struct {
	ID        uint64    `valid:"-" gorm:"primaryKey"`
	Name      string    `valid:"required" gorm:"varchar(20)"`
	ProductID uint64    `valid:"required" gorm:"column:product_id"`
	Product   Product   `valid:"required" gorm:"foreignKey:product_id;references:id"`
	UpdatedAt time.Time `valid:"-"`
	CreatedAt time.Time `valid:"-"`
} // @name Category

func (c *Category) Validate() error {
	return validator.New().Struct(c)
}

func (Category) TableName() string {
	return "category"
}