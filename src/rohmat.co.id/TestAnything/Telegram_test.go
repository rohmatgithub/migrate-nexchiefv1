package TestAnything

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"testing"
)

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", "5535654220:AAFq7yN5BTj6tXodsSGIEN8ItwkKdTZtGkQ")
}

func TestTelegram(t *testing.T)  {
	bot, err := tgbotapi.NewBotAPI("5535654220:AAFq7yN5BTj6tXodsSGIEN8ItwkKdTZtGkQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//tgbotapi.NewMessageToChannel()
	msg := tgbotapi.NewMessage(819843849, "test")
			_, _ = bot.Send(msg)

	//url := fmt.Sprintf("%s/sendMessage", getUrl())
	//body, _ := json.Marshal(map[string]string{
	//	"chat_id": "",
	//	"text":    text,
	//})
	//response, err = http.Post(
	//	url,
	//	"application/json",
	//	bytes.NewBuffer(body),
	//)
	//if err != nil {
	//	return false, err
	//}
	//
	//// Close the request at the end
	//defer response.Body.Close()
	//
	//// Body
	//body, err = ioutil.ReadAll(response.Body)
	//if err != nil {
	//	return false, err
	//}

	//u := tgbotapi.NewUpdate(0)
	//u.Timeout = 60
	//
	//updates := bot.GetUpdatesChan(u)
	//
	//for update := range updates {
	//	if update.Message != nil { // If we got a message
	//		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//		msg.ReplyToMessageID = update.Message.MessageID
	//
	//		_, _ = bot.Send(msg)
	//	}
	//}
}
