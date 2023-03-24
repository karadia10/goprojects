package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("BOT_TOKEN", "bot_token")
	os.Setenv("APP_TOKEN", "app_token")
	os.Setenv("CHANNEL_ID", "channel_id")
	bot:=slacker.NewClient(os.Getenv("BOT_TOKEN"), os.Getenv("APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples : []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err:=strconv.Atoi(year)
			if err !=nil {
				fmt.Println(err)
			}
			age := 2023-yob
			r:=fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel:=context.WithCancel(context.Background())
	defer cancel()
	err:=bot.Listen(ctx)
	if err!= nil {
		log.Fatal(err)
	}

}

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent) {
	for event:=range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)

	}
}