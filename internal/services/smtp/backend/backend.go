package backend

import (
	"context"
	"sync"

	"github.com/emersion/go-smtp"
	"github.com/lvlcn-t/DevSMTP/internal/config"
	"github.com/lvlcn-t/DevSMTP/internal/services/database"
)

type Backend interface {
	smtp.Backend
	SetContext(ctx context.Context)
}

type backend struct {
	cfg *config.SMTP
	ctx context.Context
	db  database.Database
	mu  sync.Mutex
}

func NewBackend(cfg *config.Config) Backend {
	return &backend{
		cfg: cfg.SMTPConfig,
		db:  database.NewDatabase(),
		mu:  sync.Mutex{},
	}
}

func (b *backend) SetContext(ctx context.Context) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.ctx = ctx
}
