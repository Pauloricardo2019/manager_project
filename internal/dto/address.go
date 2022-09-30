package dto

import "gerenciador/internal/model"

type Address struct {
	ID           uint64 `json:"id" validate:"required"`
	Street       string `json:"street" validate:"required"`
	Complement   string `json:"complement" validate:"required"`
	Number       uint32 `json:"number" validate:"required"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	ZipCode      string `json:"zipcode" validate:"required"`
}

func (dto *Address) ConvertToVO() *model.Address {
	var addressVO model.Address

	addressVO.Street = dto.Street
	addressVO.Complement = dto.Complement
	addressVO.Number = dto.Number
	addressVO.Neighborhood = dto.Neighborhood
	addressVO.City = dto.City
	addressVO.State = dto.State
	addressVO.ZipCode = dto.ZipCode

	return &addressVO

}

func (dto *Address) ParseFromVO(address *model.Address) {

	dto.Street = address.Street
	dto.Complement = address.Complement
	dto.Number = address.Number
	dto.Neighborhood = address.Neighborhood
	dto.City = address.City
	dto.State = address.State
	dto.ZipCode = address.ZipCode

}
