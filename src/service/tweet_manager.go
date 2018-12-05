package service

import (
  "fmt"
  "github.com/danielacarrero/Twitter/src/domain"
)

var Tweet domain.Tweet


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
  Tweet = *tweet
  return nil
}

func GetTweet() *domain.Tweet {
  return &Tweet;
}
