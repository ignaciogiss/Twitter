package service

import (
	"fmt"
	"github.com/ignaciogiss/twitter/src/domain"
)
var tweet *domain.Tweet;
var tweets []*domain.Tweet;


func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
}



func PublishTweet(t *domain.Tweet) error{
    if len( t.User) == 0 {
		return fmt.Errorf("user is required" )
	}
	if len( t.Text) == 0 {
		return fmt.Errorf("text is required" )
	}
    tweet = t
	tweets = append( tweets, t)
	return nil
}

func GetTweets() []*domain.Tweet{
	return tweets
}


func GetTweet() *domain.Tweet{
	return tweet
}