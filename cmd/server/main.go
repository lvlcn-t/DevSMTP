package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/lvlcn-t/DevSMTP/internal/config"
	"github.com/lvlcn-t/DevSMTP/internal/services/smtp"
	"github.com/lvlcn-t/DevSMTP/internal/web"
	"github.com/lvlcn-t/halog/pkg/logger"
)

func main() {
	cfg := &config.Config{
		IP:        "0.0.0.0",
		Verbosity: false,
		SMTPConfig: &config.SMTP{
			Port:            1025,
			User:            "",
			Password:        "",
			MailDirectory:   "",
			LogMailContents: false,
		},
		WebConfig: &config.Web{
			Port:     1080,
			User:     "",
			Password: "",
			Disable:  false,
		},
	}

	var cfgPath string
	flag.StringVar(&cfgPath, "cfg-file", "/config/config.yaml", "Path for provided yaml config file")

	// Load configurations in order of precedence
	config.LoadEnvVars(cfg)
	config.LoadConfigFile(cfg, cfgPath)
	config.LoadCmdFlags(cfg)

	opt := &slog.HandlerOptions{Level: slog.LevelError}
	if cfg.Verbosity {
		opt = &slog.HandlerOptions{Level: slog.LevelInfo}
	}
	log := logger.NewNamedLogger("DevSMTP", slog.NewJSONHandler(os.Stderr, opt))
	ctx := logger.IntoContext(context.Background(), log)

	smtpSrv, err := smtp.NewSMTPServer(cfg)
	if err != nil {
		log.FatalContext(ctx, "Failed to create the smtp server", "error", err)
	}
	webSrv, err := web.NewWebServer(cfg)
	if err != nil {
		log.FatalContext(ctx, "Failed to create the web server", "error", err)
	}

	errors := make(chan error, 2)
	go func() {
		if err := smtpSrv.Run(); err != nil {
			log.ErrorContext(ctx, "Failed to start SMTP server", "error", err)
			errors <- err
		}
	}()

	go func() {
		if err := webSrv.Run(); err != nil {
			log.ErrorContext(ctx, "Failed to start Web server", "error", err)
			errors <- err
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errors:
		log.ErrorContext(ctx, "Error received from server", "error", err)
	case <-sigs:
		ctx.Done()
		log.InfoContext(ctx, "Shutdown signal received")
	}
}
