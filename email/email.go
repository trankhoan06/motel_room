package emailSend

import (
	"fmt"
	"github.com/jordan-wright/email"
	"main.go/config"
	"net/smtp"
)

const (
	SmtpAuthorAddress = "smtp.gmail.com"
	SmtpSeverService  = "smtp.gmail.com:587"
)

type Sender interface {
	SendEmail(
		cfg *config.Config,
		Title string,
		Content string,
		To []string,
		Cc []string,
		Bcc []string,
		AttactFile []string,
	) error
}
type GmailSender struct {
}

func NewGmailSender() Sender {
	return &GmailSender{}
}
func (sender *GmailSender) SendEmail(
	cfg *config.Config,
	Title string,
	Content string,
	To []string,
	Cc []string,
	Bcc []string,
	AttactFile []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress)
	e.To = To
	e.Subject = Title
	e.HTML = []byte(Content)
	e.Bcc = Bcc
	e.Cc = Cc
	for _, f := range AttactFile {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s:%w", f, err)
		}

	}
	auPath := smtp.PlainAuth("", cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword, SmtpAuthorAddress)
	return e.Send(SmtpSeverService, auPath)
}
