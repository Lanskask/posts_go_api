package router

import (
	"net/http"
)

//type handleFunc = func(res http.ResponseWriter, req *http.Request)
type handleFunc = func(http.ResponseWriter, *http.Request)

type Router interface {
	Get(uri string, f handleFunc)
	Post(uri string, f handleFunc)
	ListenAndServe(port string) error
}
