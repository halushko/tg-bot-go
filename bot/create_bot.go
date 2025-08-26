package bot

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

//goland:noinspection GoUnusedExportedFunction
func Create() (Bot, error) {
	token := os.Getenv(envBotToken)
	if token == "" {
		return nil, fmt.Errorf("[ERROR] Необхідно задати токен боту в env %s", envBotToken)
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 1 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &botImpl{tb: bot}, nil
}
