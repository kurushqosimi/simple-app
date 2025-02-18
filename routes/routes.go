package routes

import "simple-app/controllers"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", controllers.GetUser)
}
