package user

import (
	userRepo "gerenciador/adapter/database/sql/user"
	"gerenciador/adapter/rest/utils"
	"gerenciador/internal/dto"
	userUseCase "gerenciador/internal/use_cases/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetByID - receives id as a parameter, and returns a user object
// @Summary - Get one user
// @Description Get user by id
// @Tags - User
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.User
// @Router /user/{id} [get]
// @Security ApiKeyAuth
func GetByID(c *gin.Context) {
	userID := c.Param("id")
	userIDParam, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	_, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	statusCode, user, err := userUseCase.GetByID(uint64(userIDParam), userRepo.GetByID)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	userDTO := &dto.User{}
	userDTO.ParseFromVO(user)

	c.JSON(http.StatusOK, userDTO)
}

// @BasePath /v1

// Create - create a user objects, and returns a user objects
// @Summary - Create user
// @Description - Create a user
// @Tags - User
// @Accept json
// @Produce json
// @Param User body dto.User true "User to be created"
// @Success 201 {object} dto.User
// @Router /user [post]
// @Security ApiKeyAuth
func Create(c *gin.Context) {
	userDTO := &dto.User{}

	err := c.BindJSON(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	userVO := userDTO.ConvertToVO()
	statusCode, err := userUseCase.Create(
		userVO,
		userRepo.Create,
	)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	userDTO.ParseFromVO(userVO)
	c.JSON(http.StatusCreated, userDTO)
}

// @BasePath /v1

// Update - receives id as a parameter, update a user, and returns a user objects
// @Summary - Update user
// @Description - Update a user by id
// @Tags - User
// @Accept json
// @Produce json
// @Param User body dto.User true "User to be updated"
// @Success 200 {object} dto.User
// @Router /user/{id} [put]
// @Security ApiKeyAuth
func Update(c *gin.Context) {
	userDTO := &dto.User{}

	err := c.ShouldBindJSON(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	paramID, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := paramID.(uint64)

	userVO := userDTO.ConvertToVO()

	err = userUseCase.Update(userID, userVO, userRepo.Update, userRepo.GetByEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}

// @BasePath /v1

// Delete - receives id as a parameter, and returns a string
// @Summary - Delete user
// @Description - Delete a user by id
// @Tags - User
// @Accept json
// @Produce json
// @Success 204
// @Param id path int true "id"
// @Error 500 {object} dto.Error
// @Router /user/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	userId := c.Param("id")
	userIDParam, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	_, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	statusCode, err := userUseCase.Delete(uint64(userIDParam), userRepo.Delete)
	if err != nil {
		utils.BurstError(c, err, statusCode)
	}

	c.JSON(http.StatusOK, nil)
}
