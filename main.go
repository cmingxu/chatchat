package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:   os.Getenv("TOKEN"),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: true,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(c.Message().Payload)
	})

	b.Handle("/get-image", func(c tele.Context) error {
		fmt.Printf("%#v\n", c)
		fmt.Printf("%#v\n", c.Message())
		fmt.Printf("%#v\n", c.Sender())

		// return c.Send(c.Sender(), f)
		//
		return c.Send("Acc")
	})

	b.Start()
}
