package service

import (
	"fmt"
	"github.com/ignaciogiss/twitter/src/domain"
)



func (tweetManager TweetManager ) GetTweetById(id int ) *domain.Tweet {
	if id < 0 {
		return nil
	}

	for _, ts := range tweetManager.tweets {
		for _, t := range ts {
			if t.Id == id {
				return t
		    }
		}
	}

	return nil
}

func (tweetManager TweetManager ) GetTweetsByUser( user string ) []*domain.Tweet {
	if user == "" {
		return nil
	}

	return tweetManager.tweets[user]
}

func (tweetManager TweetManager ) CountTweetsByUser( user string ) int {
	if user == "" {
		return 0
	}

	return len(tweetManager.tweets[user])
}

func (tweetManager TweetManager ) PublishTweet(t *domain.Tweet) (int, error ) {
    if len( t.User) == 0 {
		return -1, fmt.Errorf("user is required" )
	}
	if len( t.Text) == 0 {
		return -1, fmt.Errorf("text is required" )
	}

    t.Id = tweetManager.lastId
	tweetManager.lastId++
	tweetManager.tweet = t  // mantinene retro con primeros test de unico tweet
    _, existe := tweetManager.tweets[t.User]
    if !existe {
		tweetManager.tweets[t.User] = make([]*domain.Tweet, 0)
	}
	tweetManager.tweets[t.User] = append( tweetManager.tweets[t.User], t)
	return t.Id, nil
}

func (tweetManager TweetManager ) GetTweets() []*domain.Tweet{
	var allTweets []*domain.Tweet

	for _, t := range tweetManager.tweets {
		allTweets = append(allTweets, t...)
	}
	return allTweets
}


func (tweetManager TweetManager ) GetTweet() *domain.Tweet{
	return tweetManager.tweet
}