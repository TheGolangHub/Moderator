package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"

	"github.com/TheGolangHub/Moderator/bot"
	"github.com/TheGolangHub/Moderator/bot/utils/data"
	"github.com/TheGolangHub/Moderator/config"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	data.StoreOutFile()
	// Create bot from environment value.
	b, err := gotgbot.NewBot(config.TOKEN, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}
	bot.Admins, _ = b.GetChatAdministrators(config.CHAT_ID)

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
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: false})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\n", b.User.Username)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		data.SaveInFile()
		os.Exit(1)
	}()

	cron := cron.New()
	cron.AddFunc("@every 60m", func() { data.SaveInFile() })
	cron.Start()

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

func start(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := bot.SendMessage(ctx.EffectiveChat.Id, "Hi, I am started.", &gotgbot.SendMessageOpts{
		ReplyToMessageId: ctx.EffectiveMessage.MessageId,
	})
	return err
}
