package main

import (
 "log"
 "time"

 tele "gopkg.in/telebot.v3"
)

func Start(c tele.Context) error {
  msg := c.Message()
  return c.Send("Welcome " + msg.Sender.Username)
}

func main() {
 tokenBot := "YOUR_TOKEN"
  
 pref := tele.Settings{
		Token: tokenBot,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

 b, err := tele.NewBot(pref)
 if err != nil {
		log.Fatal(err)
 }
	
 b.Handle("/start", Start)
	
 log.Println("[*] Bot running...")
 b.Start()
}