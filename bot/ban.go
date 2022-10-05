package bot

import (
	"biliroaming-blacklist-server-go/utils"
	"log"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandBan(b *gotgbot.Bot, ctx *ext.Context) error {
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

	banUntil, _ := utils.ParseDuration("3m")
	if len(splits) > 2 {
		banUntil, err = utils.ParseDuration(splits[2])
		if err != nil {
			_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
			return err
		}
	}

	err = tg.db.BanBiliUser(uid, *banUntil)
	if err != nil {
		log.Println("failed to ban user: ", err.Error())
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(b, "已封禁", nil)
	return err
}

func (tg *TelegramBot) commandUnban(b *gotgbot.Bot, ctx *ext.Context) error {
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

	err = tg.db.UnbanBiliUser(uid)
	if err != nil {
		log.Println("failed to ban user: ", err.Error())
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(b, "已解除封禁", nil)
	return err
}
