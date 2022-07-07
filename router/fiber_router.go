package router

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

type fiberRouter struct {
	router *fiber.App
}

func NewFiberRouter() Router {
	return &fiberRouter{router: fiber.New()}
}

func (r fiberRouter) Get(path string, f handleFunc) {
	r.router.Get(path, adaptor.HTTPHandlerFunc(f))
}

func (r fiberRouter) Post(path string, f handleFunc) {
	r.router.Post(path, adaptor.HTTPHandlerFunc(f))
}

func (r fiberRouter) ListenAndServe(port string) error {
	return r.router.Listen(port)
}
