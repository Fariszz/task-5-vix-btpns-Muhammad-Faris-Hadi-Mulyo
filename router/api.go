package router

import (
	"GOlangRakamin/controllers"
	"GOlangRakamin/middlewares"
	"GOlangRakamin/modules/auth"
	"GOlangRakamin/modules/photo"
	"GOlangRakamin/modules/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	userRepository := user.NewRepository()
	photoRepository := photo.NewRepository()

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	photoService := photo.NewService(photoRepository)

	userController := controllers.NewUserController(userService, authService)
	photoController := controllers.NewPhotoController(photoService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/register", userController.RegisterUser)
	api.POST("/login", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", middlewares.AuthMiddleware(authService, userService), userController.UploadAvatar)

	api.GET("/photos", photoController.GetPhotos)
	api.POST("/photos", middlewares.AuthMiddleware(authService, userService), photoController.CreatePhoto)
	api.PUT("/photos/:id", middlewares.AuthMiddleware(authService, userService), photoController.UpdatePhoto)
	api.DELETE("/photos/:id", middlewares.AuthMiddleware(authService, userService), photoController.DeletePhoto)

	return router

}
