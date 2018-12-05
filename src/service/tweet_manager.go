package service

var tweet string;

func PublishTweet(s string){
	tweet = s;
}

func GetTweet() string{
	return tweet
}
