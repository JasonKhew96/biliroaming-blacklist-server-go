package bot

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandStats(b *gotgbot.Bot, ctx *ext.Context) error {
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

	totalUser, err := tg.db.GetTotalUser()
	if err != nil {
		tg.sugar.Errorf("failed to get total user: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	totalActiveUser, err := tg.db.GetTotalActiveUser()
	if err != nil {
		tg.sugar.Errorf("failed to get total active user: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	totalBannedUser, err := tg.db.GetTotalBannedUser()
	if err != nil {
		tg.sugar.Errorf("failed to get total banned user: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	totalRecord, err := tg.db.GetTotalRecord()
	if err != nil {
		tg.sugar.Errorf("failed to get total record: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	totalReport, err := tg.db.GetTotalReport()
	if err != nil {
		tg.sugar.Errorf("failed to get total report: %v", err)
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	text := fmt.Sprintf(
		"总用户数: %d (%d)\n被封禁数: %d\n总记录数: %d\n总举报数: %d\n",
		totalUser,
		totalActiveUser,
		totalBannedUser,
		totalRecord,
		totalReport,
	)

	_, err = ctx.EffectiveMessage.Reply(b, text, nil)
	return err
}
