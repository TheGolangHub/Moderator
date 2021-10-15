package utils

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func MentionUser(u *gotgbot.User, pmode string) string {
	pmode = strings.ToLower(pmode)
	if pmode == "html" {
		return fmt.Sprintf(`<a href="tg://user?id=%d">%s</a>`, u.Id, u.FirstName)
	} else {
		return fmt.Sprintf("[%s](tg://user?id=%d)", u.FirstName, u.Id)
	}
}

func StringContainWord(s string, word string) bool {
	s = strings.ToLower(s)
	for _, i := range strings.Fields(s) {
		if i == word {
			return true
		}
	}
	return false
}
