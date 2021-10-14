package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TheGolangHub/Moderator/bot"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	// Create bot from environment value.
	b, err := gotgbot.NewBot(os.Getenv("BOT_TOKEN"), &gotgbot.BotOpts{
		APIURL:      "http://localhost:8081",
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				fmt.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			Panic:       nil,
			ErrorLog:    nil,
			MaxRoutines: 0,
		},
	})
	dispatcher := updater.Dispatcher

	// Add echo handler to reply to all messages.
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	bot.Load(dispatcher)

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

func start(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := bot.SendMessage(ctx.EffectiveChat.Id, "Hi, I am started.", &gotgbot.SendMessageOpts{
		ReplyToMessageId: ctx.EffectiveMessage.MessageId,
	})
	return err
}
