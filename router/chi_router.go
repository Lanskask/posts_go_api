package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type chiRouter struct {
	router chi.Router
}

func NewChiRouter() Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return &chiRouter{router: r}
}

func (r chiRouter) Get(uri string, f handleFunc) {
	r.router.Get(uri, f)
}

func (r chiRouter) Post(uri string, f handleFunc) {
	r.router.Post(uri, f)
}

func (r chiRouter) ListenAndServe(port string) error {
	return http.ListenAndServe(port, r.router)
}
