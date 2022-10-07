package bot

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func genUidInlineKeyboard(uid int64, isBlacklist, isAdmin bool) *gotgbot.InlineKeyboardMarkup {
	inlineKeyboard := [][]gotgbot.InlineKeyboardButton{}
	inlineKeyboard01 := []gotgbot.InlineKeyboardButton{
		{
			Text:         "查询记录",
			CallbackData: fmt.Sprintf("uid_record_%d", uid),
		},
	}
	if isBlacklist && isAdmin {
		inlineKeyboard01 = append(inlineKeyboard01, gotgbot.InlineKeyboardButton{
			Text:         "解除黑名单",
			CallbackData: fmt.Sprintf("uid_unban_%d", uid),
		})
	}
	inlineKeyboard = append(inlineKeyboard, inlineKeyboard01)
	if !isBlacklist && isAdmin {
		inlineKeyboard = append(inlineKeyboard, []gotgbot.InlineKeyboardButton{
			{
				Text:         "封禁 3 个月",
				CallbackData: fmt.Sprintf("uid_ban_%d_3", uid),
			},
			{
				Text:         "封禁 6 个月",
				CallbackData: fmt.Sprintf("uid_ban_%d_6", uid),
			},
		})
		inlineKeyboard = append(inlineKeyboard, []gotgbot.InlineKeyboardButton{
			{
				Text:         "封禁 1 年",
				CallbackData: fmt.Sprintf("uid_ban_%d_12", uid),
			},
			{
				Text:         "封禁 10 年",
				CallbackData: fmt.Sprintf("uid_ban_%d_120", uid),
			},
		})
	}

	return &gotgbot.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard}
}

func (tg *TelegramBot) genUidResp(uid int64, isMarkdown bool, isAdmin bool) (string, *gotgbot.InlineKeyboardMarkup, error) {
	user, err := tg.db.GetBiliUser(uid)
	if err == sql.ErrNoRows {
		var text string
		if isMarkdown {
			text = fmt.Sprintf(
				"uid: `%d`",
				uid,
			)
		} else {
			text = fmt.Sprintf(
				"uid: %d",
				uid,
			)
		}
		return text, genUidInlineKeyboard(uid, false, isAdmin), nil
	} else if err != nil {
		tg.sugar.Errorf("failed to get bilibili user: %v", err)
		return "", nil, err
	}

	location, _ := time.LoadLocation("Asia/Shanghai")

	var banUntil time.Time
	var isBlacklisted bool
	if user.BanUntil.Valid && user.BanUntil.Time.After(time.Now()) {
		isBlacklisted = true
		banUntil = user.BanUntil.Time
	}

	var text string
	if isMarkdown {
		text = fmt.Sprintf("uid: `%d`\n", uid)
		if isAdmin {
			if user.Counter > 0 {
				text += fmt.Sprintf(
					"请求次数: `%d`\n最后请求时间: `%s`\n",
					user.Counter,
					EscapeMarkdownV2(user.ModifiedAt.In(location).Format(TIME_FORMAT)),
				)
			}
		}
	} else {
		text = fmt.Sprintf("uid: %d\n", uid)
		if isAdmin {
			if user.Counter > 0 {
				text += fmt.Sprintf(
					"请求次数: %d\n最后请求时间: %s\n",
					user.Counter,
					EscapeMarkdownV2(user.ModifiedAt.In(location).Format(TIME_FORMAT)),
				)
			}
		}
	}

	if user.IsWhitelist {
		if isMarkdown {
			text += "该用户是*白名单*用户\n"
		} else {
			text += "该用户是 白名单 用户\n"
		}
	}

	if isBlacklisted {
		if isMarkdown {
			text += fmt.Sprintf("黑名单解除时间: `%s`\n", EscapeMarkdownV2(banUntil.In(location).Format(TIME_FORMAT)))
		} else {
			text += fmt.Sprintf("黑名单解除时间: %s\n", EscapeMarkdownV2(banUntil.In(location).Format(TIME_FORMAT)))
		}
	}

	if !user.IsWhitelist && !isBlacklisted {
		text += "并未查找到任何记录\n"
	}

	replyMarkup := genUidInlineKeyboard(user.UID, isBlacklisted, isAdmin)

	return text, replyMarkup, nil
}

