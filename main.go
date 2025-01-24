package main

import (
	"go-blog-zleng/common"
	"go-blog-zleng/router"
	"log"
	"net/http"
)

func init() {
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
