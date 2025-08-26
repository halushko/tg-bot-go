package bot

import (
	"gopkg.in/telebot.v3"
)

const TelegramInputTextQueue = "TELEGRAM_INPUT_TEXT_QUEUE"
const TelegramOutputTextQueue = "TELEGRAM_OUTPUT_TEXT_QUEUE"
const TelegramInputFileQueue = "TELEGRAM_INPUT_FILE_QUEUE"

const envBotToken = "BOT_TOKEN"

type telegramMessage struct {
	ChatID int64
	Text   string
}

type Bot interface {
	Start()
	Stop()
	StartHandleTextMessages()
	StartHandleDocumentMessages()
	StartSendTextMessages()
}

type botImpl struct {
	tb *telebot.Bot
}

func (b *botImpl) Start() { b.tb.Start() }
func (b *botImpl) Stop()  { b.tb.Stop() }
