package mail

import (
	"log"
	"net/smtp"
	"os"

	"github.com/darcyjoven/util"
)

func NewMail(m *EmailPara) *email {
	var e email
	if m.Subject != "" {
		e.subject = m.Subject
	}
	if m.Content != nil {
		e.body.content = m.Content
	}
	if m.Types != "" {
		e.body.types = m.Types
	}
	if m.Attachment != nil {
		e.attachment = m.Attachment
	}
	if m.To != nil {
		e.to = m.To
	}
	if m.Cc != nil {
		e.cc = m.Cc
	}
	if m.Bcc != nil {
		e.bcc = m.Bcc
	}
	//  server
	e.server = &server{
		user:     m.User,
		password: m.Password,
		host:     m.Host,
		port:     m.Port,
	}
	return &e
}
func (e *email) SetSubject(sub string) {
	e.subject = sub
}

// body 支持markdown html 和text
// types markdown 1
func (e *email) SetBody(body []byte, types int) {
	switch types {
	case util.Markdown:
		e.body.content = util.MarkTotHtml(body)
		e.body.types = "text/html;charset=utf-8"
	case util.Html:
		e.body.content = body
		e.body.types = "text/html;charset=utf-8"
	case util.Text:
		e.body.content = body
		e.body.types = "text/plain;charset=utf-8"
	}
}
func (e *email) SetAttach(att map[string]string) error {
	for k, v := range att {
		_, err := os.Stat(v)
		if err != nil {
			log.Println(err)
			return err
		}
		e.attachment[k] = att[k]
	}
	return nil
}
func (e *email) SetTo(to []string) {
	e.to = append(e.to, to...)
}
func (e *email) SetCc(cc []string) {
	e.cc = append(e.cc, cc...)
}
func (e *email) SetBcc(bcc []string) {
	e.bcc = append(e.bcc, bcc...)
}
func (e *email) AddTo(to ...string) {
	e.to = append(e.to, to...)
}
func (e *email) AddCc(cc ...string) {
	e.cc = append(e.cc, cc...)
}
func (e *email) AddBcc(bcc ...string) {
	e.bcc = append(e.bcc, bcc...)
}
func (e *email) AddAttach(att map[string]string) error {
	for k, v := range att {
		_, err := os.Stat(v)
		if err != nil {
			log.Println(err)
			return err
		}
		e.attachment[k] = v
	}
	return nil
}

func (e *email) Send() error {
	// 验证服务器
	e.server.verify()
	// 设置正文
	e.setBody()
	return smtp.SendMail(
		e.server.host+":"+e.server.port,
		*e.server.auth,
		e.server.user,
		e.to,
		e.body.content,
	)
}

func (e *email) SetServer(s struct {
	user     string
	password string
	host     string
	port     string
}) {
	if s.user != "" {
		e.server.user = s.user
	}
	if s.password != "" {
		e.server.password = s.password
	}
	if s.host != "" {
		e.server.host = s.host
	}
	if s.port != "" {
		e.server.port = s.port
	}
	//  auth
	e.server.verify()
}
