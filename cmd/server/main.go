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

}
