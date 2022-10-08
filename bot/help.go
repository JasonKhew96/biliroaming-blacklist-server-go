package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandHelp(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	text := `*帮助*
/help \- 显示这条帮助信息
/uid <uid\> \- 查询用户
/record <uid\> \- 查询违规记录
/report <uid\> <描述\> \- 举报用户，与证据一起发送 \(包括回复\)
`

	userAdminLevel := tg.GetUserAdminLevel(ctx.EffectiveSender.Id())
	if IsLevelAdmin(userAdminLevel) {
		text += `
*管理员 专用*
/ban <uid\> \[时长\] \- 封禁用户，默认三个月 \(1h/1d/1w/1m/1y\)
/unban <uid\> \- 解封用户
/white <uid\> \- 白名单用户
/unwhite <uid\> \- 移除白名单
`
	}
	if IsLevelSuperAdmin(userAdminLevel) {
		text += `
*超级管理员 专用*
/addadmin <uid\> \[等级\] \- 添加管理员，1 超级管理员，2 普通管理员\(默认\)
/removeadmin <uid\> \- 移除管理员
/alteradmin <uid\> <等级\> \- 修改管理员等级
`
	}

	_, err := ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	})
	return err
}
