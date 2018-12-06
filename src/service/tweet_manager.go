package service

import (
	"fmt"
	"github.com/ignaciogiss/twitter/src/domain"
)
var tweet *domain.Tweet;
var tweets map[string][]*domain.Tweet;

var lastId = 0


func InitializeService() {
	tweets = make(map[string][]*domain.Tweet)
}

func GetTweetById(id int ) *domain.Tweet {
	if id < 0 {
		return nil
	}

	for _, ts := range tweets {
		for _, t := range ts {
			if t.Id == id {
				return t
		    }
		}
	}

	return nil
}

func GetTweetsByUser( user string ) []*domain.Tweet {
	if user == "" {
		return nil
	}

	return tweets[user]
}

func CountTweetsByUser( user string ) int {
	if user == "" {
		return 0
	}

	return len(tweets[user])
}

func PublishTweet(t *domain.Tweet) (int, error ) {
    if len( t.User) == 0 {
		return -1, fmt.Errorf("user is required" )
	}
	if len( t.Text) == 0 {
		return -1, fmt.Errorf("text is required" )
	}
    t.Id = lastId
    lastId++
    tweet = t  // mantinene retro con primeros test de unico tweet
    _, existe := tweets[t.User]
    if !existe {
		tweets[t.User] = make([]*domain.Tweet, 0)
	}
	tweets[t.User] = append( tweets[t.User], t)
	return t.Id, nil
}

func GetTweets() []*domain.Tweet{
	var allTweets []*domain.Tweet

	for _, t := range tweets {
		allTweets = append(allTweets, t...)
	}
	return allTweets
}


func GetTweet() *domain.Tweet{
	return tweet
}