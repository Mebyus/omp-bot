package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyApartmentCommander) Get(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	return
}