package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	ENV        bool
	TOKEN      string
	DB_URI     string
	RULE_KEY   string
	TG_API_URL string
	CHAT_ID    int64
)

func init() {
	if os.Getenv("ENV") == "" {
		Vi := viper.New()
		Vi.SetConfigFile("config/config.yaml")
		Vi.ReadInConfig()
		TOKEN = Vi.Get("bot-token").(string)
		DB_URI = Vi.Get("db-uri").(string)
		RULE_KEY = Vi.Get("rule-key").(string)
		TG_API_URL = Vi.Get("tg-api-url").(string)
		CHAT_ID = Vi.GetInt64("chat-id")
	} else {
		TOKEN = os.Getenv("TOKEN")
		DB_URI = os.Getenv("DB_URI")
		RULE_KEY = os.Getenv("RULE_KEY")
		TG_API_URL = os.Getenv("TG_API_URL")
		CHAT_ID, _ = strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 0)
	}
}
