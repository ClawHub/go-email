package tmpl

import (
	"fmt"
	"go-email/tmpl/app"
)

//模板信息
type TemplateMsg struct {
	Subject, Body string
}

//模板
var Templates = make(map[string]map[string]string)

//初始化
func init() {
	//后期可改为读取配置的
	//旅行记账app-重置密码
	Templates["0001"] = app.TAreset
	//旅行记账app-注册
	Templates["0002"] = app.TAsignUp

	//一键分发-重置密码
	Templates["0003"] = app.Dreset
	//一键分发-注册
	Templates["0004"] = app.DsignUp
}

//判断是否支持此模板
func CheckTemplate(templateId string) bool {
	_, ok := Templates[templateId]
	if ok {
		return true
	}
	return false
}

//获取模板，目前支持出传入一位模板值
func GetTemplateV1(templateId string, word string) TemplateMsg {
	temp := Templates[templateId]
	body := fmt.Sprintf(temp["word"], word)
	return TemplateMsg{temp["subject"], body}
}
