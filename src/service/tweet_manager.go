package service

import "github.com/ignaciogiss/twitter/src/domain"

var tweet *domain.Tweet;



func PublishTweet(t *domain.Tweet){
	tweet = t;
}

func GetTweet() *domain.Tweet{
	return tweet
}
