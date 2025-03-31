package emailSend

import "fmt"

func SendForgotPassword(email string, code int) error {
	sender := NewGmailSender()
	title := "FORGOT PASSWORD"
	content := fmt.Sprintf(`
<h1>Hello</h1>
<p>Web send code: %d</p>
`, code)
	to := []string{email}
	err := sender.SendEmail(title, content, to, nil, nil, nil)
	return err
}
