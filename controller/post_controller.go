package controller

import (
	"encoding/json"
	"entity"
	"errs"
	"fmt"
	"net/http"
	"service"
)

type IPostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type postController struct {
	service service.IPostService
}

func NewPostController(service service.IPostService) *postController {
	return &postController{service: service}
}

func (c postController) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	all, err := c.service.FindAll()
	if err != nil {
		processErr(resp, fmt.Sprintf("Error get all posts from repo: %s", err))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(all)
}

func (c postController) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		processErr(resp, fmt.Sprintf("Error unmarshalling the posts array: %s", err))
		return
	}

	if err := c.service.Validate(&post); err != nil {
		processErr(resp, fmt.Sprintf("Error validating a post: %s", err))
	}

	saved, err := c.service.Save(&post)
	if err != nil {
		processErr(resp, fmt.Sprintf("error saving post into repo: %s", err))
		return
	}

	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(saved)
}

func processErr(resp http.ResponseWriter, mess string) {
	resp.WriteHeader(http.StatusInternalServerError)
	serviceErr := errs.NewServiceError(mess).Error()
	json.NewEncoder(resp).Encode(serviceErr)
}
