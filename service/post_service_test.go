package service

import (
	"entity"
	"errors"
	"repository"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called(post)
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) Delete(post *entity.Post) (int64, error) {
	args := mock.Called(post)
	result := args.Get(0)
	return result.(int64), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) Truncate() error {
	return nil
}

func (mock *MockRepository) CloseDB() error {
	return nil
}

func TestSave(t *testing.T) {
	mockRepo := new(MockRepository)

	inPost := entity.Post{Title: "text1", Text: "Text1"}
	expPost := entity.Post{ID: 1, Title: "text1", Text: "Text1"}

	mockRepo.On("Save", &inPost).Return(&expPost, nil)

	service := NewPostService(mockRepo)
	res, err := service.Save(&inPost)
	mockRepo.AssertExpectations(t)

	assert.Nil(t, err, "Service should return with out a err")
	assert.Equal(t, res, &expPost, "Service should return the same expPost as a mocked repo")
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

func Test_service_Validate(t *testing.T) {
	tests := []struct {
		name    string
		repo    repository.PostRepo
		post    *entity.Post
		wantErr bool
		expErr  error
	}{
		{
			name:    "Post is nil",
			repo:    nil,
			post:    nil,
			wantErr: true,
			expErr:  errors.New("the post is empty"),
		},
		{
			name:    "Post with empty title",
			repo:    nil,
			post:    &entity.Post{},
			wantErr: true,
			expErr:  errors.New("the post title is empty"),
		},
		{
			name:    "Post is correct",
			repo:    nil,
			post:    &entity.Post{ID: 1, Title: "title1", Text: "text1"},
			wantErr: false,
			expErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.repo,
			}
			err := s.Validate(tt.post)
			if (err != nil) && tt.wantErr { // if I want a err and there is a err
				if strings.Compare(err.Error(), tt.expErr.Error()) != 0 { // if the errors' messages are equal
					t.Errorf("Wrong Validate() error = %v, expected: %v", err, tt.expErr)
				}
			}
			if (err != nil) != tt.wantErr { // if I don't want a err but there is a error
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
