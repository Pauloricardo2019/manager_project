package rest

import (
	"context"
	"fmt"
	"gerenciador/adapter/config"
	v1 "gerenciador/adapter/rest/router"
	"gerenciador/docs"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var httpServer *http.Server

func SetupHttpEngine() *gin.Engine {
	cfg := config.GetConfig()

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())

	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Title = "gerenciador - API"
	docs.SwaggerInfo.Description = "API para cadastro e gerenciamento de usuário"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.Use(sentryGin.New(sentryGin.Options{
		WaitForDelivery: true,
	}))

	routerV1 := engine.Group("/v1")
	{
		v1.InitializeRouter(routerV1)
	}

	engine.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.RestPort),
		Handler: engine,
	}

	go func() {
		fmt.Println("Listening on port", cfg.RestPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err.Error())
		}
	}()

	return engine
}

func StopRest(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctxWithTimeout)
	if err != nil {
		fmt.Println("http server forced to shutdown due to timeout")
	}

	fmt.Println("http server was gracefully stopped")
}
