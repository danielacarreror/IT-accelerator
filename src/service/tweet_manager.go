package service

import (
  "fmt"
  "github.com/danielacarrero/Twitter/src/domain"
)

var Tweets []*domain.Tweet
var lastTweet *domain.Tweet

func InitializeService(){
  Tweets = make([]*domain.Tweet, 0)
}

func PublishTweet(tweet *domain.Tweet) error{
  if(tweet.User == ""){
    return fmt.Errorf("user is required")
  }
  if(tweet.Text == ""){
    return fmt.Errorf("text is required")
  }
  if(len(tweet.Text) > 140){
    return fmt.Errorf("tweet exceeding 140 characters")
  }
  Tweets = append(Tweets, tweet)
  lastTweet = tweet
  return nil
}

func GetLastTweet() *domain.Tweet {
  return lastTweet;
}

func GetTweets() []*domain.Tweet{
  return Tweets
}
