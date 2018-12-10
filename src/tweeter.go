package main

import (
	"github.com/abiosoft/ishell"
	"github.com/ignaciogiss/twitter/src/domain"
	"github.com/ignaciogiss/twitter/src/rest"
	"github.com/ignaciogiss/twitter/src/service"
	"strconv"
)

func main() {
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	restServer := rest.NewRestTweetServer(tweetManager)
	restServer.StartRestServer()


	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			var tweet domain.Tweet

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet = domain.NewTextTweet(user, text)

			tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})
/*
	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetTweet()

			c.Println("Dia:     ", tweet.GetDate() )
			c.Println("Usuario: ", tweet.GetUser() )
			c.Println("Tweet:   ", tweet.GetText() )

			return
		},
	})
*/
	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetById",
		Help: "Retrieves a tweet by its ID",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			var tweet domain.Tweet

			c.Print("Search ID: ")

			id_str := c.ReadLine()
			id, _ := strconv.Atoi(id_str )

			tweet = tweetManager.GetTweetById(id)

			c.Println("Dia:     ", tweet.GetDate() )
			c.Println("Usuario: ", tweet.GetUser())
			c.Println("Tweet:   ", tweet.GetText())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsByUser",
		Help: "Retrieves a tweets by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			var tweets []domain.Tweet

			c.Print("Search User: ")

			user := c.ReadLine()

			tweets = tweetManager.GetTweetsByUser( user )

			for _, tweet := range tweets {
				c.Println("Dia:     ", tweet.GetDate() )
				c.Println("Usuario: ", tweet.GetUser())
				c.Println("Tweet:   ", tweet.GetText())
			}

			return
		},
	})

	shell.Run()

}
