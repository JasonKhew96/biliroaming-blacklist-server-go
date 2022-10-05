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
		"Title": "自助查询",
	}, "layouts/main")
}

type RecordResult struct {
	Time        string
	Description string
}

func (w *Web) checkPost(c *fiber.Ctx) error {
	uidStr := utils.CopyString(c.FormValue("uid"))
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		return c.Render("check_post", fiber.Map{
			"Error": "UID格式错误",
		}, "layouts/main")
	}

	user, err := w.db.GetBiliUser(uid)
	if err != nil && err != sql.ErrNoRows {
		return c.Render("check_post", fiber.Map{
			"Title": "查询结果",
			"Error": "查询数据库时出错",
		}, "layouts/main")
	} else if err != nil {
		return c.Render("check_post", fiber.Map{
			"Title": "查询结果",
			"Error": "未找到任何记录",
		}, "layouts/main")
	}

	records, err := w.db.GetRecords(uid, 8)
	if err != nil {
		return c.Render("check_post", fiber.Map{
			"Title": "查询结果",
			"Error": "查询数据库时出错",
		}, "layouts/main")
	}

	location, _ := time.LoadLocation("Asia/Shanghai")

	results := make([]RecordResult, 0, len(records))
	for _, record := range records {
		results = append(results, RecordResult{
			Time:        record.CreatedAt.In(location).Format("2006-01-02 15:04:05"),
			Description: record.Description,
		})
	}

	if user.BanUntil.Valid && user.BanUntil.Time.After(time.Now()) {
		return c.Render("check_post", fiber.Map{
			"Title":    "查询结果",
			"Uid":      c.FormValue("uid"),
			"BanUntil": user.BanUntil.Time.In(location).Format("2006-01-02 15:04:05"),
			"Records":  results,
		}, "layouts/main")
	}

	return c.Render("check_post", fiber.Map{
		"Title":   "查询结果",
		"Uid":     c.FormValue("uid"),
		"Records": results,
	}, "layouts/main")
}
