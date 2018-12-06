package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/danielacarrero/Twitter/src/service"
	"github.com/danielacarrero/Twitter/src/domain"
)

func main() {

	tweetManager := service.NewTweetManager()
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)


			c.Print("Write your tweet: ")
			user := "user"
			text := c.ReadLine()

			tweetManager.PublishTweet(domain.NewTweet(user, text))

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Show all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()
			fmt.Println(len(tweets))

			for i := 0; i < len(tweets); i++{
				c.Println(tweets[i].PrintableTweet())
			}

			return
		},
	})

	shell.Run()

}
