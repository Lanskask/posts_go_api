package repository

import (
	"entity"
)

const (
	tableName = "posts"
)

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(post *entity.Post) (int64, error)
}
