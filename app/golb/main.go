// app/golb/main.go
package main

import (
	"fmt"
	//"html/template"
	"net/http"

	"github.com/taadis/letgo/app/golb/controllers"
)

// indexView struct
// type indexView struct {
// 	Title string
// 	User  User
// 	Posts []Post
// }

func main() {
	fmt.Println("ready to start golb app")
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	user1 := User{
	// 		UserName: "taadis",
	// 	}
	// 	user2 := User{
	// 		UserName: "miebug",
	// 	}
	// 	posts := []Post{
	// 		Post{
	// 			User: user1,
	// 			Body: "blog1",
	// 		},
	// 		Post{
	// 			User: user2,
	// 			Body: "blog2",
	// 		},
	// 	}
	// 	viewModel := indexView{
	// 		Title: "HomePage",
	// 		User:  user1,
	// 		Posts: posts,
	// 	}
	// 	tmp, err := template.ParseFiles("views/index.html")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// 	err = tmp.Execute(w, &viewModel)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })
	controllers.Startup()
	//mux.HandleFunc("/system/user/", user.HandleFunc)
	//mux.Handle("/system/user/", midleware.PanicAndRecover(http.HandlerFunc(user.HandleFunc)))
	err := http.ListenAndServe(":5903", nil)
	if err != nil {
		fmt.Errorf("start golb error: ", err.Error())
	}
}
