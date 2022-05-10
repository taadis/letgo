package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	github_client_id     = "GITHUB_CLIENT_ID"
	github_client_secret = "GITHUB_CLIENT_SECRET"
)

func buildTokenUrl(clientId string, clientSecret string, code string, redirectUrl string) (string, error) {
	uri, err := url.Parse("https://github.com")
	if err != nil {
		return "", err
	}
	uri.Path = "/login/oauth/access_token"

	values := url.Values{}
	values.Add("client_id", clientId)
	values.Add("client_secret", clientSecret)
	values.Add("code", code)
	values.Add("redirect_url", redirectUrl)
	uri.RawQuery = values.Encode()
	return uri.String(), nil
}

func buildUserUrl() (string, error) {
	uri, err := url.Parse("https://api.github.com")
	if err != nil {
		return "", err
	}

	uri.Path = "/user"
	return uri.String(), nil
}

// for parse access_token=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&scope=&token_type=bearer
type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// 三方登录回调地址,用于拿code换token,并用token换userInfo
func callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectUrl := r.URL.Query().Get("redirect_url")

	// 拿code作为参数去github请求access_token
	tokenUrl, _ := buildTokenUrl(os.Getenv(github_client_id), os.Getenv(github_client_secret), code, redirectUrl)
	req, err := http.NewRequest(http.MethodGet, tokenUrl, nil)
	if err != nil {
		log.Printf("http.NewRequest error:%+v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenRsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("http.DefaultClient.Do error:%+v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bs, err := ioutil.ReadAll(tokenRsp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll error:%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("got tokenRsp:%s", string(bs))
	io.WriteString(w, fmt.Sprintf("got tokenRsp %s", string(bs)))

	// 尝试解析,也可以解析道结构体方便维护
	values, err := url.ParseQuery(string(bs))
	if err != nil {
		log.Printf("url.ParseQuery error:%+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken := values.Get("access_token")
	io.WriteString(w, accessToken)
	// return

	// 拿access_token去github请求user
	userUrl, _ := buildUserUrl()
	userReq, err := http.NewRequest(http.MethodGet, userUrl, nil)
	if err != nil {
		log.Printf("http.NewRequest error:%+v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userReq.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))
	useRsp, err := http.DefaultClient.Do(userReq)
	if err != nil {
		log.Printf("http.DefaultClient.Do error:%+v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bs, err = ioutil.ReadAll(useRsp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll error:%+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("got userRsp:%s", string(bs))
	io.WriteString(w, "got userRsp:"+string(bs))
}

func main() {
	fmt.Printf("oauth\n")
	http.HandleFunc("/callback/", callback)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("http.ListenAndServe error:%+v", err)
	}
}
