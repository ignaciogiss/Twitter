package domain

import (
	"fmt"
	"time"
)

const INIT_ID = 0

type Tweet interface {
	PrintableTweet() string
	GetDate() *time.Time
	GetUser() string
	GetId() int
	SetId(id int)
	GetText() string
}

type TextTweet struct {
	Id int
	User string
	Text string
	Date *time.Time
}

func NewTextTweet(user string, text string) *TextTweet {
	//var nowDate *time.Time = time.Now()
	nowDate := time.Now()

	return &TextTweet{Id: INIT_ID, User: user, Text: text, Date: &nowDate }
}

func (tweet TextTweet ) PrintableTweet() string {

	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func (tweet TextTweet ) GetId() int {
	return tweet.Id
}
func (tweet TextTweet ) GetDate() *time.Time {
	return tweet.Date
}
func (tweet *TextTweet ) SetId(id int) {
	tweet.Id = id
}
func (tweet TextTweet ) GetUser() string {
	return tweet.User
}
func (tweet TextTweet ) GetText() string {
	return tweet.Text
}

func (tweet TextTweet ) String() string {

	return tweet.PrintableTweet()
}


type ImageTweet struct {
	TextTweet
	ImgUrl string
}

func NewImageTweet(user string, text string, imgUrl string) *ImageTweet {

	return &ImageTweet{ TextTweet: *NewTextTweet(user, text), ImgUrl: imgUrl }
}

func (tweet ImageTweet ) PrintableTweet() string {

	return fmt.Sprintf("@%s: This is my image %s", tweet.User, tweet.ImgUrl )
}


type QuoteTweet struct {
	TextTweet
	secondTweet *TextTweet
}

func NewQuoteTweet(user string, text string, quotedTweet *TextTweet) *QuoteTweet {

	return &QuoteTweet{ TextTweet: *quotedTweet, secondTweet: NewTextTweet(user, text) }
}

func (tweet QuoteTweet ) PrintableTweet() string {

	return fmt.Sprintf("@%s: %s \"%s\"", tweet.secondTweet.User, tweet.secondTweet.Text, tweet.TextTweet.PrintableTweet() )
}