package api

import (
	"go-email/pkg/app"
	"go-email/pkg/e"
	"go-email/service"
	"go-email/tmpl"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SendMailForm struct {
	ToAddress  string `form:"toAddress" valid:"Email; MaxSize(100)"`
	ToName     string `form:"toName" valid:"Required;MaxSize(100)"`
	TemplateId string `form:"templateId" valid:"Required;MaxSize(10)"`
	Word       string `form:"word"`
}

//发送邮件
func SendMail(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form SendMailForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//检查是否支持此模板
	flag := tmpl.CheckTemplate(form.TemplateId)
	if !flag {
		appG.Response(httpCode, e.INVALID_TEMPLATE, nil)
		return
	}
	//获取模板
	templateMsg := tmpl.GetTemplateV1(form.TemplateId, form.Word)

	//发送邮件
	mail := service.Mail{
		Subject:   templateMsg.Subject,
		ToAddress: form.ToAddress,
		ToName:    form.ToName,
		Body:      templateMsg.Body,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	}
	mail.SendMsg()
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
