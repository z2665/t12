package models

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 *    user : example@example.com login smtp server user
 *    password: xxxxx login smtp server password
 *    host: smtp.example.com:port   smtp.163.com:25
 *    to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

func sendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func ForGotSend(to string, vcode string) {
	user := "WeNedFound@163.com"
	password := "TimeToFound"
	host := "smtp.163.com:25"

	subject := "Verification account From Found"

	body := `
    <html>
    <body>
    <h3>
    {{.}}
    </h3>
    </body>
    </html>
    `
	var buf bytes.Buffer
	t := template.New("email")
	t, _ = t.Parse(body)
	t.Execute(&buf, vcode)
	sendMail(user, password, host, to, subject, buf.String(), "html")

}

func MakeVcode(uname string) string {
	p := uname + time.Now().Format("2006-01-02 15:04:05") + strconv.Itoa(os.Getpid())
	md := md5.New()
	md.Write([]byte(p))
	sum := md.Sum(nil)
	return hex.EncodeToString(sum)
}
