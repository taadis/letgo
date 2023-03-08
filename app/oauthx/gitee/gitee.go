// Package gitee provides constants for using OAuth2 to access Gitee.
package gitee // import "golang.org/x/oauth2/github"

import (
	"os"

	"golang.org/x/oauth2"
)

// Endpoint is Gitee's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://gitee.com/oauth/authorize",
	TokenURL: "https://gitee.com/oauth/token",
}

func ClientID() string {
	return os.Getenv("GITEE_CLIENT_ID")
}

func ClientSecret() string {
	return os.Getenv("GITEE_CLIENT_SECRET")
}
