package main

import (
	"github.com/abiosoft/ishell"
	"github.com/danielacarrero/Twitter/src/service"
	"github.com/danielacarrero/Twitter/src/domain"
)

func main() {
	var user string
	tweetManager := service.NewTweetManager()
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "login to Twitter",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("User: ")
			user = c.ReadLine()
			for user == "" {
				c.Println("User is empty, please login with valid user")
				c.Print("User: ")
				user = c.ReadLine()
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			if user == "" {
				c.Println("Please login first: 'login' command")
				return
			}

			c.Print("Write your tweet: ")
			text := c.ReadLine()

			for text == "" {
				c.Println("Your tweet is empty")
				c.Print("Please, write your tweet: ")
				text = c.ReadLine()
			}

			tweetManager.PublishTweet(domain.NewTextTweet(user, text))

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Show all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			if user == "" {
				c.Println("Please login first: 'login' command")
				return
			}

			tweets := tweetManager.GetTweets()

			for i := 0; i < len(tweets); i++{
				c.Println(tweets[i].PrintableTweet())
			}

			return
		},
	})

	shell.Run()

}
