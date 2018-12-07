package service

import (
  "fmt"
  "strings"
  "github.com/danielacarrero/Twitter/src/domain"
)

type TweetManager struct {
  Tweets []domain.Tweet
  lastTweet domain.Tweet
  userTweets map[string][]domain.Tweet
  tweetWriter TweetWriter
}

/*
var Tweets []*domain.Tweet
var lastTweet *domain.Tweet
var userTweets map[string][]*domain.Tweet
*/

func NewTweetManager(tweetWriter TweetWriter) TweetManager{
  tweets := make([]domain.Tweet, 0)
  userTweets := make(map[string][]domain.Tweet)
  return TweetManager{tweets, nil, userTweets, tweetWriter}
}

func (tm *TweetManager) PublishTweet(tweet domain.Tweet) (int, error){
  if(tweet.GetUser() == ""){
    return 0, fmt.Errorf("user is required")
  }
  if(tweet.GetText() == ""){
    return 0, fmt.Errorf("text is required")
  }
  if(len(tweet.GetText()) > 140){
    return 0, fmt.Errorf("tweet exceeding 140 characters")
  }

  switch castedTweet := tweet.(type) {
    case *domain.TextTweet:
    case *domain.ImageTweet:
      if castedTweet.GetUrlImg() == "" {
        return 0, fmt.Errorf("Image tweet without url image")
      }
    case *domain.QuoteTweet:
      if castedTweet.QuotedTweet == nil {
        return 0, fmt.Errorf("QuoteTweet without a quote")
      }
  }

  tweetsFromUser := tm.userTweets[tweet.GetUser()]
  tweetsFromUser = append(tweetsFromUser, tweet)
  tm.userTweets[tweet.GetUser()] = tweetsFromUser

  tm.Tweets = append(tm.Tweets, tweet)
  tm.lastTweet = tweet

  tm.tweetWriter.WriteTweet(tweet)

  return tweet.GetId(), nil
}

func (tm *TweetManager) GetLastTweet() domain.Tweet {
  return tm.lastTweet;
}

func (tm *TweetManager) GetTweets() []domain.Tweet{
  return tm.Tweets
}

func (tm *TweetManager) GetTweetById(id int) domain.Tweet {
  for _ , valor := range tm.Tweets {
    if valor.GetId() == id {
      return valor
    }
  }
  return nil
}

func (tm *TweetManager) CountTweetsByUser(user string) int{
  counter := 0
  for _ , valor := range tm.Tweets {
    if valor.GetUser() == user {
      counter ++
    }
  }
  return counter
}

func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet{
  for clave, valor := range(tm.userTweets) {
    if clave == user{
      return valor
    }
  }
  return nil
}

func (tm *TweetManager) SearchTweetsContaining(query string, searchResult chan domain.Tweet){
  go func() {
    for _, tweet := range tm.Tweets{
      if strings.Contains(tweet.GetText(), query){
        searchResult <- tweet
      }
    }
  }()
}
