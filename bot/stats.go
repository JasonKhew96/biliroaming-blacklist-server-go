package bot

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandStats(b *gotgbot.Bot, ctx *ext.Context) error {
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
		"总用户数: %d\n总记录数: %d\n总举报数: %d\n",
		totalUser,
		totalRecord,
		totalReport,
	)

	_, err = ctx.EffectiveMessage.Reply(b, text, nil)
	return err
}
