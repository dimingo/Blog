package models

import "html/template"

type BlogPost struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Date    string `json:"date"`
}

type Post struct {
	Title   string
	Content template.HTML
}
