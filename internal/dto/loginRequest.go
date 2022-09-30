package dto

import "gerenciador/internal/model"

type LoginRequest struct {
	Login    string `valid:"notnull" json:"email"`
	Password string `valid:"notnull" json:"password"`
} // @name LoginRequest

func (dto *LoginRequest) ConvertToVO() *model.LoginRequest {
	result := &model.LoginRequest{}

	result.Login = dto.Login
	result.Password = dto.Password

	return result
}
