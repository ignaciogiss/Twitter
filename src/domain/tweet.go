package domain

import "time"

type date struct {
	Date *time.Time
}


type Tweet struct {
	User string
	Text string
	Date *time.Time
}


func NewTweet(user string, text string) *Tweet {
	//var nowDate *time.Time = time.Now()
	nowDate := time.Now()
	return &Tweet{User: user, Text: text, Date: &nowDate }
}