package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) genRecordResp(uid int64, page int) (string, *gotgbot.InlineKeyboardMarkup, error) {
	recordCount, err := tg.db.GetRecordCount(uid)
	if err != nil {
		return "", nil, err
	}

	if recordCount == 0 {
		return "未找到记录", &gotgbot.InlineKeyboardMarkup{InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{
					Text:         "查询 UID",
					CallbackData: fmt.Sprintf("record_uid_%d", uid),
				},
			},
		}}, nil
	}

	record, err := tg.db.GetRecord(uid, page-1)
	if err != nil {
		tg.sugar.Errorf("GetRecord: %v", err)
		return "", nil, err
	}

	text := fmt.Sprintf("UID: `%d`\n描述:\n%s", uid, EscapeMarkdownV2(record.Description))

	inlineKeyboard := [][]gotgbot.InlineKeyboardButton{}
	if record.ChatID.Valid && record.MessageID.Valid {
		chatIdBot := record.ChatID.Int64
		chatIdBotStr := strconv.FormatInt(chatIdBot, 10)
		chatIdStr := strings.TrimPrefix(chatIdBotStr, "-100")
		inlineKeyboard = append(inlineKeyboard, []gotgbot.InlineKeyboardButton{
			{
				Text: "查看证据",
				Url:  fmt.Sprintf("https://t.me/c/%s/%d", chatIdStr, record.MessageID.Int64),
			},
		})
	}
	keyboard := []gotgbot.InlineKeyboardButton{}
	// prev
	callbackData := "nil"
	if page > 1 {
		callbackData = fmt.Sprintf("record_page_%d_%d", uid, page-1)
	}
	keyboard = append(keyboard, gotgbot.InlineKeyboardButton{
		Text:         "上一页",
		CallbackData: callbackData,
	})
	// middle
	keyboard = append(keyboard,
		gotgbot.InlineKeyboardButton{
			Text:         fmt.Sprintf("%d / %d", page, recordCount),
			CallbackData: "nil",
		},
	)
	// next
	callbackData = "nil"
	if page < int(recordCount) {
		callbackData = fmt.Sprintf("record_page_%d_%d", uid, page+1)
	}
	keyboard = append(keyboard, gotgbot.InlineKeyboardButton{
		Text:         "下一页",
		CallbackData: callbackData,
	})

	inlineKeyboard = append(inlineKeyboard, keyboard)

	return text, &gotgbot.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard}, nil
}

func (tg *TelegramBot) commandRecord(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	splits := strings.Split(msg.Text, " ")
	if len(splits) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	idStr := splits[1]
	uid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	text, replyMarkup, err := tg.genRecordResp(uid, 1)
	if err != nil {
		tg.sugar.Errorf("failed to generate record response: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "查询失败", nil)
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode:             "MarkdownV2",
		DisableWebPagePreview: true,
		ReplyMarkup:           replyMarkup,
	})
	if err != nil {
		return err
	}

	return nil
}

func (tg *TelegramBot) callbackRecord(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, "record_")
}

func (tg *TelegramBot) callbackRecordResp(b *gotgbot.Bot, ctx *ext.Context) error {
	callbackData := ctx.CallbackQuery.Data
	switch {
	case strings.HasPrefix(callbackData, "record_uid_"):
		uidStr := strings.TrimPrefix(callbackData, "record_uid_")
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "参数错误",
			})
			return err
		}
		isAdmin := IsLevelAdmin(tg.GetUserAdminLevel(ctx.CallbackQuery.From.Id))
		text, replyMarkup, err := tg.genUidResp(uid, true, isAdmin)
		if err != nil {
			tg.sugar.Errorf("failed to generate uid response: %v", err)
			_, err := ctx.EffectiveMessage.Reply(b, "查询失败", nil)
			return err
		}
		editMessageTextOpts := &gotgbot.EditMessageTextOpts{
			ParseMode:             "MarkdownV2",
			DisableWebPagePreview: true,
		}
		if replyMarkup != nil {
			editMessageTextOpts.ReplyMarkup = *replyMarkup
		}
		_, _, err = ctx.EffectiveMessage.EditText(b, text, editMessageTextOpts)
		return err
	case strings.HasPrefix(callbackData, "record_page_"):
		data := strings.TrimPrefix(callbackData, "record_page_")
		splits := strings.Split(data, "_")
		uidStr := splits[0]
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "参数错误",
			})
			return err
		}
		pageStr := splits[1]
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "参数错误",
			})
			return err
		}
		text, replyMarkup, err := tg.genRecordResp(uid, page)
		if err != nil {
			tg.sugar.Errorf("failed to generate record response: %v", err)
			_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "查询错误",
			})
			return err
		}
		editMessageTextOpts := &gotgbot.EditMessageTextOpts{
			ParseMode:             "MarkdownV2",
			DisableWebPagePreview: true,
		}
		if replyMarkup != nil {
			editMessageTextOpts.ReplyMarkup = *replyMarkup
		}
		_, _, err = ctx.EffectiveMessage.EditText(b, text, editMessageTextOpts)
		return err
	}
	return nil
}
