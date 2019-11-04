package service

import (
	"go-email/mail"
	"go-email/pkg/gmongo"
	"go-email/pkg/logging"
	"go.uber.org/zap"
)

type Mail struct {
	Subject   string
	ToAddress string
	ToName    string
	Body      string
	Time      string
	Res       string
}

//发送邮件
func (m *Mail) SendMsg() {
	//发送邮件
	err := mail.Send(m.Subject, m.ToAddress, m.ToName, m.Body)
	if err != nil {
		m.Res = err.Error()
		logging.AppLogger.Error("send mail fail", zap.Error(err))
	}
	//数据入mongo库
	collection := gmongo.Client.Database("email").Collection("email")
	_, err = collection.InsertOne(gmongo.GetContext(), m)
	if err != nil {
		logging.AppLogger.Error("insert mongo mail fail", zap.Error(err))
	}
}
