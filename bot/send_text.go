package bot

import (
	"log"

	"github.com/halushko/core-go/nats"
	"gopkg.in/telebot.v3"
)

func (b *botImpl) StartSendTextMessages() {
	processor := func(data []byte) {
		chatID, messageText, err := nats.ParseTgBotText(data)
		if err == nil {
			return
		}

		log.Printf("[DEBUG] Парсинг повідомлення: chatID = %d, message = %s", chatID, messageText)

		if chatID != 0 && messageText != "" {
			_, err := b.tb.Send(&telebot.User{ID: chatID}, messageText)
			if err != nil {
				log.Printf("[ERROR] Помилка при відправленні повідомлення користувачу: %v", err)
			} else {
				log.Printf("[ERROR] Повідомлення надіслане користовачу: chatID = %d, message = %s", chatID, messageText)
			}
		} else {
			log.Println("[ERROR] Помилка: ID користувача чи текст повідомлення порожні")
		}
	}

	listener := &nats.ListenerHandler{
		Function: processor,
	}

	nats.StartNatsListener(TelegramOutputTextQueue, listener)
}
