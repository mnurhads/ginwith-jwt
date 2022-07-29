package main

import (
	"ginwith-jwt/controllers"
	"ginwith-jwt/database"
	"ginwith-jwt/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root@tcp(localhost:3306)/jwt_learn?parseTime=true")
	database.Migrate()

	// set router
	router := initRouter()
	router.Run(":9090")
}

// func router
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/generate-token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping-connection", controllers.Ping)
			secured.GET("/user-all", controllers.GetAllUser)
		}
	}
	return router
}
