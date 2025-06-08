package emailSend

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	SmtpAuthorAddress = "smtp.gmail.com"
	SmtpSeverService  = "smtp.gmail.com:587"
)

type Sender interface {
	SendEmail(
		Title string,
		Content string,
		To []string,
		Cc []string,
		Bcc []string,
		AttachFile []string,
	) error
}
type GmailSender struct {
	name              string
	fromEmailAddr     string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddr string, fromEmailPassword string) Sender {
	return &GmailSender{
		name:              name,
		fromEmailAddr:     fromEmailAddr,
		fromEmailPassword: fromEmailPassword,
	}
}
func (sender *GmailSender) SendEmail(
	Title string,
	Content string,
	To []string,
	Cc []string,
	Bcc []string,
	AttachFile []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("<%s>", sender.fromEmailAddr)
	e.To = To
	e.Subject = Title
	e.HTML = []byte(Content)
	e.Bcc = Bcc
	e.Cc = Cc
	for _, f := range AttachFile {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s:%w", f, err)
		}

	}
	auPath := smtp.PlainAuth("", sender.fromEmailAddr, sender.fromEmailPassword, SmtpAuthorAddress)
	return e.Send(SmtpSeverService, auPath)
}
