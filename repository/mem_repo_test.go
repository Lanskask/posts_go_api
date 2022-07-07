package repository

import (
	"entity"
	"reflect"
	"sort"
	"testing"
)

func TestMemRepo_FindAll(t *testing.T) {
	tests := []struct {
		name      string
		inputData []entity.Post
		wantErr   bool
	}{
		{
			name: "Simple test",
			inputData: []entity.Post{
				{
					ID:    1,
					Title: "title1",
					Text:  "test1",
				},
				{
					ID:    2,
					Title: "title2",
					Text:  "test2",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		repo, _ := NewMemRepo()
		for i := range tt.inputData {
			repo.Save(&tt.inputData[i])
		}

		t.Run(tt.name, func(t *testing.T) {
			got, _ := repo.FindAll()

			if !CompareArrs(got, tt.inputData) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.inputData)
			}
		})
	}
}

func CompareArrs(a, b []entity.Post) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool { return a[i].Title > a[j].Title })
	sort.Slice(b, func(i, j int) bool { return b[i].Title > b[j].Title })
	return reflect.DeepEqual(a, b)
}
