package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Plan struct {
	ID        uint64    `valid:"-" gorm:"primaryKey"`
	Name      string    `valid:"required" gorm:"varchar(50)"`
	CompanyID uint64    `valid:"required" gorm:"column:company_id"`
	Company   Company   `valid:"required" gorm:"foreignKey:company_id;references:id"`
	ExpiresAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:""`
	CreatedAt time.Time `valid:""`
} // @name Plan

func (p *Plan) Validate() error {
	return validator.New().Struct(p)
}

func (Plan) TableName() string {
	return "plan"
}
