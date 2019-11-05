package user

import (
	"fmt"
	//"html/template"
	//"io"
	"net/http"
	//"github.com/taadis/letgo/store"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RequestURI)
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	case http.MethodDelete:
		delete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("该方法不支持, 当前路由: %s %s", r.Method, r.RequestURI)))
	}
}

//
func get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("未实现"))
}

//
func post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

/*
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
*/

//
func delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("未实现"))
}
