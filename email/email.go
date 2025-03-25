package emailSend

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	SenderEmailAccount  = "trankhoan06@gmail.com"
	PassWordSenderEmail = "klptazdfrrcwvgfp"
	NameSender          = "Web seller"
	SmtpAuthorAddress   = "smtp.gmail.com"
	SmtpSeverService    = "smtp.gmail.com:587"
)

type Sender interface {
	SendEmail(
		Title string,
		Content string,
		To []string,
		Cc []string,
		Bcc []string,
		AttactFile []string,
	) error
}
type GmailSender struct {
	Name                   string
	FromEmailAddress       string
	FromEmailPasswordEmail string
}

func NewGmailSender() Sender {
	return &GmailSender{
		Name:                   NameSender,
		FromEmailAddress:       SenderEmailAccount,
		FromEmailPasswordEmail: PassWordSenderEmail,
	}
}
func (sender *GmailSender) SendEmail(
	Title string,
	Content string,
	To []string,
	Cc []string,
	Bcc []string,
	AttactFile []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", NameSender, SenderEmailAccount)
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
	auPath := smtp.PlainAuth("", sender.FromEmailAddress, sender.FromEmailPasswordEmail, SmtpAuthorAddress)
	return e.Send(SmtpSeverService, auPath)
}
