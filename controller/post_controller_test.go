package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"posts_api/entity"
	"posts_api/repository"
	"posts_api/service"
	"sort"
	"testing"
)

func TestAddPost(t *testing.T) {
	// initialization
	repo, err := repository.NewSQLiteRepo(false)
	if err != nil {
		t.Fatalf("failed to create a repo")
	}
	serv := service.NewPostService(repo)
	contr := NewPostController(serv)

	expPost := entity.Post{
		Title: "Title 1",
		Text:  "Text 1",
	}
	expStatus := http.StatusCreated

	// prepare test
	respBody := []byte(fmt.Sprintf(`{"title": "%s", "text": "%s"}`, expPost.Title, expPost.Text))
	req, _ := http.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(respBody))

	handler := http.HandlerFunc(contr.AddPost)

	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != expStatus {
		t.Errorf("Handler returned a wrong status returned: %d, expected: %d", status, expStatus)
	}

	var post entity.Post
	json.NewDecoder(io.Reader(resp.Body)).Decode(&post)
	if !ComparePosts(post, expPost) {
		t.Errorf("Not expected result. got %v, expected: %v", post, expPost)
	}

	cleanDB(repo)
}

func TestGetPosts(t *testing.T) {
	// initialization
	//repo, err := repository.NewSQLiteRepo(false)
	//repo, err := repository.NewMemRepo()
	repo, err := repository.NewFirebaseRepo()

	if err != nil {
		t.Fatalf("failed to create a repo: %s", err)
	}
	_ = repo.Truncate()

	serv := service.NewPostService(repo)
	contr := NewPostController(serv)

	expPosts := []entity.Post{
		{
			ID:    11,
			Title: "Title 1",
			Text:  "Text 1",
		},
		{
			ID:    12,
			Title: "Title 2",
			Text:  "Text 2",
		},
	}
	expStatus := http.StatusOK

	for i := 0; i < len(expPosts); i++ {
		_, _ = repo.Save(&expPosts[i])
	}

	// prepare test
	req, _ := http.NewRequest(http.MethodGet, "/posts", nil)

	handler := http.HandlerFunc(contr.GetPosts)

	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != expStatus {
		t.Errorf("Handler returned a wrong status returned: %d, expected: %d", status, expStatus)
	}

	var posts []entity.Post
	json.NewDecoder(io.Reader(resp.Body)).Decode(&posts)
	if !ComparePostsArr(posts, expPosts) {
		t.Errorf("Not expected result. got %v, expected: %v", posts, expPosts)
	}

	cleanDB(repo)
}

func ComparePosts(a, b entity.Post) bool {
	if a.Title != b.Title || a.Text != b.Text {
		return false
	}
	return true
}

func ComparePostsArr(a, b []entity.Post) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool { return a[i].Title > a[j].Title })
	sort.Slice(b, func(i, j int) bool { return b[i].Title > b[j].Title })

	for i := range a {
		if !ComparePosts(a[i], b[i]) {
			return false
		}
	}
	return true
}

func cleanDB(repo repository.PostRepo) {
	repo.Truncate()
	repo.CloseDB()
}
