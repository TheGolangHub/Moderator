package bot

import (
	"log"

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

func IsUserAdmin(bot *gotgbot.Bot, ctx *ext.Context) bool {
	if ctx.EffectiveUser.Id == 1087968824 {
		return true
	}
	member, err := ctx.EffectiveChat.GetMember(bot, ctx.EffectiveUser.Id)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	for _, x := range Admins {
		if x == member {
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
