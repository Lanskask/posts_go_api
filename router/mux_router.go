package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	router *mux.Router
}

func NewMuxRouter() Router {
	return &muxRouter{router: mux.NewRouter()}
}

func (r muxRouter) Get(path string, f handleFunc) {
	r.router.HandleFunc(path, f).Methods(http.MethodGet)
}

func (r muxRouter) Post(path string, f handleFunc) {
	r.router.HandleFunc(path, f).Methods(http.MethodPost)
}

func (r muxRouter) ListenAndServe(port string) error {
	return http.ListenAndServe(port, r.router)
}
