package email

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(content,email string)error{
	d := gomail.NewDialer("smtp.163.com",25,"y484742285@163.com","YMZDBQFXSWIRXSQR")
	//YMZDBQFXSWIRXSQR
	m := gomail.NewMessage()
	m.SetAddressHeader("From","y484742285@163.com","yinqingping")
	m.SetHeader("To",email)
	m.SetHeader("Subject", "通知")
	m.SetBody("text/html",content)

	if err := d.DialAndSend(m);err != nil {
		return err
	}
	return nil
}

