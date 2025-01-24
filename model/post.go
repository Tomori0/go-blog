package model

import (
	"go-blog-zleng/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	CategoryId int       `json:"category_id"`
	UserId     int       `json:"user_id"`
	ViewCount  int       `json:"view_count"`
	Type       int       `json:"type"`
	CreatAt    time.Time `json:"created_at"`
}

type PostMore struct {
	Pid          int           `json:"pid"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"`
	Content      template.HTML `json:"content"`
	CategoryId   int           `json:"category_id"`
	CategoryName string        `json:"category_name"`
	UserId       int           `json:"user_id"`
	UserName     string        `json:"user_name"`
	ViewCount    int           `json:"view_count"`
	Type         int           `json:"type"`
	CreateAt     string        `json:"created_at"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"category_id"`
	UserId     int    `json:"user_id"`
	Type       int    `json:"type"`
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}

type PostResp struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
