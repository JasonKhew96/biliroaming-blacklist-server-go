package bot

import (
	"biliroaming-blacklist-server-go/utils"
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandKey(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	isAdmin := IsLevelAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id()))
	if !isAdmin {
		return nil
	}

	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	splits := strings.Split(msg.Text, " ")
	if len(splits) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	key := splits[1]
	if len(key) != 32 {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	info, err := utils.GetMyInfo(key)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("未知错误: %v", err), nil)
		return err
	}
	uid := info.Mid

	text, replyMarkup, err := tg.genUidResp(uid, true, isAdmin)
	if err != nil {
		tg.sugar.Errorf("failed to generate key response: %v", err)
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
