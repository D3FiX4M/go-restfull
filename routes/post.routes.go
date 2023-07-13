package routes

import (
	"github.com/D3FiX4M/go-restfull/controller"
	"github.com/D3FiX4M/go-restfull/middleware"
	"github.com/gin-gonic/gin"
)

type PostControllerRoute struct {
	postController controller.PostController
}

func NewPostControllerRoute(postController controller.PostController) PostControllerRoute {
	return PostControllerRoute{postController: postController}
}

func (cr *PostControllerRoute) PostRoute(rg *gin.RouterGroup) {
	router := rg.Group("/post")
	router.Use(middleware.DeserializeUser())

	router.GET("/:postId", cr.postController.GetById)
	router.GET("", cr.postController.GetAll)
	router.POST("", cr.postController.Create)
	router.PUT("/:postId", cr.postController.Update)
	router.DELETE("/:postId", cr.postController.Delete)
}
