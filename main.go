package main

import (
	"github.com/D3FiX4M/go-restfull/controller"
	"github.com/D3FiX4M/go-restfull/initializers"
	"github.com/D3FiX4M/go-restfull/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine

	AuthController controller.AuthController
	UserController controller.UserController
	PostController controller.PostController

	AuthControllerRoute routes.AuthControllerRoute
	UserControllerRoute routes.UserControllerRoute
	PostControllerRoute routes.PostControllerRoute
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controller.NewAuthController(initializers.DB)
	UserController = controller.NewUserController(initializers.DB)
	PostController = controller.NewPostController(initializers.DB)

	AuthControllerRoute = routes.NewAuthControllerRoute(AuthController)
	UserControllerRoute = routes.NewUserControllerRoute(UserController)
	PostControllerRoute = routes.NewPostControllerRoute(PostController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/health-checker", func(context *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		context.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthControllerRoute.AuthRoute(router)
	UserControllerRoute.UserRoute(router)
	PostControllerRoute.PostRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
