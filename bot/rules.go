package bot

import (
	"fmt"
	"os"
	"strings"

	"github.com/TheGolangHub/Moderator/bot/utils"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func userReadRules(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	if strings.Contains(msg.Text, os.Getenv("RULE_KEY")) {
		msg.Delete(b)
		b.SendMessage(ctx.EffectiveChat.Id, fmt.Sprintf("Hello %s, Hope you will follow rules as carefully as you read them!", utils.MentionUser(ctx.EffectiveUser, "html")), &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
	}
	return nil
}

func loadRules(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewMessage(message.Text, userReadRules))
}
