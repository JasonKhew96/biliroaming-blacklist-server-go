package bot

import (
	"biliroaming-blacklist-server-go/db"
	"biliroaming-blacklist-server-go/utils"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GenReportInlineKeyboard(uid int64, reportId int) gotgbot.InlineKeyboardMarkup {
	inlineKeyboard := [][]gotgbot.InlineKeyboardButton{
		{
			{
				Text:         "查询记录",
				CallbackData: fmt.Sprintf("report_record_%d", uid),
			},
		},
		{
			{
				Text:         "接受",
				CallbackData: fmt.Sprintf("report_confirm_%d", reportId),
			},
			{
				Text:         "拒绝",
				CallbackData: fmt.Sprintf("report_decline_%d", reportId),
			},
		},
	}

	return gotgbot.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard}
}

func (tg *TelegramBot) commandReport(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	msg := ctx.EffectiveMessage

	var text string
	if msg.Text != "" {
		text = msg.Text
	} else if msg.Caption != "" {
		text = msg.Caption
	} else {
		return nil
	}

	splits := strings.Split(text, " ")
	if len(splits) < 3 {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	idStr := splits[1]
	uid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}
	description := strings.Join(splits[2:], " ")

	var fileType int16
	var fileId string

	switch {
	case msg.Animation != nil:
		fileType = db.FILE_TYPE_ANIMATION
		fileId = msg.Animation.FileId
	case msg.Document != nil:
		fileType = db.FILE_TYPE_DOCUMENT
		fileId = msg.Document.FileId
	case msg.Photo != nil:
		fileType = db.FILE_TYPE_PHOTO
		fileId = msg.Photo[len(msg.Photo)-1].FileId
	case msg.Video != nil:
		fileType = db.FILE_TYPE_VIDEO
		fileId = msg.Video.FileId
	case msg.ReplyToMessage != nil:
		repliedMsg := msg.ReplyToMessage
		switch {
		case repliedMsg.Animation != nil:
			fileType = db.FILE_TYPE_ANIMATION
			fileId = repliedMsg.Animation.FileId
		case repliedMsg.Document != nil:
			fileType = db.FILE_TYPE_DOCUMENT
			fileId = repliedMsg.Document.FileId
		case repliedMsg.Photo != nil:
			fileType = db.FILE_TYPE_PHOTO
			fileId = repliedMsg.Photo[len(repliedMsg.Photo)-1].FileId
		case repliedMsg.Video != nil:
			fileType = db.FILE_TYPE_VIDEO
			fileId = repliedMsg.Video.FileId
		}
	}

	if fileType == db.FILE_TYPE_NONE {
		_, err := ctx.EffectiveMessage.Reply(b, "未找到证据", nil)
		return err
	}

	accInfo, err := utils.GetUserInfo(uid)
	if err != nil {
		tg.sugar.Errorf("failed to get acc info: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("获取用户信息失败: %s", err.Error()), nil)
		return err
	}

	newText := fmt.Sprintf(
		"*新举报*\nUID: `%d`\n昵称: [%s](https://space.bilibili.com/%d)\n描述:\n%s",
		accInfo.Mid,
		EscapeMarkdownV2(accInfo.Name), accInfo.Mid,
		EscapeMarkdownV2(description),
	)

	reportId, err := tg.db.InsertReport(uid, description, fileType, fileId, fmt.Sprintf("TG@%d", ctx.EffectiveSender.Id()))
	if err != nil {
		tg.sugar.Errorf("failed to insert report: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	replyMarkup := GenReportInlineKeyboard(uid, reportId)

	switch fileType {
	case db.FILE_TYPE_ANIMATION:
		_, err = b.SendAnimation(tg.ReportChatId, fileId, &gotgbot.SendAnimationOpts{
			Caption:     newText,
			ParseMode:   "MarkdownV2",
			ReplyMarkup: replyMarkup,
		})
	case db.FILE_TYPE_DOCUMENT:
		_, err = b.SendDocument(tg.ReportChatId, fileId, &gotgbot.SendDocumentOpts{
			Caption:     newText,
			ParseMode:   "MarkdownV2",
			ReplyMarkup: replyMarkup,
		})
	case db.FILE_TYPE_PHOTO:
		_, err = b.SendPhoto(tg.ReportChatId, fileId, &gotgbot.SendPhotoOpts{
			Caption:     newText,
			ParseMode:   "MarkdownV2",
			ReplyMarkup: replyMarkup,
		})
	case db.FILE_TYPE_VIDEO:
		_, err = b.SendVideo(tg.ReportChatId, fileId, &gotgbot.SendVideoOpts{
			Caption:     newText,
			ParseMode:   "MarkdownV2",
			ReplyMarkup: replyMarkup,
		})
	}
	if err != nil {
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(b, "已提交举报", nil)
	return err
}

func (tg *TelegramBot) callbackReport(cq *gotgbot.CallbackQuery) bool {
	if !strings.HasPrefix(cq.Data, "report_") {
		return false
	}
	if !IsLevelAdmin(tg.GetUserAdminLevel(cq.From.Id)) {
		return false
	}
	return true
}

func (tg *TelegramBot) callbackReportResp(b *gotgbot.Bot, ctx *ext.Context) error {
	callbackData := ctx.CallbackQuery.Data
	switch {
	case strings.HasPrefix(callbackData, "report_record_"):
		uidStr := strings.TrimPrefix(callbackData, "report_record_")
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "参数错误",
			})
			return err
		}
		text, _, err := tg.genUidResp(uid, false, true)
		if err == sql.ErrNoRows {
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "未找到记录",
			})
			return err
		} else if err != nil {
			tg.sugar.Errorf("failed to gen uid resp: %v", err)
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "数据库错误",
			})
			return err
		}
		_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      text,
			ShowAlert: true,
		})
		return err
	case strings.HasPrefix(callbackData, "report_confirm_"):
		reportIdStr := strings.TrimPrefix(callbackData, "report_confirm_")
		reportId, err := strconv.Atoi(reportIdStr)
		if err != nil {
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "参数错误",
			})
			return err
		}
		report, err := tg.db.GetReport(reportId)
		if err != nil {
			tg.sugar.Errorf("failed to get report: %v", err)
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "数据库错误",
			})
			return err
		}

		accInfo, err := utils.GetUserInfo(report.UID)
		if err != nil {
			tg.sugar.Errorf("failed to get bili acc info: %v", err)
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text: "B站API错误",
			})
			return err
		}

		newText := fmt.Sprintf(
			"*新证据*\nUID: `%d`\n昵称: [%s](https://space.bilibili.com/%d)\n描述:\n%s",
			accInfo.Mid,
			EscapeMarkdownV2(accInfo.Name), accInfo.Mid,
			EscapeMarkdownV2(report.Description),
		)

		var msg *gotgbot.Message
		switch report.FileType {
		case db.FILE_TYPE_ANIMATION:
			msg, err = b.SendAnimation(tg.AnnouceChatId, report.FileID, &gotgbot.SendAnimationOpts{
				Caption:   newText,
				ParseMode: "MarkdownV2",
			})
		case db.FILE_TYPE_DOCUMENT:
			msg, err = b.SendDocument(tg.AnnouceChatId, report.FileID, &gotgbot.SendDocumentOpts{
				Caption:   newText,
				ParseMode: "MarkdownV2",
			})
		case db.FILE_TYPE_PHOTO:
			msg, err = b.SendPhoto(tg.AnnouceChatId, report.FileID, &gotgbot.SendPhotoOpts{
				Caption:   newText,
				ParseMode: "MarkdownV2",
			})
		case db.FILE_TYPE_VIDEO:
			msg, err = b.SendVideo(tg.AnnouceChatId, report.FileID, &gotgbot.SendVideoOpts{
				Caption:   newText,
				ParseMode: "MarkdownV2",
			})
		}
		if err != nil {
			return err
		}

		if _, err = tg.db.InsertRecord(report.UID, report.Description, msg.Chat.Id, msg.MessageId, ctx.CallbackQuery.From.Id); err != nil {
			tg.sugar.Errorf("failed to insert record: %v", err)
			return err
		}

		user, err := tg.db.GetBiliUser(report.UID)
		if err != nil || user.BanUntil.IsZero() || !user.BanUntil.IsZero() && user.BanUntil.Time.Before(time.Now()) {
			banUntil, _ := utils.ParseDuration(tg.DefaultBanTime)

			_, err = tg.db.BanBiliUser(report.UID, *banUntil)
			if err != nil {
				tg.sugar.Errorf("failed to ban user: %v", err)
			}
		}

		if _, err = b.DeleteMessage(ctx.EffectiveChat.Id, ctx.EffectiveMessage.MessageId, nil); err != nil {
			return err
		}
		return nil
	case strings.HasPrefix(callbackData, "report_decline_"):
		if _, err := b.DeleteMessage(ctx.EffectiveChat.Id, ctx.EffectiveMessage.MessageId, nil); err != nil {
			return err
		}
		return nil
	}
	return nil
}
