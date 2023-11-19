package smtp

import (
	"context"
	"fmt"

	"github.com/emersion/go-smtp"
	"github.com/lvlcn-t/DevSMTP/internal/config"
	"github.com/lvlcn-t/DevSMTP/internal/services/smtp/backend"
)

type Server interface {
	Run(ctx context.Context) error
}

type server struct {
	srv *smtp.Server
	be  backend.Backend
}

func NewSMTPServer(cfg *config.Config) (Server, error) {
	be := backend.NewBackend(cfg)
	srv := smtp.NewServer(be)

	srv.Addr = fmt.Sprintf("%s:%d", cfg.IP, cfg.SMTPConfig.Port)
	srv.AllowInsecureAuth = true

	return &server{
		srv: srv,
		be:  be,
	}, nil
}

func (s *server) Run(ctx context.Context) error {
	s.be.SetContext(ctx)
	return s.srv.ListenAndServe()
}
