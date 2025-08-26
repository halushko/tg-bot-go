package bot

import (
	"log"

	"github.com/halushko/core-go/nats"
	"gopkg.in/telebot.v3"
)

func (b *botImpl) StartHandleDocumentMessages() {
	b.tb.Handle(telebot.OnDocument, func(c telebot.Context) error {
		document := c.Message().Document

		log.Printf("[DEBUG] Отримано файл: %s", document.FileName)

		chatId := c.Chat().ID
		fileID := document.FileID
		fileName := document.FileName
		fileSize := document.FileSize
		mimeType := document.MIME
		messageText := c.Message().Text
		caption := c.Message().Caption
		url := document.FileURL

		log.Printf(
			"[DEBUG] chatId:%d, uploadedFileId:%s, fileName:%s, message:%s, caption:%s",
			chatId, fileID, fileName, messageText, caption,
		)

		nats.PublishTgFileInfoMessage(TelegramInputFileQueue, chatId, fileID, fileName, fileSize, mimeType, url)
		//nats.PublishTgTextMessage(TelegramOutputTextQueue, chatId, "Файл "+document.FileName+" додано до обробки")
		return nil
	})
}
