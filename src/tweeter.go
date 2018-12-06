package main

import (
	"github.com/abiosoft/ishell"
	"github.com/ignaciogiss/twitter/src/domain"
	"github.com/ignaciogiss/twitter/src/service"
	"strconv"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			var tweet *domain.Tweet

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet = domain.NewTweet(user, text)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println("Dia:     ", tweet.Date )
			c.Println("Usuario: ", tweet.User)
			c.Println("Tweet:   ", tweet.Text)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetById",
		Help: "Retrieves a tweet by its ID",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			var tweet *domain.Tweet

			c.Print("Search ID: ")

			id_str := c.ReadLine()
			id, _ := strconv.Atoi(id_str )

			tweet = service.GetTweetById(id)

			c.Println("Dia:     ", tweet.Date )
			c.Println("Usuario: ", tweet.User)
			c.Println("Tweet:   ", tweet.Text)

			return
		},
	})

	shell.Run()

}



