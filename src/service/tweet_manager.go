package service

import (
  "github.com/danielacarrero/Twitter/src/domain"
)

var Tweet domain.Tweet


func PublishTweet(tweet *domain.Tweet){
  Tweet = *tweet
}

func GetTweet() *domain.Tweet {
  return &Tweet;
}
