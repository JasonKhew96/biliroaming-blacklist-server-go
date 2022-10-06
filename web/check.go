package web

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (w *Web) checkGet(c *fiber.Ctx) error {
	return c.Render("check_get", fiber.Map{
		"Title":   "查询",
		"SiteKey": w.config.SiteKey,
	}, "layouts/main")
}

type RecordResult struct {
	Time        string
	Description string
}

func (w *Web) checkPost(c *fiber.Ctx) error {
	fiberMap := fiber.Map{
		"Title":   "查询",
	}
	uidStr := utils.CopyString(c.FormValue("uid"))
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		fiberMap["Error"] = "UID 格式错误"
		return c.Render("check_post", fiberMap, "layouts/main")
	}

	token := utils.CopyString(c.FormValue("cf-turnstile-response"))
	ip := string(utils.CopyBytes(c.Request().Header.Peek("CF-Connecting-IP")))
	if len(token) == 0 || len(ip) == 0 {
		fiberMap["Error"] = "请完成验证"
		return c.Render("check_post", fiberMap, "layouts/main")
	}
	success, err := w.verifyCaptchas(token, ip)
	if err != nil || !success {
		fiberMap["Error"] = "验证失败"
		return c.Render("check_post", fiberMap, "layouts/main")
	}

	user, err := w.db.GetBiliUser(uid)
	if err != nil && err != sql.ErrNoRows {
		fiberMap["Error"] = "查询失败"
		return c.Render("check_post", fiberMap, "layouts/main")
	} else if err != nil {
		fiberMap["Error"] = "未查询到该用户"
		return c.Render("check_post", fiberMap, "layouts/main")
	}

	records, err := w.db.GetRecords(uid, 8)
	if err != nil {
		fiberMap["Error"] = "查询失败"
		return c.Render("check_post", fiberMap, "layouts/main")
	}

	location, _ := time.LoadLocation("Asia/Shanghai")

	results := make([]RecordResult, 0, len(records))
	for _, record := range records {
		results = append(results, RecordResult{
			Time:        record.CreatedAt.In(location).Format("2006-01-02 15:04:05"),
			Description: record.Description,
		})
	}

	fiberMap["Uid"] = uid
	fiberMap["Records"] = results

	if user.BanUntil.Valid && user.BanUntil.Time.After(time.Now()) {
		fiberMap["BanUntil"] = user.BanUntil.Time.In(location).Format("2006-01-02 15:04:05")

		return c.Render("check_post", fiberMap, "layouts/main")
	}

	return c.Render("check_post", fiberMap, "layouts/main")
}
