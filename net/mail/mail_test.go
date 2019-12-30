package mail

import (
	"fmt"
	"testing"
	"time"
)

//
func TestSendMail1(t *testing.T) {
	host := "smtp.qq.com"
	port := 25
	username := ""
	password := ""
	from := ""
	to := []string{ "" }
	subject := "subject"
	content := []byte(fmt.Sprintf("mail-go test, current datetime %s", time.Now().Format("2006-01-02 15:04:05")))

	err := SendMail(host, port, username, password, from, to, subject, content)
	if err != nil{
		t.Error(err)
	}
}
