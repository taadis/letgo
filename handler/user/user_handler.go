package user

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/taadis/letgo/store"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RequestURI)
	switch r.Method {
	case "GET":
		get(w, r)
	case "POST":
		post(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/user/add.html")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		io.WriteString(w, err.Error())
	}
}

// 添加用户
func post(w http.ResponseWriter, r *http.Request) {
	inputModel := new(model.Article)
	inputModel.Title = "测试标题"
	inputModel.Content = "测试内容"
	inputModel.AuthorId = 11
	article := new(model.Article)
	lastInsertId, err := article.Add(inputModel)
	//sql := "insert into Articles(Title, Content, Author1) value(?,?,?)"
	//lastInsertId, err := models.AddRow(sql, model.Title, model.Content, model.Author)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	if lastInsertId <= 0 {
		io.WriteString(w, "Add()失败!")
		return
	}
	io.WriteString(w, "Add()成功!")
}
