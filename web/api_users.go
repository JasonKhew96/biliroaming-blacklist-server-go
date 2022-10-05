package web

import (
	"biliroaming-blacklist-server-go/entity"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (w *Web) usersUid(c *fiber.Ctx) error {
	resp := &entity.RespStatus{}
	uidStr := utils.CopyString(c.Params("uid"))
	if uidStr == "" {
		uidStr = utils.CopyString(c.Query("uid"))
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil || uid <= 0 {
		resp.Code = 400
		resp.Message = "Invalid UID"
		return c.JSON(resp)
	}

	userAgent := string(utils.CopyBytes(c.Request().Header.UserAgent()))
	if strings.HasPrefix(userAgent, "biliroaming") {
		w.db.IncBiliUserCounter(uid)
	}

	user, err := w.db.GetBiliUser(uid)
	if err != nil {
		resp.Data = &entity.RespStatusData{
			UID:         uid,
			IsBlacklist: false,
			IsWhitelist: false,
			Status:      StatusNone,
			BanUntil:    0,
		}
		return c.JSON(resp)
	}

	if user.BanUntil.Valid && user.BanUntil.Time.After(time.Now()) {
		resp.Data = &entity.RespStatusData{
			UID:         uid,
			IsBlacklist: true,
			IsWhitelist: false,
			Status:      StatusBlacklist,
			BanUntil:    user.BanUntil.Time.Unix(),
		}
		return c.JSON(resp)
	}

	var status int8
	if user.IsWhitelist {
		status = StatusWhitelist
	} else {
		status = StatusNone
	}

	resp.Code = 0

	resp.Data = &entity.RespStatusData{
		UID:         uid,
		IsBlacklist: false,
		IsWhitelist: user.IsWhitelist,
		Status:      status,
		BanUntil:    0,
	}

	return c.JSON(resp)
}
