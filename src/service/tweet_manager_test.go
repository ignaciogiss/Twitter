package service_test

import (
    _ "github.com/abiosoft/ishell"
    "github.com/ignaciogiss/twitter/src/domain"
    "github.com/ignaciogiss/twitter/src/service"
    "testing"
)


func isValidTweet(t *testing.T, tweet *domain.Tweet, user string, text string) bool {
    if tweet == nil {
        t.Error("Tweet is nil")
        return false
    }
    if tweet.User != user {
        t.Error("Expected user is ", user)
        return false
    }
    if tweet.Text != text {
        t.Error("Expected text is ", text)
        return false
    }
    return true
}


/*
func TestPublishedTweetIsSave(t *testing.T){
    var tweet string = "This is my first tweet"

    service.PublishTweet(tweet)

    if (service.GetTweet() != tweet) {
        t.Error("Expected tweet is", tweet)
    }
}
*/


/*
func TestPublishedTweetIsSaved( t *testing.T) {
    tweetManager := service.NewTweetManager()
    var tweet *domain.Tweet
    user := "ignaciogiss"
    text := "this is a tweet"
    tweet = domain.NewTweet(user, text)

    tweetManager.PublishTweet(tweet)

    publishedTweet := tweetManager.GetTweet()
    if publishedTweet.User != user &&
        publishedTweet.Text != text {
        t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
            user, text, publishedTweet.User, publishedTweet.Text)
    }
    if publishedTweet.Date == nil {
        t.Error("Expected date can't be nil")
    }

}
*/

func TestTweetWithoutUserIsntPublished( t *testing.T) {
    tweetManager := service.NewTweetManager()

    var tweet *domain.Tweet

    var user string

    text := "this is a tweet"
    tweet = domain.NewTweet(user, text)

    var err error
    _, err = tweetManager.PublishTweet(tweet)

    if err != nil && err.Error() != "user is required" {
        t.Error("Expected error user is required")
    }

}


func TestTweetWithoutTextIsNotPublished( t *testing.T) {
    tweetManager := service.NewTweetManager()

    var tweet *domain.Tweet

    var text string

    user := "usuario"
    tweet = domain.NewTweet(user, text)

    var err error
    _, err = tweetManager.PublishTweet(tweet)

    if err != nil && err.Error() != "text is required" {
        t.Error("Expected error text is required")
    }
}

//func TestTweetWichExceeding140CharacterIsNotPublished( t *testing.T) {}

func TestCanPublishAndRetrieveMoreThanOneTweet( t *testing.T) {
    tweetManager := service.NewTweetManager()

    var tweet, secondTweet *domain.Tweet
    user1 := "ignaciogiss"
    user2 := "bot2"
    text1 := "hola es nuevo tweet"
    text2 := "hola mundo this is a bot tweet"
    tweet = domain.NewTweet(user1, text1)
    secondTweet = domain.NewTweet(user2, text2)

    tweetManager.PublishTweet(tweet)
    tweetManager.PublishTweet(secondTweet)


    publishedTweets := tweetManager.GetTweets()
    if len(publishedTweets) != 2 {
        t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
        return
    }

    firstPublishedTweet := publishedTweets[0]
    secondPublishedTweet := publishedTweets[1]


    if !isValidTweet(t, firstPublishedTweet, user1, text1 ) {
        return
    }
    if !isValidTweet(t, secondPublishedTweet, user2, text2 ) {
        return
    }
}


func TestCanRetrieveTweetById(t *testing.T){
    tweetManager := service.NewTweetManager()

    var tweet *domain.Tweet
    var id int

    user := "grupoesfera"
    text := "This is my first tweet"

    tweet = domain.NewTweet(user, text)

    id, _ = tweetManager.PublishTweet(tweet)

    publishedTweet := tweetManager.GetTweetById( id )

    isValidTweet(t, publishedTweet, user,   text )
}

func TestCannotRetrieveTweetByNegativeId(t *testing.T){
    tweetManager := service.NewTweetManager()

    var tweet *domain.Tweet

    user := "grupoesfera"
    text := "This is my first tweet"

    tweet = domain.NewTweet(user, text)

    _, _ = tweetManager.PublishTweet(tweet)


    if (tweetManager.GetTweetById( -1 ) != nil) {
        t.Error("Expected tweet is nil")
    }

}

