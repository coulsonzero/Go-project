package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"net/textproto"
)

// 需要开启smtp：163邮箱-邮箱中心-开启smtp-获取授权码

func main() {
	emailExample()
}

func emailExample() {
	userEmail := "****@163.com"
	smtpAuthCode := "WOTXXVGYHYIPFBAM"
	emailType := "163"

	e := &email.Email{
		To:      []string{"***@163.com", "***@qq.com"},
		From:    userEmail,
		Subject: "Email Send Test",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h3>This a test email</h3>"),
		Headers: textproto.MIMEHeader{},
	}

	ch := make(chan int, 1)
	go func() {
		err := e.Send("smtp."+emailType+".com:25", smtp.PlainAuth("", userEmail, smtpAuthCode, "smtp."+emailType+".com"))
		if err != nil {
			log.Fatal(err)
		}
		ch <- 1
	}()

	<-ch
	fmt.Println("end")
}
