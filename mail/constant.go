package mail

import "net/smtp"

type email struct {
	subject string
	body    struct {
		content []byte
		types   string
	}
	attachment map[string]string
	to         []string
	cc         []string
	bcc        []string
	server     *server
}
type EmailPara struct {
	Subject    string
	Content    []byte
	Types      string
	Attachment map[string]string
	To         []string
	Cc         []string
	Bcc        []string
	User       string
	Password   string
	Host       string
	Port       string
}
type server struct {
	user     string
	password string
	host     string
	port     string
	auth     *smtp.Auth
}
