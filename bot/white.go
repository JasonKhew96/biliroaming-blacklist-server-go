package bot

import (
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandWhite(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	if !IsLevelAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id())) {
		_, err := ctx.EffectiveMessage.Reply(b, "权限不足", nil)
		return err
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

	err = tg.db.WhiteBiliUser(uid, true)
	if err != nil {
		tg.sugar.Errorf("failed to whitelist user: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	return nil
}

func (tg *TelegramBot) commandUnwhite(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	if !IsLevelAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id())) {
		_, err := ctx.EffectiveMessage.Reply(b, "权限不足", nil)
		return err
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

	err = tg.db.WhiteBiliUser(uid, false)
	if err != nil {
		tg.sugar.Errorf("failed to unwhitelist user: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	return nil
}
