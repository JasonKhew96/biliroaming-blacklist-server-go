package web

import (
	"biliroaming-blacklist-server-go/bot"
	"biliroaming-blacklist-server-go/config"
	"biliroaming-blacklist-server-go/db"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"go.uber.org/zap"
)

const (
	StatusNone int8 = iota
	StatusBlacklist
	StatusWhitelist
)

type Web struct {
	sugar *zap.SugaredLogger

	app *fiber.App
	db  *db.Database
	ctx context.Context
	tg  *bot.TelegramBot

	config config.CaptchasConfig
	auth   string
}

func (w *Web) index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "首页",
	}, "layouts/main")
}

func New(db *db.Database, tg *bot.TelegramBot, port int, conf config.CaptchasConfig, auth string, sugar *zap.SugaredLogger) {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:        engine,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	})

	web := Web{
		sugar:  sugar,
		db:     db,
		ctx:    context.Background(),
		app:    app,
		tg:     tg,
		config: conf,
		auth:   auth,
	}

	app.Get("/", web.index)
	app.Get("/api/users/:uid", web.usersUid)
	app.Get("/status.php", web.usersUid)
	app.Get("/check", web.checkGet)
	app.Post("/check", web.checkPost)
	app.Get("/report", web.reportGet)
	app.Post("/report", web.reportPost)

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		sugar.Fatalln(err)
	}
}
