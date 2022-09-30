package Company

import (
	companyRepo "gerenciador/adapter/database/sql/company"
	"gerenciador/adapter/rest/utils"
	"gerenciador/internal/dto"
	CompanyUseCase "gerenciador/internal/use_cases/company"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @BasePath /v1

// Create - create a company
// @Summary - Create company
// @Description - Create a company
// @Tags - User
// @Accept json
// @Produce json
// @Param Company body dto.Company true "Company"
// @Success 201 {object} dto.Company
// @Router /company [post]
// @Security ApiKeyAuth
func Create(c *gin.Context) {
	Company := &dto.Company{}

	err := c.BindJSON(Company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}

	subVO := Company.ConvertToVO()

	statusCode, err := CompanyUseCase.Create(
		subVO,
		companyRepo.Create,
	)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	Company.ParseFromVO(subVO)

	c.JSON(http.StatusCreated, Company)
}

// @BasePath /v1

// List - List all Companys
// @Summary - List all Companys
// @Description - List all Companys
// @Tags - User
// @Accept json
// @Produce json
// @Success 200 {array} model.Company
// @Router /company [get]
// @Security ApiKeyAuth
func List(c *gin.Context) {

	statusCode, companies, err := CompanyUseCase.List(
		companyRepo.List,
	)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	c.JSON(http.StatusOK, companies)

}

// @BasePath /v1

// GetByID - create a company
// @Summary - GetByID company
// @Description - GetByID a company
// @Tags - User
// @Accept json
// @Produce json
// @param id path string true "Company id"
// @Param Company body dto.Company true "Company"
// @Success 200 {object} dto.Company
// @Router /company/:id [get]
// @Security ApiKeyAuth
func GetByID(c *gin.Context) {
	idParam := c.Param("id")
	companyID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}

	statusCode, company, err := CompanyUseCase.GetByID(
		companyID,
		companyRepo.GetByID,
	)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	c.JSON(http.StatusOK, company)
}

// @BasePath /v1

// Update - create a company
// @Summary - Update company
// @Description - Update a company
// @Tags - User
// @Accept json
// @Produce json
// @Param Company body dto.Company true "Company"
// @Param id path string true "company id"
// @Success 200
// @Router /company/:id [put]
// @Security ApiKeyAuth
func Update(c *gin.Context) {
	idParam := c.Param("id")
	companyID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}
	companyDTO := &dto.Company{}
	companyVo := companyDTO.ConvertToVO()

	statusCode, err := CompanyUseCase.Update(
		companyID,
		companyVo,
		companyRepo.Update,
	)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	c.JSON(http.StatusOK, nil)
}

// @BasePath /v1

// Delete - create a company
// @Summary - Create company
// @Description - Create a company
// @Tags - User
// @Accept json
// @Produce json
// @Param id path string true "company id"
// @Success 200 {object} dto.Company
// @Router /company [post]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	idParam := c.Param("id")
	companyID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}

	statusCode, err := CompanyUseCase.Delete(
		companyID,
		companyRepo.GetByID,
		companyRepo.Delete,
	)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	c.JSON(http.StatusOK, nil)
}
