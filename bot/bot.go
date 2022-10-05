package bot

import (
	"biliroaming-blacklist-server-go/config"
	"biliroaming-blacklist-server-go/db"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

type TelegramBot struct {
	db            *db.Database
	Bot           *gotgbot.Bot
	AnnouceChatId int64
	ReportChatId  int64
}

func New(db *db.Database, config config.TelegramConfig) (*TelegramBot, error) {
	bot, err := gotgbot.NewBot(config.BotToken, nil)
	if err != nil {
		return nil, err
	}

	tg := &TelegramBot{
		db:            db,
		Bot:           bot,
		AnnouceChatId: config.AnnounceChatId,
		ReportChatId:  config.ReportChatId,
	}

	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update: ", err.Error())
				return ext.DispatcherActionNoop
			},
		},
	})
	dispatcher := updater.Dispatcher

	// super admin
	dispatcher.AddHandler(handlers.NewCommand("addadmin", tg.commandAddAdmin))
	dispatcher.AddHandler(handlers.NewCommand("removeadmin", tg.commandRemoveAdmin))
	dispatcher.AddHandler(handlers.NewCommand("alteradmin", tg.commandAlterAdmin))

	// admin
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

	// dispatcher.AddHandler(
	// 	handlers.NewInlineQuery(
	// 		func(iq *gotgbot.InlineQuery) bool {
	// 			if _, err := strconv.ParseInt(iq.Query, 10, 64); err != nil {
	// 				return false
	// 			}
	// 			return true
	// 		},
	// 		func(b *gotgbot.Bot, ctx *ext.Context) error {
	// 			// b.AnswerInlineQuery(ctx.InlineQuery.Id, nil)
	// 			return nil
	// 		},
	// 	),
	// )

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 60,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 60,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s has been started...\n", bot.User.Username)

	// updater.Idle()

	return tg, nil
}
