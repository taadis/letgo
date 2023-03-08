package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/taadis/letgo/app/oauthx/gitee"
	"golang.org/x/oauth2"
)

// main PasswordCredentialsToken模式可行
func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     gitee.ClientID(),
		ClientSecret: gitee.ClientSecret(),
		Endpoint:     gitee.Endpoint,
		RedirectURL:  "",
		Scopes:       []string{"user_info", "emails"},
	}

	token, err := conf.PasswordCredentialsToken(ctx, "username", "password")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("token data:%+v", token)

	client := conf.Client(ctx, token)
	rsp, err := client.Get("https://gitee.com/api/v5/users/{username}")
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	data := make(map[string]interface{})
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data:%+v", data)

	rsp, err = client.Get("https://gitee.com/api/v5/repos/taadis/mirrors-in-china/branches")
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		log.Fatalf("status code error, code:%d", rsp.StatusCode)
	}

	bs, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("branchs data:%+v", string(bs))
}

// Code Token模式需要用户回调授权,纯后端好像不好实现?
func mainWithCode() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     gitee.ClientID(),
		ClientSecret: gitee.ClientSecret(),
		Endpoint:     gitee.Endpoint,
		RedirectURL:  "",
		Scopes:       nil, // []string{"user_info", "emails"},
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Printf("visit the URL for the auth dialog: %v\n", url)

	token, err := conf.Exchange(ctx, "code")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, token)
	rsp, err := client.Get("https://gitee.com/api/v5/users/taadis")
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	data := make(map[string]interface{})
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data:%+v", data)
}
