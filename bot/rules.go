package bot

import (
	"fmt"
	"strings"

	"github.com/TheGolangHub/Moderator/bot/utils"
	"github.com/TheGolangHub/Moderator/bot/utils/data"
	"github.com/TheGolangHub/Moderator/config"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func userReadRules(b *gotgbot.Bot, ctx *ext.Context) error {
	if IsUserAdmin(ctx.EffectiveUser.Id) {
		return ext.EndGroups
	}
	msg := ctx.EffectiveMessage
	user := ctx.EffectiveUser
	if strings.Contains(strings.ToLower(msg.Text), config.RULE_KEY) {
		data.D.RuledUsers = append(data.D.RuledUsers, user.Id)
		msg.Delete(b)
		b.SendMessage(ctx.EffectiveChat.Id, fmt.Sprintf("Hello %s, Hope you will follow the rules as carefully as you read them!", utils.MentionUser(user, "html")), &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
	} else {
		c := data.D.NotruledCount[user.Id]
		if c <= 5 {
			data.D.NotruledCount[user.Id] += 1
		} else {
			data.D.NotruledCount[user.Id] = 0
			kicked, _ := b.UnbanChatMember(ctx.EffectiveChat.Id, user.Id, nil)
			if kicked {
				b.SendMessage(ctx.EffectiveChat.Id, fmt.Sprintf("%s has been kicked as he was trying to talk here without reading rules.", utils.MentionUser(user, "html")), &gotgbot.SendMessageOpts{
					ParseMode: "html",
				})
			}
		}
	}
	return ext.EndGroups
}

func loadRules(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandlerToGroup(handlers.NewMessage(func(msg *gotgbot.Message) bool {
		if msg.Chat.Id != config.CHAT_ID {
			return false
		}
		if !message.Text(msg) {
			return false
		}
		if data.Int64InSlice(msg.From.Id, data.D.RuledUsers) {
			return false
		}
		return true
	}, userReadRules), RulesGroup)
}
