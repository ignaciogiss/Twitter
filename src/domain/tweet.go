package domain

import (
	"fmt"
	"time"
)

const INIT_ID = 0

type Tweet struct {
	Id int
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	//var nowDate *time.Time = time.Now()
	nowDate := time.Now()


	return &Tweet{Id: INIT_ID, User: user, Text: text, Date: &nowDate }
}

func (tweet Tweet ) PrintableTweet() string {

	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}