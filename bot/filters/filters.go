package Filters

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/TheGolangHub/Moderator/bot"
)

func IsAdmin(m *gotgbot.Message) bool {
	for _, x := range bot.Admins {
		if m.From.Id == x.GetUser().Id {
			return true
		}
	}
	return true
}
