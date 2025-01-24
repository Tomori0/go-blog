package router

import (
	"go-blog-zleng/api"
	"go-blog-zleng/views"
	"net/http"
)

func Router() {
	// 1. view 2. data 3. static resource
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
