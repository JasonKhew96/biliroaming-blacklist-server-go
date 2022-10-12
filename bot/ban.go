package bot

import (
	"biliroaming-blacklist-server-go/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandBan(b *gotgbot.Bot, ctx *ext.Context) error {
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

	accInfo, err := utils.GetBiliAccInfo(uid)
	if err != nil {
		tg.sugar.Errorf("failed to get acc info: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "获取用户信息失败", nil)
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

	_, err = tg.db.BanBiliUser(uid, *banUntil)
	if err != nil {
		tg.sugar.Errorf("failed to ban user: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	location, _ := time.LoadLocation("Asia/Shanghai")
	text := fmt.Sprintf("*已封禁*\nUID: `%d`\n昵称: [%s](https://space.bilibili.com/%d)\n将在 `%s` 后解除",
		accInfo.Mid,
		accInfo.Name,
		accInfo.Mid,
		banUntil.In(location).Format(TIME_FORMAT),
	)

	_, err = ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	})
	return err
}

func (tg *TelegramBot) commandUnban(b *gotgbot.Bot, ctx *ext.Context) error {
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

	accInfo, err := utils.GetBiliAccInfo(uid)
	if err != nil {
		tg.sugar.Errorf("failed to get acc info: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "获取用户信息失败", nil)
		return err
	}

	_, err = tg.db.UnbanBiliUser(uid)
	if err != nil {
		tg.sugar.Errorf("failed to unban user: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	text := fmt.Sprintf("*已解除封禁*\nUID: `%d`\n昵称: [%s](https://space.bilibili.com/%d)\n",
		accInfo.Mid,
		accInfo.Name,
		accInfo.Mid,
	)

	_, err = ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	})
	return err
}
