package routes

import (
	"go-socialmedia/controllers"
	"go-socialmedia/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Butuh login
		api.POST("/upload", middleware.AuthMiddleware(), controllers.UploadPost)
		
		api.GET("/posts", controllers.GetAllPosts)

		// user Profile
		api.GET("/user/:username/posts", controllers.GetPostsByUsername)

		// myProfile
		api.GET("/me", middleware.AuthMiddleware(), controllers.GetMyProfile)

	}
}
