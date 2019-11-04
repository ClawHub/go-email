# go-email
Go版本邮件发送服务
###简单的画了个架构图：
![邮件服务架构.png](https://upload-images.jianshu.io/upload_images/8803909-eb8914cc92612944.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 发送邮件
- 记录mongo
- 日志系统
- 配置文件读取
- REST服务
- vendor包管理

###项目代码结构
![image.png](https://upload-images.jianshu.io/upload_images/8803909-345664ab91a412fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

比较简单的项目结构。

###技术核心：发送邮件
```
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
```
###业务核心：发送邮件
```
//发送邮件
func (m *Mail) Send() {
	//处理模板，TODO
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
```

###github地址
[https://github.com/ClawHub/go-email](https://github.com/ClawHub/go-email)
