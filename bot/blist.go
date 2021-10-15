package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/TheGolangHub/Moderator/bot/utils"
	"github.com/TheGolangHub/Moderator/bot/utils/data"
	"github.com/TheGolangHub/Moderator/config"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func profanityCheck(b *gotgbot.Bot, ctx *ext.Context) error {
	if IsUserAdmin(ctx.EffectiveUser.Id) {
		return ext.EndGroups
	}
	msg := ctx.EffectiveMessage
	chat := ctx.EffectiveChat
	msg.Delete(b)
	if data.D.RulebreakCount[ctx.EffectiveUser.Id] > 5 {
		data.D.RulebreakCount[ctx.EffectiveUser.Id] = 0
		banned, _ := b.BanChatMember(chat.Id, ctx.EffectiveUser.Id, &gotgbot.BanChatMemberOpts{
			UntilDate: time.Now().Unix() + int64(time.Hour*24),
		})
		if banned {
			b.SendMessage(chat.Id, fmt.Sprintf("%s has broken the rules more than 5 times so has been banned for a day.", utils.MentionUser(ctx.EffectiveUser, "html")), &gotgbot.SendMessageOpts{
				ParseMode: "html",
			})
		}
	}
	data.D.RulebreakCount[ctx.EffectiveUser.Id] += 1
	b.SendMessage(chat.Id, fmt.Sprintf("Hey %s, It looks like you've forgotten the rules. Read them again and follow properly otherwise I will have to take strict actions against you.", utils.MentionUser(ctx.EffectiveUser, "html")), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	return ext.EndGroups
}

func offTopicChat(b *gotgbot.Bot, ctx *ext.Context) error {
	if IsUserAdmin(ctx.EffectiveUser.Id) {
		return ext.EndGroups
	}
	ctx.EffectiveMessage.Delete(b)
	chat := ctx.EffectiveChat
	if data.D.RulebreakCount[ctx.EffectiveUser.Id] > 5 {
		data.D.RulebreakCount[ctx.EffectiveUser.Id] = 0
		banned, _ := b.BanChatMember(chat.Id, ctx.EffectiveUser.Id, &gotgbot.BanChatMemberOpts{
			UntilDate: time.Now().Unix() + int64(time.Hour*24),
		})
		if banned {
			b.SendMessage(chat.Id, fmt.Sprintf("%s has broken the rules more than 5 times so has been banned for a day.", utils.MentionUser(ctx.EffectiveUser, "html")), &gotgbot.SendMessageOpts{
				ParseMode: "html",
			})
		}
	}
	data.D.RulebreakCount[ctx.EffectiveUser.Id] += 1
	b.SendMessage(chat.Id, fmt.Sprintf("Hey %s, It looks like you've forgotten the rules. Read them again and follow properly otherwise I will have to take strict actions against you.", utils.MentionUser(ctx.EffectiveUser, "html")), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	return ext.EndGroups
}

func loadBlist(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandlerToGroup(handlers.NewMessage(func(msg *gotgbot.Message) bool {
		if msg.Chat.Id != config.CHAT_ID {
			return false
		}
		if !message.Text(msg) {
			return false
		}
		sep := " "
		if len(strings.Split(msg.Text, " ")) < 2 {
			sep = ""
		}
		for _, x := range data.ProfanityList {
			if strings.Contains(strings.ToLower(msg.Text)+sep, x) {
				return true
			}
		}
		return false
	}, profanityCheck), ProfanityGroup)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(func(msg *gotgbot.Message) bool {
		if msg.Chat.Id != config.CHAT_ID {
			return false
		}
		if !message.Text(msg) {
			return false
		}
		sep := " "
		if len(strings.Split(msg.Text, " ")) < 2 {
			sep = ""
		}
		for _, x := range data.OffTalk {
			if strings.Contains(strings.ToLower(msg.Text)+sep, x) {
				return true
			}
		}
		return false
	}, offTopicChat), NormTalkGroup)
}
