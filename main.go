package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Date struct {
	Name          string
	Date          time.Time
	GentitiveName string
}

func declOfNum(number int, titles []string) string {
	if number < 0 {
		number = -number
	}

	cases := []int{2, 0, 1, 1, 1, 2}
	var currentCase int
	if number%100 > 4 && number%100 < 20 {
		currentCase = 2
	} else if number%10 < 5 {
		currentCase = cases[number%10]
	} else {
		currentCase = cases[5]
	}

	return titles[currentCase]
}

func tillString(today, till time.Time) string {
	parts := make([]string, 0, 3)

	diff := till.Sub(today)

	days := int(diff.Hours()) / 24
	hours := int(diff.Hours()) % 24
	mins := int(diff.Minutes()) % 60

	if days != 0 {
		parts = append(parts, fmt.Sprintf("%d %s", days, declOfNum(days, []string{"день", "дня", "дней"})))
	}
	parts = append(parts, fmt.Sprintf("%d %s", hours, declOfNum(hours, []string{"час", "часа", "часов"})))
	parts = append(parts, fmt.Sprintf("%d %s", mins, declOfNum(mins, []string{"минута", "минуты", "минут"})))

	return strings.Join(parts, " ")
}

var dates = []Date{
	{Date: time.Date(0000, 6, 01, 00, 00, 00, 00, time.Local), Name: "Лето", GentitiveName: "лета"},
	{Date: time.Date(0000, 1, 01, 00, 00, 00, 00, time.Local), Name: "Новый год", GentitiveName: "нового года"},
	{Date: time.Date(0000, 9, 01, 00, 00, 00, 00, time.Local), Name: "Первое сентября", GentitiveName: "первого сентября"},
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	ctx := context.Background()

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatal(err)
	}

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatal(err)
	}
	defer bh.Stop()

	bh.HandleInlineQuery(func(ctx *th.Context, query telego.InlineQuery) error {
		results := make([]telego.InlineQueryResult, len(dates))
		now := time.Now()

		for i, date := range dates {
			date.Date = date.Date.AddDate(now.Year(), 0, 0)
			if now.Sub(date.Date) > 0 {
				date.Date = date.Date.AddDate(1, 0, 0)
			}

			results[i] = tu.ResultArticle(
				date.Name,
				date.Name,
				tu.TextMessage(
					fmt.Sprintf(
						"До %s осталось %s",
						date.GentitiveName,
						tillString(now, date.Date),
					),
				),
			)
		}

		return ctx.Bot().AnswerInlineQuery(ctx, tu.InlineQuery(query.ID, results...).WithCacheTime(0))
	})

	bh.Start()

}
