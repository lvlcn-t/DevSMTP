package config

type Config struct {
	IP        string `yaml:"ip"`
	Verbosity bool   `yaml:"verbosity"`

	SMTPConfig *SMTP `yaml:"smtp"`
	WebConfig  *Web  `yaml:"web"`
}

type SMTP struct {
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	MailDirectory   string `yaml:"mailDirectory"`
	LogMailContents bool   `yaml:"logMailContents"`
}

type Web struct {
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Disable  bool   `yaml:"disable"`
}
