package config

import (
	"flag"
	"os"

	"github.com/lvlcn-t/DevSMTP/pkg/utils"
	"gopkg.in/yaml.v3"
)

func LoadCmdFlags(cfg *Config) {
	flag.StringVar(&cfg.IP, "ip", "0.0.0.0", "IP address for binding")
	flag.BoolVar(&cfg.Verbosity, "v", false, "Verbosity")

	flag.IntVar(&cfg.SMTPConfig.Port, "s", 1025, "SMTP port")
	flag.IntVar(&cfg.SMTPConfig.Port, "smtp", 1025, "SMTP port")
	flag.StringVar(&cfg.SMTPConfig.User, "smtp-user", "", "SMTP user")
	flag.StringVar(&cfg.SMTPConfig.Password, "smtp-password", "", "SMTP password")
	flag.StringVar(&cfg.SMTPConfig.MailDirectory, "mail-directory", "", "Path for persisting mails")
	flag.BoolVar(&cfg.SMTPConfig.LogMailContents, "log-mail-contents", false, "Log mail contents")

	flag.IntVar(&cfg.WebConfig.Port, "w", 1080, "Web port")
	flag.IntVar(&cfg.WebConfig.Port, "web", 1080, "Web port")
	flag.StringVar(&cfg.WebConfig.User, "web-user", "", "Web user")
	flag.StringVar(&cfg.WebConfig.Password, "web-password", "", "Web password")
	flag.BoolVar(&cfg.WebConfig.Disable, "disable-web", false, "Disable web interface")

	flag.Parse()
}

func LoadConfigFile(cfg *Config, filePath string) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		return
	}
}

func LoadEnvVars(cfg *Config) {
	ip, err := utils.GetEnv("DEVSMTP_IP", utils.ParseString)
	if err != nil {
		ip = cfg.IP
	}
	v, err := utils.GetEnv("DEVSMTP_VERBOSITY", utils.ParseBool)
	if err != nil {
		v = cfg.Verbosity
	}

	sp, err := utils.GetEnv("DEVSMTP_SMTP_PORT", utils.ParseInt)
	if err != nil {
		sp = cfg.SMTPConfig.Port
	}
	su, err := utils.GetEnv("DEVSMTP_SMTP_USER", utils.ParseString)
	if err != nil {
		su = cfg.SMTPConfig.User
	}
	sPass, err := utils.GetEnv("DEVSMTP_SMTP_PASSWORD", utils.ParseString)
	if err != nil {
		sPass = cfg.SMTPConfig.Password
	}
	md, err := utils.GetEnv("DEVSMTP_MAIL_DIRECTORY", utils.ParseString)
	if err != nil {
		md = cfg.SMTPConfig.MailDirectory
	}
	logMc, err := utils.GetEnv("DEVSMTP_LOG_MAIL_CONTENTS", utils.ParseBool)
	if err != nil {
		logMc = cfg.SMTPConfig.LogMailContents
	}

	wp, err := utils.GetEnv("DEVSMTP_WEB_PORT", utils.ParseInt)
	if err != nil {
		wp = cfg.WebConfig.Port
	}
	wu, err := utils.GetEnv("DEVSMTP_WEB_USER", utils.ParseString)
	if err != nil {
		wu = cfg.WebConfig.User
	}
	wPass, err := utils.GetEnv("DEVSMTP_WEB_PASSWORD", utils.ParseString)
	if err != nil {
		wPass = cfg.WebConfig.Password
	}
	disWeb, err := utils.GetEnv("DEVSMTP_DISABLE_WEB", utils.ParseBool)
	if err != nil {
		disWeb = cfg.WebConfig.Disable
	}

	cfg = &Config{
		IP:        ip,
		Verbosity: v,
		SMTPConfig: &SMTP{
			Port:            sp,
			User:            su,
			Password:        sPass,
			MailDirectory:   md,
			LogMailContents: logMc,
		},
		WebConfig: &Web{
			Port:     wp,
			User:     wu,
			Password: wPass,
			Disable:  disWeb,
		},
	}
}
