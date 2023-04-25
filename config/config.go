package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type TelegramConfig struct {
	BotToken       string `yaml:"bot_token"`
	AnnounceChatId int64  `yaml:"announce_chat_id"`
	ReportChatId   int64  `yaml:"report_chat_id"`
	DefaultBanTime string `yaml:"default_ban_time"`
}

type CaptchasConfig struct {
	SiteKey   string `yaml:"site_key"`
	SecretKey string `yaml:"secret_key"`
}

type Config struct {
	Port           int    `yaml:"port"`
	Auth           string `yaml:"auth"`
	DatabaseConfig `yaml:"database"`
	TelegramConfig `yaml:"telegram"`
	CaptchasConfig `yaml:"captchas"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