func TestCannotRetrieveTweetByInvalidId(t *testing.T){
    tweetManager := service.NewTweetManager()

    var tweet *domain.Tweet

    user := "grupoesfera"
    text := "This is my first tweet"

    tweet = domain.NewTweet(user, text)

    _, _ = tweetManager.PublishTweet(tweet)


    if (tweetManager.GetTweetById( 123 ) != nil) {
        t.Error("Expected tweet is nil")
    }

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
    tweetManager := service.NewTweetManager()

    var tweet, secondTweet, thirdTweet *domain.Tweet
    user := "grupoesfera"
    anotherUser := "nick"
    text := "This is my first tweet"
    secondText := "This is my second tweet"
    tweet = domain.NewTweet(user, text)
    secondTweet = domain.NewTweet(user, secondText)
    thirdTweet = domain.NewTweet(anotherUser, text)
    tweetManager.PublishTweet(tweet)
    tweetManager.PublishTweet(secondTweet)
    tweetManager.PublishTweet(thirdTweet)
    // Operation
    count := tweetManager.CountTweetsByUser(user)
    // Validation
    if count != 2 {
        t.Errorf("Expected count is 2 but was %d", count)
    }
}

func TestCountTheTweetsSentByAnInvalidUserIsZero(t *testing.T) {
    tweetManager := service.NewTweetManager()

    var tweet, secondTweet, thirdTweet *domain.Tweet
    user := "grupoesfera"
    anotherUser := "nick"
    text := "This is my first tweet"
    secondText := "This is my second tweet"
    tweet = domain.NewTweet(user, text)
    secondTweet = domain.NewTweet(user, secondText)
    thirdTweet = domain.NewTweet(anotherUser, text)
    tweetManager.PublishTweet(tweet)
    tweetManager.PublishTweet(secondTweet)
    tweetManager.PublishTweet(thirdTweet)
    // Operation
    user = ""
    count := tweetManager.CountTweetsByUser(user)
    // Validation
    if count != 0 {
        t.Errorf("Expected count is 0 but was %d", count)
    }
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
    tweetManager := service.NewTweetManager()

    var tweet, secondTweet, thirdTweet *domain.Tweet
    user := "grupoesfera"
    anotherUser := "nick"
    text := "This is my first tweet"
    secondText := "This is my second tweet"
    tweet = domain.NewTweet(user, text)
    secondTweet = domain.NewTweet(user, secondText)
    thirdTweet = domain.NewTweet(anotherUser, text)
    // publish the 3 tweets
    tweetManager.PublishTweet(tweet)
    tweetManager.PublishTweet(secondTweet)
    tweetManager.PublishTweet(thirdTweet)

    // Operation
    tweets := tweetManager.GetTweetsByUser(user)

    // Validation
    if len(tweets) != 2 { /* handle error */ }
    firstPublishedTweet := tweets[0]
    secondPublishedTweet := tweets[1]
    if !isValidTweet(t, firstPublishedTweet, user, text ) {
        return
    }
    if !isValidTweet(t, secondPublishedTweet, user, secondText ) {
        return
    }
    if !isValidTweet(t, thirdTweet, anotherUser, text ) {
        return
    }
}

func TestCannotRetrieveTheTweetsSentByAnInvalidUser(t *testing.T) {
    tweetManager := service.NewTweetManager()

    var tweet, secondTweet, thirdTweet *domain.Tweet
    user := "grupoesfera"
    anotherUser := "nick"
    text := "This is my first tweet"
    secondText := "This is my second tweet"
    tweet = domain.NewTweet(user, text)
    secondTweet = domain.NewTweet(user, secondText)
    thirdTweet = domain.NewTweet(anotherUser, text)
    // publish the 3 tweets
    tweetManager.PublishTweet(tweet)
    tweetManager.PublishTweet(secondTweet)
    tweetManager.PublishTweet(thirdTweet)

    // Operation
    user = ""
     tweetManager.GetTweetsByUser(user)

    if (tweetManager.GetTweetsByUser(user) != nil) {
        t.Error("Expected tweet is nil")
    }

}