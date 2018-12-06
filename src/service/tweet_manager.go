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

func GetTweetById(id int ) *domain.Tweet {
	if id < 0 {
		return nil
	}
	return tweets[ id ]
}

func PublishTweet(t *domain.Tweet) (int, error ) {
    if len( t.User) == 0 {
		return -1, fmt.Errorf("user is required" )
	}
	if len( t.Text) == 0 {
		return -1, fmt.Errorf("text is required" )
	}
    t.Id = len(tweets) // tamano porque es la posicion final donde se appendea
    tweet = t
	tweets = append( tweets, t)
	return t.Id, nil
}

func GetTweets() []*domain.Tweet{
	return tweets
}


func GetTweet() *domain.Tweet{
	return tweet
}