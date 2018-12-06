package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	//var nowDate *time.Time = time.Now()
	nowDate := time.Now()

	if text == "" {
		text = "lakds"
	}

	return &Tweet{User: user, Text: text, Date: &nowDate }
}