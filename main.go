package main

import (
	"biliroaming-blacklist-server-go/bot"
	"biliroaming-blacklist-server-go/config"
	"biliroaming-blacklist-server-go/db"
	"biliroaming-blacklist-server-go/web"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.New(config.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}
	tg, err := bot.New(db, config.TelegramConfig)
	if err != nil {
		log.Fatal(err)
	}
	web.New(db, tg, config.CaptchasConfig)
}
