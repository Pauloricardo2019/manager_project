package dto

import "gerenciador/internal/model"

type Company struct {
	Company   string  `json:"company" validate:"required"`
	Owner     string  `json:"owner" validate:"required"`
	Password  string  `json:"password" validate:"required"`
	AddressID uint64  `json:"address_id" validate:"required"`
	Address   Address `json:"address" validate:"required"`
}

func (dto *Company) ConvertToVO() *model.Company {
	var SubVo model.Company

	SubVo.Owner = dto.Owner
	SubVo.Password = dto.Password
	SubVo.AddressID = dto.AddressID
	addressVO := dto.Address
	SubVo.Address = *addressVO.ConvertToVO()

	return &SubVo
}

func (dto *Company) ParseFromVO(Company *model.Company) {

	dto.Owner = Company.Owner
	dto.Password = Company.Password
	dto.AddressID = Company.AddressID
	addressDTO := Address{}
	addressDTO.ParseFromVO(&Company.Address)
	dto.Address = addressDTO
}
