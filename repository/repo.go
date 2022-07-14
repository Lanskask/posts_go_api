package repository

import (
	"posts_api/entity"
)

const (
	tableName = "posts"
)

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	// Delete returns Last Updated Id or error
	Delete(post *entity.Post) (int64, error)
	Truncate() error
	CloseDB() error
}
