package emailSend

//
//import (
//	"fmt"
//	"main.go/config"
//)
//
//func SendForgotPassword(email string, code int, cfg *config.Config) error {
//	sender := NewGmailSender()
//	title := "FORGOT PASSWORD"
//	content := fmt.Sprintf(`
//<h1>Hello</h1>
//<p>Web send code: %d</p>
//`, code)
//	to := []string{email}
//	err := sender.SendEmail(cfg, title, content, to, nil, nil, nil)
//	return err
//}
