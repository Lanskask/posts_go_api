package service

import (
	"crypto/rand"
	"entity"
	"errors"
	"math"
	"math/big"
	"repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct {
	repo repository.PostRepo
}

func NewPostService(repo repository.PostRepo) PostService {
	return &service{repo: repo}
}

func (s service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("the post is empty")
	}
	if post.Title == "" {
		return errors.New("the post title is empty")
	}
	return nil
}

func (s service) Save(post *entity.Post) (*entity.Post, error) {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	post.ID = int(bigInt.Int64())

	return s.repo.Save(post)
}

func (s service) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}
