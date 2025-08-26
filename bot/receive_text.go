package bot

import (
	"log"

	"github.com/halushko/core-go/nats"
	"gopkg.in/telebot.v3"
)

func (b *botImpl) StartHandleTextMessages() {
	b.tb.Handle(telebot.OnText, func(c telebot.Context) error {
		chatId := c.Chat().ID
		message := c.Message().Text

		log.Printf("[DEBUG] chatId:%d, message:%s", chatId, message)
		nats.PublishTgTextMessage(TelegramInputTextQueue, chatId, message)
		//nats.PublishTgTextMessage(TelegramOutputTextQueue, chatId, "Ваше повідомлення "+message+" додано до обробки")

		return nil
	})
}
