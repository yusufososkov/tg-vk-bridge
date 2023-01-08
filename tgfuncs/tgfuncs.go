package tgfuncs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"tg-vk-bridge/tokens"
)

func GetUpdates() string {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", tokens.TgToken))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func SendMessage(txtMsg string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=1866924348&text=%s", tokens.TgToken, txtMsg))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
