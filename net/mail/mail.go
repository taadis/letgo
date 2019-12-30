package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
)

// 发送邮件
func SendMail(host string, port int, username string, password string, from string, to []string, subject string, content []byte) (err error) {
	// 参数
	addr := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth("text/plain", username, password, host)
	//msg := []byte("")
	line := "\r\n"
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("From: %s%s", from, line))
	buffer.WriteString(fmt.Sprintf("To: %s%s", strings.Join(to, ";"), line))
	buffer.WriteString(fmt.Sprintf("Subject: %s%s", subject, line))
	buffer.WriteString(fmt.Sprintf("Content-Type: %s;charset=UTF-8%s%s", "text/html", line, line))
	buffer.Write(content)

	// 发送邮件
	err = smtp.SendMail(addr, auth, from, to, buffer.Bytes())
	if err != nil {
		return
	}
	return
}
