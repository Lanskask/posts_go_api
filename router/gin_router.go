package router

import (
	"github.com/gin-gonic/gin"
)

type ginRouter struct {
	router *gin.Engine
}

func NewGinRouter() Router {
	return &ginRouter{router: gin.Default()}
}

func (r ginRouter) Get(uri string, f handleFunc) {
	r.router.GET(uri, gin.WrapF(f))
}

func (r ginRouter) Post(uri string, f handleFunc) {
	r.router.POST(uri, gin.WrapF(f))
}

func (r ginRouter) ListenAndServe(port string) error {
	return r.router.Run(port)
}
