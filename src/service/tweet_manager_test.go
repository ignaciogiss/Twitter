package service_test

import ( "testing"
    _ "github.com/abiosoft/ishell"
         "github.com/ignaciogiss/Twitter/src/service"
)

func TestPublishedTweetIsSave(t *testing.T){
    var tweet string = "This is my first tweet"

    service.PublishTweet(tweet)

    if (service.GetTweet() != tweet) {
        t.Error("Expected tweet is", tweet)
    }
}
