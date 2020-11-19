package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"testing"
	"time"
)

// 链接 qq 的可以
func TestSendMail1(t *testing.T) {
	host := "smtp.qq.com"
	port := 25
	username := "xxx@qq.com"
	password := "授权码"
	from := "xxx@qq.com"
	to := []string{"xxx@qq.com"}
	subject := "subject"
	content := []byte(fmt.Sprintf("mail-go test, current datetime %s", time.Now().Format("2006-01-02 15:04:05")))

	err := SendMail(host, port, username, password, from, to, subject, content)
	if err != nil {
		t.Error(err)
	}
}

// LoginAuth
// ## links
// - http://being23.github.io/2015/09/17/%E4%BD%BF%E7%94%A8golang%E5%8F%91%E9%80%81%E9%82%AE%E4%BB%B6/
// - https://github.com/go-gomail/gomail/issues/16
// - https://blog.csdn.net/u010918487/article/details/108473196
func TestSendMail2(t *testing.T) {
	host := "mail.xxx.com"
	port := 587
	username := "xxx@xxx.com"
	password := "xxx"
	from := "xxx@xxx.com"
	to := []string{"xxx@qq.com", "xxx@xxx.com"}
	subject := "subject"
	content := []byte(fmt.Sprintf("mail-go test, current datetime %s", time.Now().Format("2006-01-02 15:04:05")))

	// 参数
	addr := fmt.Sprintf("%s:%d", host, port)
	auth := NewLoginAuth(username, password)
	//msg := []byte("")
	line := "\r\n"
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("From: %s%s", from, line))
	buffer.WriteString(fmt.Sprintf("To: %s%s", strings.Join(to, ";"), line))
	buffer.WriteString(fmt.Sprintf("Subject: %s%s", subject, line))
	buffer.WriteString(fmt.Sprintf("Content-Type: %s;charset=UTF-8%s%s", "text/html", line, line))
	buffer.Write(content)

	// 发送邮件
	err := smtp.SendMail(addr, auth, from, to, buffer.Bytes())
	if err != nil {
		t.Logf(err.Error())
	}
}
