package etcd

import "github.com/taadis/letgo/registry"

type authKey struct {
}

type logConfigKey struct {
}

type authCreds struct {
	Username string
	Password string
}

// Auth 允许指定 username/password
func Auth(username string, password string) registry.Option {

}
