package bot

import (
	"strconv"

	"github.com/halushko/core-go/nats"
	"gopkg.in/telebot.v3"
)

func (b *botImpl) StartHandleMemberJoined() {
	b.tb.Handle(telebot.OnUserJoined, func(c telebot.Context) error {
		chatID := c.Chat().ID
		members := c.Message().UsersJoined

		for _, u := range members {
			nats.PublishTgTextMessage(TelegramMemberJoinedQueue, chatID, strconv.FormatInt(u.ID, 10))
		}
		return nil
	})
}
