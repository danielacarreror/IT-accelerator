package service

import (
  "fmt"
  "github.com/danielacarrero/Twitter/src/domain"
)

type TweetManager struct {
  Tweets []*domain.Tweet
  lastTweet *domain.Tweet
  userTweets map[string][]*domain.Tweet
}

/*
var Tweets []*domain.Tweet
var lastTweet *domain.Tweet
var userTweets map[string][]*domain.Tweet
*/

func NewTweetManager() TweetManager{
  tweets := make([]*domain.Tweet, 0)
  userTweets := make(map[string][]*domain.Tweet)
  return TweetManager{tweets, nil, userTweets}
}

func (tm *TweetManager) PublishTweet(tweet *domain.Tweet) (int, error){
  if(tweet.User == ""){
    return 0, fmt.Errorf("user is required")
  }
  if(tweet.Text == ""){
    return 0, fmt.Errorf("text is required")
  }
  if(len(tweet.Text) > 140){
    return 0, fmt.Errorf("tweet exceeding 140 characters")
  }

  tweetsFromUser := tm.userTweets[tweet.User]
  tweetsFromUser = append(tweetsFromUser, tweet)
  tm.userTweets[tweet.User] = tweetsFromUser

  tm.Tweets = append(tm.Tweets, tweet)
  tm.lastTweet = tweet

  return tweet.Id, nil
}

func (tm *TweetManager) GetLastTweet() *domain.Tweet {
  return tm.lastTweet;
}

func (tm *TweetManager) GetTweets() []*domain.Tweet{
  return tm.Tweets
}

func (tm *TweetManager) GetTweetById(id int) *domain.Tweet {
  for _ , valor := range tm.Tweets {
    if valor.Id == id {
      return valor
    }
  }
  return nil
}

func (tm *TweetManager) CountTweetsByUser(user string) int{
  counter := 0
  for _ , valor := range tm.Tweets {
    if valor.User == user {
      counter ++
    }
  }
  return counter
}

func (tm *TweetManager) GetTweetsByUser(user string) []*domain.Tweet{
  for clave, valor := range(tm.userTweets) {
    if clave == user{
      return valor
    }
  }
  return nil
}
