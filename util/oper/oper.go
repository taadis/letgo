// oper project oper.go
package oper

import (
	"errors"
	"os/exec"
	"runtime"
)

// 指定一个链接地址,并调用系统默认浏览器访问
func Access(uri string) (err error) {
	os := runtime.GOOS

	switch os {
	case "windows":
		err = accessOnWindows(uri)
		break
	case "linux":
		err = accessOnLinux(uri)
		break
	case "darwin":
		err = accessOnDarwin(uri)
		break
	default:
		err = errors.New("not support " + os)
		break
	}

	return err
}

// 在 windows 系统下调用浏览器访问
func accessOnWindows(uri string) error {
	cmd := exec.Command("explorer", uri)
	err := cmd.Start()
	return err
}

// call default browser access on linux
func accessOnLinux(uri string) (err error) {
	cmd := exec.Command("x-www-browser", uri)
	err = cmd.Start()
	return err
}

// call default browser access on darwin(Mac)
func accessOnDarwin(uri string) (err error) {
	cmd := exec.Command("open", uri)
	err = cmd.Start()
	return err
}

// TODO: 在其他系统下调用浏览器访问
