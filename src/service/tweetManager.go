package service

import (
	"github.com/ignaciogiss/twitter/src/domain"
)

// func (tweetManager TweetManager ) InitializeService() {
// 	tweetManager.tweets = make(map[string][]*domain.Tweet)
// }

type TweetManager struct {
	tweet domain.Tweet
	tweets map[string][]domain.Tweet
	lastId int
	registeredUsers []*domain.User
	TweetWriter
}


func  NewTweetManager( writer TweetWriter ) *TweetManager {

	return &TweetManager{ TweetWriter: writer, tweets : make(map[string][]domain.Tweet), lastId: 0 }
}
