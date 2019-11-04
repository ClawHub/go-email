package mail

import (
	"crypto/tls"
	"go-email/pkg/setting"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	Username string
	Password string
	Host     string
	Port     int
}

var EMailSetting = &Mail{}

func Setup() {
	setting.MapTo("mail", EMailSetting)
}
func Send(subject, toAddress, toName, body string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "clawhub@163.com", "ClawHub")
	m.SetAddressHeader("To", toAddress, toName)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(EMailSetting.Host, EMailSetting.Port, EMailSetting.Username, EMailSetting.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
