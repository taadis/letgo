package user

import (
	"encoding/json"
	"fmt"

	//"path"
	//"html/template"
	//"io"
	"net/http"

	//"github.com/taadis/letgo/store"
	"github.com/taadis/letgo/store/mysql"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RequestURI)
	var err error
	switch r.Method {
	case http.MethodGet:
		err = getFunc(w, r)
	case http.MethodPost:
		err = postFunc(w, r)
	case http.MethodPut:
		err = putFunc(w, r)
	case http.MethodDelete:
		err = deleteFunc(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
func recoverFunc(w http.ResponseWriter, r *http.Request) {
	if rc := recover(); rc != nil {
		s := fmt.Sprintf("panic recover %v", rc)
		fmt.Println(s)
		http.Error(w, s, http.StatusInternalServerError)
	}
}

//
func getFunc(w http.ResponseWriter, r *http.Request) (err error) {
	//w.WriteHeader(http.StatusNotImplemented)
	//err = errors.New("get user error")
	// defer func() {
	// 	if rc := recover(); rc != nil {
	// 		s := fmt.Sprintf("panic recover %v", rc)
	// 		fmt.Println(s)
	// 		http.Error(w, s, http.StatusInternalServerError)
	// 	}
	// }()
	//defer recoverFunc(w, r)
	values := r.URL.Query()
	id := values.Get("id")
	//panic("鬼知道为什么崩溃了!")
	//user := store.User{}
	userStore := mysql.UserStore{}
	user, err := userStore.User(id)
	if err != nil {
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	return
}

//
func postFunc(w http.ResponseWriter, r *http.Request) (err error) {
	w.WriteHeader(http.StatusNotImplemented)
	return
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
func putFunc(w http.ResponseWriter, r *http.Request) (err error) {
	w.WriteHeader(http.StatusNotImplemented)
	return
}

//
func deleteFunc(w http.ResponseWriter, r *http.Request) (err error) {
	w.WriteHeader(http.StatusNotImplemented)
	return
}
