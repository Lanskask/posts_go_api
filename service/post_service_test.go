package service

import (
	"entity"
	"errors"
	"repository"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_service_Save(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	postRepoMock := repository.NewMockIPostRepo(ctrl)

	inputPost := entity.Post{Title: "text1", Text: "Text1"}
	expectedPots := entity.Post{ID: 1, Title: "text1", Text: "Text1"}

	postRepoMock.
		EXPECT().
		Save(&inputPost).
		Return(&expectedPots, nil)

	s := NewPostService(postRepoMock)
	res, err := s.Save(&inputPost)
	assert.Nil(t, err, "Service should return with out a err")
	assert.Equal(t, &expectedPots, res, "Service should return the same post as a mocked repo")
}

func Test_service_Validate(t *testing.T) {
	tests := []struct {
		name    string
		repo    repository.IPostRepo
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
