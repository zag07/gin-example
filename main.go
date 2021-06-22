package main

import (
	"net/http"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/internal/pkg/routing"
)

func main() {
	r := routing.NewRouter()
	s := &http.Server{
		Addr:    ":" + configs.App.Port(),
		Handler: r,
	}

	s.ListenAndServe()
}
