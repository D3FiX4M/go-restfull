package routes

import (
	"github.com/D3FiX4M/go-restfull/controller"
	"github.com/D3FiX4M/go-restfull/middleware"
	"github.com/gin-gonic/gin"
)

type AuthControllerRoute struct {
	authController controller.AuthController
}

func NewAuthControllerRoute(authController controller.AuthController) AuthControllerRoute {
	return AuthControllerRoute{authController: authController}
}

func (cr *AuthControllerRoute) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/sign/up", cr.authController.SignUp)
	router.POST("/sign/in", cr.authController.SignIn)
	router.GET("/refresh/token", cr.authController.RefreshToken)
	router.GET("/logout", middleware.DeserializeUser(), cr.authController.Logout)
}
