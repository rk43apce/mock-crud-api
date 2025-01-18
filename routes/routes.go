package routes

import (
	"mock-crud-api/controllers"
	"mock-crud-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware()) // Add CORS middleware

	router.POST("/login", controllers.LoginHandler)

	// router.GET("/dashboard", controllers.DashboardHandler)

	//Protected route group
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/dashboard", controllers.DashboardHandler)

	v1 := router.Group("/api/v1/products")
	{

		v1.GET("", controllers.GetProducts)
		v1.POST("", controllers.CreateProduct)
		v1.GET("/:id", controllers.GetProductByID)
		v1.PUT("/:id", controllers.UpdateProduct)
		v1.DELETE("/:id", controllers.DeleteProduct)
	}

	return router
}
