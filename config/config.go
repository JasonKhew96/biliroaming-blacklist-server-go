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
}

type Config struct {
	DatabaseConfig `yaml:"database"`
	TelegramConfig `yaml:"telegram"`
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
