package home

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"runtime"
	"time"
)

//
type RuntimeInfo struct {
	Now            string `json:"now"`             // 当前时间
	RuntimeVersion string `json:"runtime_version"` // 运行时版本
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		runtimeInfo := &RuntimeInfo{
			Now:            time.Now().Format("2006-01-02 15:04:05"),
			RuntimeVersion: runtime.Version(),
		}
		result, err := json.Marshal(runtimeInfo)
		if err != nil {
			io.WriteString(w, err.Error())
		}
		io.WriteString(w, string(result))
	}
}

//
func Html(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "index")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		t, err := template.ParseFiles("views/home/index.html", "views/shared/footer.html")
		if err != nil {
			io.WriteString(w, err.Error())
		}
		err = t.Execute(w, RuntimeInfo{
			Now:            time.Now().Format("2006-01-02 15:04:05"),
			RuntimeVersion: runtime.Version(),
		})
		if err != nil {
			io.WriteString(w, err.Error())
		}
	}
}
