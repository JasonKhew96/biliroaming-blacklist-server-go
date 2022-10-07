package bot

import (
	"biliroaming-blacklist-server-go/models"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (tg *TelegramBot) commandAddAdmin(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	if !IsLevelSuperAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id())) {
		_, err := ctx.EffectiveMessage.Reply(b, "权限不足", nil)
		return err
	}

	splits := strings.Split(msg.Text, " ")
	if len(splits) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	idStr := splits[1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	var level int64 = 2
	if len(splits) > 2 {
		level, err = strconv.ParseInt(splits[2], 10, 16)
		if err != nil {
			_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
			return err
		}
	}

	_, err = tg.db.GetAdmin(id)
	if err != nil && err != sql.ErrNoRows {
		tg.sugar.Errorf("failed to get admin: %s", err.Error())
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	} else if err == nil {
		_, err := ctx.EffectiveMessage.Reply(b, "该用户已经是管理员了", nil)
		return err
	}

	err = tg.db.UpsertAdmin(&models.Admin{
		ID:         id,
		Level:      int16(level),
		ModifiedAt: time.Now(),
	})
	if err != nil {
		tg.sugar.Errorf("failed to upsert admin: %s", err.Error())
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	return nil
}

func (tg *TelegramBot) commandRemoveAdmin(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	if !IsLevelSuperAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id())) {
		_, err := ctx.EffectiveMessage.Reply(b, "权限不足", nil)
		return err
	}

	splits := strings.Split(msg.Text, " ")
	if len(splits) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	idStr := splits[1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	if _, err = tg.db.RemoveAdmin(id); err != nil {
		tg.sugar.Errorf("failed to remove admin: %s", err.Error())
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	return nil
}

func (tg *TelegramBot) commandAlterAdmin(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() || ctx.EffectiveChat.Id != ctx.EffectiveSender.Id() {
		return nil
	}

	msg := ctx.EffectiveMessage
	if msg.Text == "" {
		return nil
	}

	if !IsLevelSuperAdmin(tg.GetUserAdminLevel(ctx.EffectiveSender.Id())) {
		_, err := ctx.EffectiveMessage.Reply(b, "权限不足", nil)
		return err
	}

	splits := strings.Split(msg.Text, " ")
	if len(splits) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	idStr := splits[1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
		return err
	}

	var level int64 = 2
	if len(splits) > 2 {
		level, err = strconv.ParseInt(splits[2], 10, 16)
		if err != nil {
			_, err := ctx.EffectiveMessage.Reply(b, "参数错误", nil)
			return err
		}
	}

	admin, err := tg.db.GetAdmin(id)
	if err != nil {
		tg.sugar.Errorf("failed to get admin: %s", err.Error())
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	admin.Level = int16(level)
	admin.ModifiedAt = time.Now()

	err = tg.db.UpsertAdmin(admin)
	if err != nil {
		tg.sugar.Errorf("failed to upsert admin: %s", err.Error())
		_, err := ctx.EffectiveMessage.Reply(b, "数据库错误", nil)
		return err
	}

	return nil
}
