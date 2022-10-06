package web

import (
	"fmt"
	"image"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"biliroaming-blacklist-server-go/bot"
	"biliroaming-blacklist-server-go/db"
	ut "biliroaming-blacklist-server-go/utils"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func (w *Web) reportGet(c *fiber.Ctx) error {
	fiberMap := fiber.Map{
		"Title":   "自助举报",
		"SiteKey": w.config.SiteKey,
	}
	return c.Render("report", fiberMap, "layouts/main")
}

func isUploadFileSupported(contentType string) bool {
	if strings.Contains(contentType, "image/jpeg") || strings.Contains(contentType, "image/png") || strings.Contains(contentType, "video/mp4") {
		return true
	}
	return false
}

func (w *Web) reportPost(c *fiber.Ctx) error {
	fiberMap := fiber.Map{
		"Title":   "自助举报",
		"SiteKey": w.config.SiteKey,
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		fiberMap["Message"] = "请选择文件"
		return c.Render("report", fiberMap, "layouts/main")
	}

	contentTypes := fileHeader.Header.Values("Content-Type")
	if len(contentTypes) != 1 {
		fiberMap["Message"] = "未知文件类型"
		return c.Render("report", fiberMap, "layouts/main")
	} else if !isUploadFileSupported(contentTypes[0]) {
		fiberMap["Message"] = "文件类型不支持"
		return c.Render("report", fiberMap, "layouts/main")
	}

	if fileHeader.Size > 50*1024*1024 { // 50 MiB
		fiberMap["Message"] = "文件大小超过限制"
		return c.Render("report", fiberMap, "layouts/main")
	}

	description := strings.TrimSpace(utils.CopyString(c.FormValue("description")))
	if len(description) == 0 {
		fiberMap["Message"] = "请填写举报描述"
		return c.Render("report", fiberMap, "layouts/main")
	}
	uidStr := utils.CopyString(c.FormValue("uid"))
	if len(uidStr) == 0 {
		fiberMap["Message"] = "请填写 UID"
		return c.Render("report", fiberMap, "layouts/main")
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil || uid <= 0 {
		fiberMap["Message"] = "UID 格式错误"
		return c.Render("report", fiberMap, "layouts/main")
	}

	token := utils.CopyString(c.FormValue("cf-turnstile-response"))
	ip := string(utils.CopyBytes(c.Request().Header.Peek("CF-Connecting-IP")))
	if len(token) == 0 || len(ip) == 0 {
		fiberMap["Message"] = "请完成验证"
		return c.Render("report", fiberMap, "layouts/main")
	}
	success, err := w.verifyCaptchas(token, ip)
	if err != nil || !success {
		fiberMap["Message"] = "验证失败"
		return c.Render("report", fiberMap, "layouts/main")
	}

	accInfo, err := ut.GetBiliAccInfo(uid)
	if err != nil {
		fiberMap["Message"] = "获取用户信息失败"
		return c.Render("report", fiberMap, "layouts/main")
	}

	if accInfo.Mid != uid {
		fiberMap["Message"] = "UID 不存在"
		return c.Render("report", fiberMap, "layouts/main")
	}

	fileName := path.Join(os.TempDir(), fmt.Sprintf("biliroaming_%d.%s", time.Now().Unix(), path.Ext(fileHeader.Filename)))
	defer os.Remove(fileName)
	if err := c.SaveFile(fileHeader, fileName); err != nil {
		fiberMap["Message"] = "保存文件失败"
		return c.Render("report", fiberMap, "layouts/main")
	}

	reportId, err := w.db.InsertReport(uid, description, 0, "")
	if err != nil {
		fiberMap["Message"] = "保存举报信息失败"
		return c.Render("report", fiberMap, "layouts/main")
	}

	newText := fmt.Sprintf(
		"*新举报*\nUID: `%d`\n昵称: [%s](https://space.bilibili.com/%d)\n描述:\n%s",
		accInfo.Mid,
		bot.EscapeMarkdownV2(accInfo.Name), accInfo.Mid,
		bot.EscapeMarkdownV2(description),
	)

	fileType := db.FILE_TYPE_DOCUMENT
	if strings.Contains(contentTypes[0], "image/") {
		fileType = db.FILE_TYPE_DOCUMENT
		if reader, err := os.Open(fileName); err == nil {
			defer reader.Close()
			im, _, err := image.DecodeConfig(reader)
			if err != nil || im.Width == 0 || im.Height == 0 {
				fiberMap["Message"] = "解析文件失败"
				return c.Render("report", fiberMap, "layouts/main")
			}
			if !(im.Width/im.Height > 5 || im.Height/im.Width > 5) {
				fileType = db.FILE_TYPE_PHOTO
			}
		}
	} else {
		fileType = db.FILE_TYPE_VIDEO
	}

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fiberMap["Message"] = "读取文件失败"
		return c.Render("report", fiberMap, "layouts/main")
	}

	var msg *gotgbot.Message
	var fileId string
	switch fileType {
	case db.FILE_TYPE_PHOTO:
		msg, err = w.tg.Bot.SendPhoto(w.tg.ReportChatId, fileData, &gotgbot.SendPhotoOpts{
			Caption:   newText,
			ParseMode: "MarkdownV2",
		})
		if err != nil {
			fiberMap["Message"] = "发送文件失败"
			return c.Render("report", fiberMap, "layouts/main")
		}
		fileId = msg.Photo[0].FileId
	case db.FILE_TYPE_VIDEO:
		msg, err = w.tg.Bot.SendVideo(w.tg.ReportChatId, fileData, &gotgbot.SendVideoOpts{
			Caption:   newText,
			ParseMode: "MarkdownV2",
		})
		if err != nil {
			fiberMap["Message"] = "发送文件失败"
			return c.Render("report", fiberMap, "layouts/main")
		}
		fileId = msg.Video.FileId
	case db.FILE_TYPE_DOCUMENT:
		msg, err = w.tg.Bot.SendDocument(w.tg.ReportChatId, fileData, &gotgbot.SendDocumentOpts{
			Caption:   newText,
			ParseMode: "MarkdownV2",
		})
		if err != nil {
			fiberMap["Message"] = "发送文件失败"
			return c.Render("report", fiberMap, "layouts/main")
		}
		fileId = msg.Document.FileId
	}

	if _, err := w.db.UpdateReport(reportId, fileType, fileId); err != nil {
		fiberMap["Message"] = "保存举报信息失败"
		return c.Render("report", fiberMap, "layouts/main")
	}

	inlineKeyboard := bot.GenReportInlineKeyboard(uid, reportId)

	if _, _, err := w.tg.Bot.EditMessageReplyMarkup(&gotgbot.EditMessageReplyMarkupOpts{
		ChatId:      w.tg.ReportChatId,
		MessageId:   msg.MessageId,
		ReplyMarkup: inlineKeyboard,
	}); err != nil {
		fiberMap["Message"] = "发送文件失败"
		return c.Render("report", fiberMap, "layouts/main")
	}

	fiberMap["Message"] = "举报成功"
	return c.Render("report", fiberMap, "layouts/main")
}
