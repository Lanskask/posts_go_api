package repository

import (
	"entity"
)

const (
	collectionName = "posts"
)

type IPostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
