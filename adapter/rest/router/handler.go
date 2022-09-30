package router

import (
	CompanyController "gerenciador/adapter/rest/controller/v1/company"
	loginController "gerenciador/adapter/rest/controller/v1/login"
	userController "gerenciador/adapter/rest/controller/v1/user"
	"gerenciador/adapter/rest/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.RouterGroup) {
	//adminGroup := router.Group("admin")
	//{
	//	adminGroup.POST("/login", adminController.Login)
	//}
	//
	loginGroup := router.Group("login")
	{
		loginGroup.POST("/", loginController.Login)
	}

	CompanyGroup := router.Group("company")
	{
		CompanyGroup.POST("/", CompanyController.Create)
		CompanyGroup.GET("", CompanyController.List)
		CompanyGroup.GET("/:id", CompanyController.GetByID)
		CompanyGroup.PUT("/:id", CompanyController.Update)
		CompanyGroup.DELETE("/:id", CompanyController.Delete)
	}

	userGroup := router.Group("user")
	{
		userGroup.GET("/:id", middleware.Auth(), userController.GetByID)
		userGroup.POST("/", userController.Create)
		userGroup.PUT("/:id", middleware.Auth(), userController.Update)
		userGroup.DELETE("/:id", middleware.Auth(), userController.Delete)
	}

}
