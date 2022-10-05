package bot

import "strings"

func (tg *TelegramBot) GetUserAdminLevel(id int64) int16 {
	user, err := tg.db.GetAdmin(id)
	if err != nil {
		return 0
	}
	return user.Level
}

func IsLevelSuperAdmin(level int16) bool {
	return level == 1
}

func IsLevelAdmin(level int16) bool {
	return level > 0
}

var allMdV2 = []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
var mdV2Repl = strings.NewReplacer(func() (out []string) {
	for _, x := range allMdV2 {
		out = append(out, x, "\\"+x)
	}
	return out
}()...)

func EscapeMarkdownV2(s string) string {
	return mdV2Repl.Replace(s)
}
