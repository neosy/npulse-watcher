package ntelegram

import (
	"log"
	"net/http"
	"net/url"
)

const (
	TG_URL_BOT string = "https://api.telegram.org/bot"
)

type Telegram struct {
	token string
}

func New(token string) *Telegram {
	return &Telegram{token}
}

func (tg Telegram) Send(channel_id string, text string) error {
	uri := TG_URL_BOT
	uri += tg.token
	uri += "/sendMessage"

	data := url.Values{
		"chat_id": {channel_id},
		"text":    {text},
	}
	_, err := http.PostForm(uri, data)

	if err != nil {
		log.Println(err)
	}

	return err
}
