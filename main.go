package main

import (
	"net/http"

	"github.com/zs368/gin-example/internal/routes"
)

func main() {
	r := routes.NewRouter()
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	s.ListenAndServe()
}
