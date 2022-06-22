package service

import (
	"entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.Post{ID: 1, Title: "text1", Text: "Text1"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	service := NewPostService(mockRepo)
	res, err := service.FindAll()
	mockRepo.AssertExpectations(t)

	assert.Nil(t, err, "Service should return with out a err")
	assert.Equal(t, res, []entity.Post{post}, "Service should return the same post as a mocked repo")
}
