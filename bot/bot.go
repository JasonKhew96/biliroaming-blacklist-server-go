package bot

import (
	"biliroaming-blacklist-server-go/config"
	"biliroaming-blacklist-server-go/db"
	"fmt"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"go.uber.org/zap"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

type TelegramBot struct {
	sugar *zap.SugaredLogger

	db            *db.Database
	Bot           *gotgbot.Bot
	AnnouceChatId int64
	ReportChatId  int64

	DefaultBanTime string
}

func New(db *db.Database, config config.TelegramConfig, sugar *zap.SugaredLogger) (*TelegramBot, error) {
	bot, err := gotgbot.NewBot(config.BotToken, nil)
	if err != nil {
		return nil, err
	}

	tg := &TelegramBot{
		sugar:          sugar,
		db:             db,
		Bot:            bot,
		AnnouceChatId:  config.AnnounceChatId,
		ReportChatId:   config.ReportChatId,
		DefaultBanTime: config.DefaultBanTime,
	}

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("help", tg.commandHelp))

	// super admin
	dispatcher.AddHandler(handlers.NewCommand("addadmin", tg.commandAddAdmin))
	dispatcher.AddHandler(handlers.NewCommand("removeadmin", tg.commandRemoveAdmin))
	dispatcher.AddHandler(handlers.NewCommand("alteradmin", tg.commandAlterAdmin))

	// admin
	dispatcher.AddHandler(handlers.NewCommand("key", tg.commandKey))

	dispatcher.AddHandler(handlers.NewCommand("ban", tg.commandBan))
	dispatcher.AddHandler(handlers.NewCommand("unban", tg.commandUnban))

	dispatcher.AddHandler(handlers.NewCommand("white", tg.commandWhite))
	dispatcher.AddHandler(handlers.NewCommand("unwhite", tg.commandUnwhite))

	dispatcher.AddHandler(handlers.NewCommand("stats", tg.commandStats))

	// user
	dispatcher.AddHandler(handlers.NewCommand("uid", tg.commandUid))
	dispatcher.AddHandler(handlers.NewCommand("record", tg.commandRecord))
	dispatcher.AddHandler(handlers.NewCommand("report", tg.commandReport))

	// callback
	dispatcher.AddHandler(handlers.NewCallback(tg.callbackUid, tg.callbackUidResp))
	dispatcher.AddHandler(handlers.NewCallback(tg.callbackRecord, tg.callbackRecordResp))
	dispatcher.AddHandler(handlers.NewCallback(tg.callbackReport, tg.callbackReportResp))

	// nil
	dispatcher.AddHandler(handlers.NewCallback(func(cq *gotgbot.CallbackQuery) bool {
		return strings.HasPrefix(cq.Data, "nil")
	}, func(b *gotgbot.Bot, ctx *ext.Context) error {
		_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			CacheTime: 300,
		})
		return err
	}))

	dispatcher.AddHandler(handlers.NewInlineQuery(tg.inlineQuery, tg.inlineQueryResp))

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 60,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 60,
			},
			AllowedUpdates: []string{"message", "callback_query", "inline_query"},
		},
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s has been started...\n", bot.User.Username)

	// updater.Idle()

	return tg, nil
}
