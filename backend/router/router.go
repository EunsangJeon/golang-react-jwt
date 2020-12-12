package router

import (
	"github.com/EunsangJeon/golang-react-jwt/backend/controller"
	"github.com/EunsangJeon/golang-react-jwt/backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middleware.ErrorHandler)
	router.Use(middleware.CORSMiddleware())

	// routes
	router.GET("/ping", controller.Pong)
	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)
	router.GET("/session", controller.Session)

	return router
}
