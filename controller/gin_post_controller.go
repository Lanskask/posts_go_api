package controller

import (
	"entity"
	"errs"
	"fmt"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

type ginPostController struct {
	service service.PostService
}

func NewGinPostController(service service.PostService) *ginPostController {
	return &ginPostController{service: service}
}

func (controller ginPostController) GetPosts(c *gin.Context) {
	all, err := controller.service.FindAll()
	if err != nil {
		processGinErr(c, fmt.Sprintf("Error get all posts from repo: %s", err))
		return
	}

	c.JSON(http.StatusOK, all)
}

func (controller ginPostController) AddPost(c *gin.Context) {
	var post entity.Post
	err := c.BindJSON(&post)
	if err != nil {
		processGinErr(c, fmt.Sprintf("Error unmarshalling the posts array: %s", err))
		return
	}

	if err := controller.service.Validate(&post); err != nil {
		processGinErr(c, fmt.Sprintf("Error validating a post: %s", err))
	}

	saved, err := controller.service.Save(&post)
	if err != nil {
		processGinErr(c, fmt.Sprintf("error saving post into repo: %s", err))
		return
	}

	c.JSON(http.StatusOK, saved)
}

func processGinErr(c *gin.Context, mess string) {
	serviceErr := errs.NewServiceError(mess).Error()
	c.JSON(http.StatusInternalServerError, gin.H{"error": serviceErr})
}
