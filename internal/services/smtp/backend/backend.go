package backend

import (
	"errors"
	"io"

	"github.com/emersion/go-smtp"
	"github.com/lvlcn-t/DevSMTP/internal/config"
)

type backend struct {
	*config.SMTP
}

func NewBackend(cfg *config.Config) backend {
	return backend{
		cfg.SMTPConfig,
	}
}

func (b backend) NewSession(conn *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

func (b backend) Login(conn *smtp.Conn, username, password string) (smtp.Session, error) {
	if username != b.User || password != b.Password {
		return nil, errors.New("invalid credentials")
	}
	return &Session{}, nil
}

func (b backend) AnonymousLogin(conn *smtp.Conn) (smtp.Session, error) {
	// TODO: implement allowed anonymous login if credentials aren't configured
	return nil, smtp.ErrAuthRequired
}

type Session struct{}

func (s *Session) AuthPlain(username, password string) error {
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	return nil
}

func (s *Session) Data(r io.Reader) error {
	// TODO: Implement the logic to handle the received data
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