func (tg *TelegramBot) commandUid(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveSender.IsUser() && ctx.EffectiveChat.Id == ctx.EffectiveSender.Id() {
		return nil
	}

	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	if !ctx.EffectiveSender.IsUser() {
		return nil
	}

	isAdmin := IsLevelAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id()))

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

	text, replyMarkup, err := tg.genUidResp(uid, true, isAdmin)
	if err != nil {
		tg.sugar.Errorf("failed to generate uid response: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "查询失败", nil)
		return err
	}

	sendMessageOpts := &gotgbot.SendMessageOpts{
		DisableWebPagePreview: true,
		ParseMode:             "MarkdownV2",
	}

	if replyMarkup != nil {
		sendMessageOpts.ReplyMarkup = replyMarkup
	}

	if _, err = ctx.EffectiveMessage.Reply(b, text, sendMessageOpts); err != nil {
		return err
	}

	return err
}

func (tg *TelegramBot) callbackUid(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, "uid_")
}

func (tg *TelegramBot) callbackUidResp(b *gotgbot.Bot, ctx *ext.Context) error {
	callbackData := ctx.CallbackQuery.Data
	switch {
	case strings.HasPrefix(callbackData, "uid_record_"):
		uidStr := strings.TrimPrefix(callbackData, "uid_record_")
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			return err
		}
		text, replyMarkup, err := tg.genRecordResp(uid, 1)
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
	case strings.HasPrefix(callbackData, "uid_ban_"):
		if !IsLevelAdmin(tg.GetUserAdminLevel(ctx.CallbackQuery.From.Id)) {
			return nil
		}

		data := strings.TrimPrefix(callbackData, "uid_ban_")
		splits := strings.Split(data, "_")
		uidStr := splits[0]
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			return err
		}
		monthsStr := splits[1]
		months, err := strconv.Atoi(monthsStr)
		if err != nil {
			return err
		}
		banUntil := time.Now().AddDate(0, months, 0)
		if err := tg.db.BanBiliUser(uid, banUntil); err != nil {
			tg.sugar.Errorf("failed to ban bilibili user: %v", err)
			_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "封禁失败",
				ShowAlert: true,
			})
			return err
		}
		_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text: "封禁成功",
		})
		if err != nil {
			return err
		}
		text, replyMarkup, err := tg.genUidResp(uid, true, true)
		if err != nil {
			tg.sugar.Errorf("failed to generate uid response: %v", err)
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "查询失败",
			})
			return err
		}
		_, _, err = ctx.EffectiveMessage.EditText(b, text, &gotgbot.EditMessageTextOpts{
			DisableWebPagePreview: true,
			ParseMode:             "MarkdownV2",
			ReplyMarkup:           *replyMarkup,
		})
		return err
	case strings.HasPrefix(callbackData, "uid_unban_"):
		if !IsLevelAdmin(tg.GetUserAdminLevel(ctx.CallbackQuery.From.Id)) {
			return nil
		}

		uidStr := strings.TrimPrefix(callbackData, "uid_unban_")
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			return err
		}
		if err := tg.db.UnbanBiliUser(uid); err != nil {
			tg.sugar.Errorf("failed to unban bilibili user: %v", err)
			_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "解封失败",
				ShowAlert: true,
			})
			return err
		}
		_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text: "解封成功",
		})
		if err != nil {
			return err
		}
		text, replyMarkup, err := tg.genUidResp(uid, true, true)
		if err != nil {
			tg.sugar.Errorf("failed to generate uid response: %v", err)
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "查询失败",
			})
			return err
		}
		_, _, err = ctx.EffectiveMessage.EditText(b, text, &gotgbot.EditMessageTextOpts{
			DisableWebPagePreview: true,
			ParseMode:             "MarkdownV2",
			ReplyMarkup:           *replyMarkup,
		})
		return err
	}
	return nil
}

func (tg *TelegramBot) inlineQuery(iq *gotgbot.InlineQuery) bool {
	if iq.Query == "" {
		return false
	}
	if _, err := strconv.ParseInt(iq.Query, 10, 64); err != nil {
		return false
	}
	return true
}

func (tg *TelegramBot) inlineQueryResp(b *gotgbot.Bot, ctx *ext.Context) error {
	uidStr := ctx.InlineQuery.Query
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		return err
	}

	isAdmin := IsLevelAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id()))

	text, _, err := tg.genUidResp(uid, true, isAdmin)
	if err != nil {
		tg.sugar.Errorf("failed to generate uid response: %v", err)
		return nil
	}

	inputMessageText := &gotgbot.InputTextMessageContent{
		MessageText:           text,
		ParseMode:             "MarkdownV2",
		DisableWebPagePreview: true,
	}
	inlineQueryResultArticle := gotgbot.InlineQueryResultArticle{
		Id:                  "uid_" + uidStr,
		Title:               "UID: " + uidStr,
		InputMessageContent: inputMessageText,
	}

	_, err = b.AnswerInlineQuery(ctx.InlineQuery.Id, []gotgbot.InlineQueryResult{inlineQueryResultArticle}, nil)
	return err
}
