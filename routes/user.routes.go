package routes

import (
	"github.com/D3FiX4M/go-restfull/controller"
	"github.com/D3FiX4M/go-restfull/middleware"
	"github.com/gin-gonic/gin"
)

type UserControllerRoute struct {
	userController controller.UserController
}

func NewUserControllerRoute(userController controller.UserController) UserControllerRoute {
	return UserControllerRoute{userController: userController}
}

func (cr *UserControllerRoute) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/user")

	router.GET("/current", middleware.DeserializeUser(), cr.userController.GetCurrentUser)
}
