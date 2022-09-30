package dto

import (
	"gerenciador/internal/model"
	"time"
)

type User struct {
	FirstName string    `json:"first_name" valid:"notnull" `
	LastName  string    `json:"last_name" valid:"notnull"`
	CPF       string    `json:"cpf" valid:"notnull"`
	DDD       string    `json:"ddd" valid:"notnull"`
	Phone     string    `json:"phone" valid:"notnull"`
	Birth     time.Time `json:"birth" valid:"-"`
} // @name User

func (dto *User) ParseFromVO(user *model.User) {
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.DDD = user.DDD
	dto.Phone = user.Phone
	dto.Birth = user.Birth

}

func (dto *User) ConvertToVO() *model.User {
	result := &model.User{}

	result.FirstName = dto.FirstName
	result.LastName = dto.LastName
	result.DDD = dto.DDD
	result.Phone = dto.Phone
	result.Birth = dto.Birth

	return result
}

func (dto *User) ParseFromArrayVO(users []model.User) []User {
	var usersDTO []User

	for _, user := range users {
		usersDTO = append(usersDTO, User{

			FirstName: user.FirstName,
			LastName:  user.LastName,
			DDD:       user.DDD,
			Phone:     user.Phone,
			Birth:     user.Birth,
		})
	}
	return usersDTO
}
