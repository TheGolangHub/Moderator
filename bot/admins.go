package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/TheGolangHub/Moderator/config"
)

var Admins = []gotgbot.ChatMember{}

func RefreshAdmins(b *gotgbot.Bot, ctx *ext.Context) error {
	Admins, _ = b.GetChatAdministrators(ctx.EffectiveChat.Id)
	return nil
}

func IsUserAdmin(userId int64) bool {
	if userId == 1087968824 || userId == 777000 {
		return true
	}
	for _, x := range Admins {
		if x.GetUser().Id == userId {
			return true
		}
	}
	return false
}

func loadAdmins(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewChatMember(func(u *gotgbot.ChatMemberUpdated) bool {
		if u.Chat.Id != config.CHAT_ID {
			return false
		}
		if u.NewChatMember.MergeChatMember().Status == "administrator" && u.OldChatMember.MergeChatMember().Status != "administrator" {
			return true
		}
		if u.NewChatMember.MergeChatMember().Status != "administrator" && u.OldChatMember.MergeChatMember().Status == "administrator" {
			return true
		}
		return false
	}, RefreshAdmins))
}
