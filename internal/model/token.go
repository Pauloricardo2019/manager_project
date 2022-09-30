package model

import "time"

type Token struct {
	Value     string `valid:"notnull" gorm:"varchar(255)"`
	CompanyID uint64
	Company   Company   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" valid:"notnull"`
	CreatedAt time.Time `valid:"-" gorm:"autoCreateTime"`
	ExpiresAt time.Time `valid:"-" gorm:"autoUpdateTime:milli"`
}

func (Token) TableName() string {
	return "token"
}
