package service

import (
	"fmt"
	"github.com/ignaciogiss/twitter/src/domain"
	"strings"
)



func (tweetManager TweetManager ) GetTweetById(id int ) domain.Tweet {
	if id < 0 {
		return nil
	}

	for _, ts := range tweetManager.tweets {
		for _, t := range ts {
			if t.GetId() == id {
				return t
		    }
		}
	}

	return nil
}

func (tweetManager TweetManager ) GetTweetsByUser( user string ) []domain.Tweet {
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

func (tweetManager *TweetManager ) PublishTweet(t domain.Tweet) (int, error ) {
    if len( t.GetUser()) == 0 {
		return -1, fmt.Errorf("user is required" )
	}
	if len( t.GetText()) == 0 {
		return -1, fmt.Errorf("text is required" )
	}

    t.SetId( tweetManager.lastId )
	tweetManager.lastId++
	tweetManager.tweet = t  // mantiene retro con primeros test de unico tweet
    _, existe := tweetManager.tweets[t.GetUser()]
    if !existe {
		tweetManager.tweets[t.GetUser()] = make([]domain.Tweet, 0)
	}
	tweetManager.tweets[t.GetUser()] = append( tweetManager.tweets[t.GetUser()], t)
	tweetManager.TweetWriter.WriteTweet( t )
	return t.GetId(), nil
}

func (tweetManager TweetManager ) GetTweets() []domain.Tweet{
	var allTweets []domain.Tweet

	for _, t := range tweetManager.tweets {
		allTweets = append(allTweets, t...)
	}
	return allTweets
}

/*
func (tweetManager TweetManager ) GetTweet() domain.Tweet{
	return tweetManager.tweet
}
*/
func (tweetManager *TweetManager ) RegisterUser(name string, mail string, nick string, password string) bool{
	tweetManager.registeredUsers = append( tweetManager.registeredUsers, domain.NewUser(name , mail , nick , password ))
	return true // TODO: caso false
}

func (tweetManager TweetManager ) LoginUser(name string, password string) bool{
	for _, u := range tweetManager.registeredUsers {
		if u.Name == name {
			if u.PasswordOk(password) {
				return true
			}
		}
	}
     return false
}


func (tweetManager *TweetManager ) SearchTweetsContaining(query string, searchResult chan domain.Tweet) {

	go func() {
		for _, ts := range tweetManager.tweets {
			for _, t := range ts {
				if strings.Contains( t.GetText(), query )	 {
					searchResult <- t
					}
			}
		}
	}()
}
