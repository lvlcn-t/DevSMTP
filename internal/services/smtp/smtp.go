package smtp

import (
	"fmt"

	"github.com/emersion/go-smtp"
	"github.com/lvlcn-t/DevSMTP/internal/config"
	"github.com/lvlcn-t/DevSMTP/internal/services/smtp/backend"
)

type Server interface {
	Run() error
}

type server struct {
	srv *smtp.Server
}

func NewSMTPServer(cfg *config.Config) (Server, error) {
	srv := smtp.NewServer(backend.NewBackend(cfg))

	srv.Addr = fmt.Sprintf("%s:%d", cfg.IP, cfg.SMTPConfig.Port)
	srv.AllowInsecureAuth = true

	return &server{
		srv: srv,
	}, nil
}

func (s *server) Run() error {
	return s.srv.ListenAndServe()
}
