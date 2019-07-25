package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//
type Article struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

//
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//
type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

func main() {
	log.Println("use_json.go")
	article := Article{
		Id:      1,
		Title:   "标题",
		Content: "内容",
		Author: Author{
			Id:   2,
			Name: "zhang san",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "一楼",
				Author: Author{
					Id:   5,
					Name: "王五",
				},
			},
			Comment{
				Id:      6,
				Content: "小板凳",
				Author: Author{
					Id:   7,
					Name: "小七",
				},
			},
		},
	}
	bytes, err := json.Marshal(&article)
	if err != nil {
		log.Panicln(err)
	}
	err = ioutil.WriteFile("some.json", bytes, 0644)
	if err != nil {
		log.Panicln(err)
	}
}
