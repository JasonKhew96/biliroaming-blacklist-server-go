package web

import (
	"biliroaming-blacklist-server-go/bot"
	"biliroaming-blacklist-server-go/db"
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/template/html"
)

const (
	StatusNone int8 = iota
	StatusBlacklist
	StatusWhitelist
)

type Web struct {
	app *fiber.App
	db  *db.Database
	ctx context.Context
	tg  *bot.TelegramBot
}

func (w *Web) index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "首页",
	}, "layouts/main")
}

func New(db *db.Database, tg *bot.TelegramBot) {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(csrf.New(
		csrf.Config{
			Next: func(c *fiber.Ctx) bool {
				return strings.HasPrefix(c.Path(), "/api/") || strings.HasPrefix(c.Path(), "/status.php")
			},
		},
	))

	web := Web{
		db:  db,
		ctx: context.Background(),
		app: app,
		tg:  tg,
	}

	app.Get("/", web.index)
	app.Get("/api/users/:uid", web.usersUid)
	app.Get("/status.php", web.usersUid)
	app.Get("/check", web.checkGet)
	app.Post("/check", web.checkPost)
	app.Get("/report", web.reportGet)
	app.Post("/report", web.reportPost)

	if err := app.Listen(":7181"); err != nil {
		log.Fatal(err)
	}
}
