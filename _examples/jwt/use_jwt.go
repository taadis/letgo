package main

import (
	"fmt"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

//
type CustomPayload struct {
	jwt.Payload
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"bar,omitempty"`
}

//
func main() {
	fmt.Println("use_jwt.go")

	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			Issuer:         "gbrlsnchs",
			Subject:        "someone",
			Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
		Foo: "foo",
		Bar: 1337,
	}

	token, err := jwt.Sign(pl, jwt.NewHS256([]byte("secret")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
