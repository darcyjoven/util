package mail

import "net/smtp"

func (s *server) verify() {
	a := smtp.PlainAuth("", s.user, s.password, s.host)
	(s.auth) = &a
}
