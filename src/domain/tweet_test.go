package domain_test

import "testing"
import "github.com/ignaciogiss/twitter/src/domain"

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)
	text := tweet.PrintableTweet()
	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}


func TestCanSetIdOnTweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	id := 17
	tweet.SetId(id)

	// Validation
	if tweet.GetId() != id {
		t.Errorf("The expected text id is %d but was %d", id, tweet.GetId())
	}

}

func TestCanGetUserFromTweet(t *testing.T) {

	// Initialization
	user := "grupoesfera"
	tweet := domain.NewTextTweet(user, "This is my tweet")

	// Operation


	// Validation
	if tweet.GetUser() != user {
		t.Errorf("The expected text id is %s but was %s", user, tweet.GetUser())
	}

}

func TestCanGetTextFromTweet(t *testing.T) {

	// Initialization
	text := "This is my tweet"
	tweet := domain.NewTextTweet("grupoesfera", text)

	// Operation


	// Validation
	if tweet.GetText() != text {
		t.Errorf("The expected text id is %s but was %s", text, tweet.GetText())
	}

}