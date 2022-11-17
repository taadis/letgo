package main

import (
	"log"
	"net/http"

	_ "github.com/taadis/letgo/app/i18n/internal/translations"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const AcceptLanguage = "Accept-Language"

func newServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)
	return mux
}

func main() {
	mux := newServer()

	err := http.ListenAndServe(":5903", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get(AcceptLanguage)
	log.Printf("got Accept-Language is %s", lang)

	//t, _, err := language.ParseAcceptLanguage(lang)
	//if err != nil {
	//	log.Printf("language.ParseAcceptLanguage error:%+v", err)
	//}

	matcher := language.NewMatcher([]language.Tag{
		language.Chinese,
		language.SimplifiedChinese,
		language.TraditionalChinese,
		language.English,
		language.AmericanEnglish,
		language.BritishEnglish,
	}, language.PreferSameScript(false))
	tag, _ := language.MatchStrings(matcher, lang, r.Header.Get(AcceptLanguage))
	log.Printf("language.MatchStrings tag is %s", tag.String())
	b, _ := tag.Base()
	log.Printf("base is %s", b.String())

	ctxLanguage := "zh-Hans"
	if b.String() == "en" {
		ctxLanguage = "en-US"
	}

	//var tag language.Tag
	//switch lang {
	//case "en-US":
	//	tag = language.MustParse("en-US")
	//default:
	//	tag = language.MustParse("zh-CN")
	//}

	// 使用对应语言初始化一个message.Printer实例
	parsedTag := language.MustParse(ctxLanguage)
	log.Printf("parsedTag is %s", parsedTag.String())
	p := message.NewPrinter(parsedTag)
	// 获取欢迎信息翻译成目标语言
	p.Fprintf(w, "hello")
}
