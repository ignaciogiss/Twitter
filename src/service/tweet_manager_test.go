package service_test

import (
    "github.com/ignaciogiss/twitter/src/domain"
    "testing"
    _ "github.com/abiosoft/ishell"
         "github.com/ignaciogiss/twitter/src/service"
)

/*
func TestPublishedTweetIsSave(t *testing.T){
    var tweet string = "This is my first tweet"

    service.PublishTweet(tweet)

    if (service.GetTweet() != tweet) {
        t.Error("Expected tweet is", tweet)
    }
}
*/

func TestPublishedTweetIsSaved( t *testing.T) {
    var tweet *domain.Tweet
    user := "ignaciogiss"
    text := "this is a tweet"
    tweet = domain.NewTweet(user, text)

    service.PublishTweet(tweet)

    publishedTweet := service.GetTweet()
    if publishedTweet.User != user &&
        publishedTweet.Text != text {
        t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
            user, text, publishedTweet.User, publishedTweet.Text)
    }
    if publishedTweet.Date == nil {
        t.Error("Expected date can't be nil")
    }
}