package main

import (
	"net/http"

	"github.com/zs368/gin-example/configs"
	_ "github.com/zs368/gin-example/init"
	"github.com/zs368/gin-example/internal/pkg/routing"
)

func main() {
	r := routing.NewRouter()
	s := &http.Server{
		Addr:    ":" + configs.App.HttpPort,
		Handler: r,
	}

	s.ListenAndServe()
}
