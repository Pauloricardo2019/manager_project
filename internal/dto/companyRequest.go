package dto

import "gerenciador/internal/model"

type CompanyRequest struct {
	Owner     string  `json:"owner"  valid:"required"`
	CNPJ      string  `json:"cnpj" valid:"required"`
	Login     string  `json:"login" valid:"required"`
	Password  string  `json:"password" valid:"required"`
	AddressID uint64  `json:"address_id" valid:"required"`
	Address   Address `json:"address" valid:"required"`
}

func (c *CompanyRequest) ConvertToVO() *model.Company {
	companyVO := &model.Company{}

	companyVO.Owner = c.Owner
	companyVO.CNPJ = c.CNPJ
	companyVO.Login = c.Login
	companyVO.Password = c.Password
	companyVO.AddressID = c.AddressID
	addressDTO := c.Address
	addressVO := addressDTO.ConvertToVO()
	companyVO.Address = *addressVO

	return companyVO
}
