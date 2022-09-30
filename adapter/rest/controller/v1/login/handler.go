package login

import (
	companyRepo "gerenciador/adapter/database/sql/company"
	tokenRepo "gerenciador/adapter/database/sql/token"
	"gerenciador/adapter/rest/utils"
	"gerenciador/internal/dto"
	loginUseCase "gerenciador/internal/use_cases/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	loginDTO := &dto.LoginRequest{}

	err := c.BindJSON(loginDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}

	loginRequestVO := loginDTO.ConvertToVO()

	statusCode, token, err := loginUseCase.Login(
		loginRequestVO,
		tokenRepo.CreateToken,
		companyRepo.GetByLogin,
	)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	c.JSON(http.StatusCreated, token)
}
