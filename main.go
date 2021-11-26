package main

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/slack-go/slack"
	"os"
	"time"
)

func main() {
	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	client := slack.New(token, slack.OptionDebug(true))
	c := cron.New()

	c.AddFunc("0 0 18 * * FRI", func(){
		_, _, err := client.PostMessage(
			channelID,
			slack.MsgOptionAttachments(slack.Attachment{
				ImageURL: "https://qovery.slack.com/files/U02NH5MTZKQ/F02NJACA6CW/james-bond-the-weeknd.gif",
			}),
		)

		if err != nil {
			fmt.Sprint(err.Error())
		}

	})

	c.Start()

	for {
		fmt.Printf("Next weekend : %s", c.Entries()[0].Next.String())
		time.Sleep(12 * time.Hour)
	}
}