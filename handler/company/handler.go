package company

import (
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		io.WriteString("home")
	}
}

/*
func Index(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "index")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		t, err := template.ParseFiles("views/home/index.html", "views/shared/footer.html")
		if err != nil {
			io.WriteString(w, err.Error())
		}
		err = t.Execute(w, IndexModel{
			Now:            time.Now().Format("2006-01-02 15:04:05"),
			RuntimeVersion: runtime.Version(),
		})
		if err != nil {
			io.WriteString(w, err.Error())
		}
	}
}
*/
