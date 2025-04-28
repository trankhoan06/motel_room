package emailSend

import (
	"fmt"
	"main.go/config"
)

func SendVerifyEmail(email string, code int, cfg *config.Config) error {
	sender := NewGmailSender()
	title := "Verify Email"
	content := fmt.Sprintf(`
<h1>Hello</h1>
<p>Web send code: %d</p>
`, code)
	to := []string{email}
	err := sender.SendEmail(cfg, title, content, to, nil, nil, nil)
	return err
}
