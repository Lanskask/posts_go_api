package repository

import (
	"entity"
	"sync"
)

type MemRepo struct {
	mx   sync.RWMutex
	data map[int]*entity.Post
}

func NewMemRepo() (*MemRepo, error) {
	return &MemRepo{
		mx:   sync.RWMutex{},
		data: make(map[int]*entity.Post),
	}, nil
}

func (r *MemRepo) Save(post *entity.Post) (*entity.Post, error) {
	r.mx.Lock()
	r.data[post.ID] = post
	r.mx.Unlock()

	return post, nil
}

func (r *MemRepo) FindAll() ([]entity.Post, error) {
	var res []entity.Post

	r.mx.RLock()
	for _, v := range r.data {
		res = append(res, *v)
	}
	//for k, _ := range r.data {
	//	res = append(res, *r.data[k])
	//}
	r.mx.RUnlock()

	return res, nil
}

func (r *MemRepo) Delete(post *entity.Post) (int64, error) {
	r.mx.Lock()
	r.data[post.ID] = nil
	r.mx.Unlock()

	return int64(post.ID), nil
}

func (r *MemRepo) Truncate() error {
	r.mx.Lock()
	r.data = make(map[int]*entity.Post)
	r.mx.Unlock()

	return nil
}

func (r *MemRepo) CloseDB() error {
	return nil
}
