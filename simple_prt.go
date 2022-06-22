package main

import (
	"fmt"
	"net/http"
)

func simplePrt(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Up and running ...")
}
