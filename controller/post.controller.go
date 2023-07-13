package controller

import (
	"github.com/D3FiX4M/go-restfull/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB: DB}
}

func (pc *PostController) GetById(ctx *gin.Context) {
	postId := ctx.Param("postId")
	var post models.Post
	result := pc.DB.First(&post, "id = ?", postId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})

}

func (pc *PostController) GetAll(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "20")

	pageCount, _ := strconv.Atoi(page)
	limitCount, _ := strconv.Atoi(limit)
	offset := (pageCount - 1) * limitCount

	var posts []models.Post
	results := pc.DB.Limit(limitCount).Offset(offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

func (pc *PostController) Create(ctx *gin.Context) {

	var payload *models.CreatePostRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	currentUser := ctx.MustGet("currentUser").(models.User)

	newPost := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

func (pc *PostController) Update(ctx *gin.Context) {
	postId := ctx.Param("postId")

	var payload *models.UpdatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	var postById models.Post
	result := pc.DB.First(&postById, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      postById.User,
		CreatedAt: postById.CreatedAt,
		UpdatedAt: time.Now(),
	}

	pc.DB.Model(&postById).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": postById})
}

func (pc *PostController) Delete(ctx *gin.Context) {
	postId := ctx.Param("postId")

	result := pc.DB.Delete(&models.Post{}, "id = ?", postId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
