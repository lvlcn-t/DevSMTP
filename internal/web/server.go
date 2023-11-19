package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvlcn-t/DevSMTP/internal/config"
	"github.com/lvlcn-t/DevSMTP/internal/web/renderer"
)

type Server interface {
	Run(ctx context.Context) error
}

type server struct {
	r   *gin.Engine
	cfg *config.Config
}

func NewWebServer(cfg *config.Config) (Server, error) {
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
		gin.Recovery(),
	)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.HTMLRender = &renderer.TemplRender{}
	return &server{
		r:   r,
		cfg: cfg,
	}, nil
}

func (s *server) Run(ctx context.Context) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	return s.r.Run(fmt.Sprintf("%s:%d", s.cfg.IP, s.cfg.WebConfig.Port))
}
