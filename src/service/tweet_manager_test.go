package service_test


import (
  "testing"
  "github.com/danielacarrero/Twitter/src/service"
  "github.com/danielacarrero/Twitter/src/domain"
)

//Se debe empezar con Test en la función para indicar que es test
// Debe recibir como parametro *testing.T, importando del package testing el tipo T, este tipo es el definido para test
// t es un puntero al tipo T
func TestPublishedTweetIsSaved(t *testing.T){
  //Initialization
  var tweet *domain.Tweet
  user := "danielacarrero"
  text := "This is my first tweet"
  tweet = domain.NewTweet(user, text)

  //Operation
  service.PublishTweet(tweet)

  //Validation

  /*if service.Tweet != tweet {
    t.Error("Expected tweet is", tweet)
  }*/

  publishedTweet := service.GetLastTweet()
  if publishedTweet.User != user || publishedTweet.Text != text {
    t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
  }

  if publishedTweet.Date == nil {
    t.Error("Expected date can't be nil")
  }
}

func TestTweetWithoutUserIsNotPublished(t *testing.T){
  //Initialization
  var tweet *domain.Tweet

  var user string
  text := "This is my first tweet"

  tweet = domain.NewTweet(user, text)

  //Operation
  var err error
  err = service.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "user is required" {
    t.Error("Expected error is user is required")
  }
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
  //Initialization
  var tweet *domain.Tweet

  var text string
  user := "danielacarrero"

  tweet = domain.NewTweet(user, text)

  //Operation
  var err error
  err = service.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "text is required" {
    t.Error("Expected error is text is required")
  }
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T){
  //Initialization
  var tweet *domain.Tweet

  text := `asdasdasdslajflkdkkkkkkaslkdjalkdjlaskjdlkajdl
            kjsdljdlkajdlkasjlksajflksjflksajflkjlkajsdlk
            jlkajsdlkjaflkjlkjdlkajdlajdlajdlakjdlkajdlak
            jflkasjdlkjsdlkjsadlkafjlakdjlkasjflkasjdljflkdljlaa`
  user := "danielacarrero"

  tweet = domain.NewTweet(user, text)

  //Operation
  var err error
  err = service.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "tweet exceeding 140 characters" {
    t.Error("Expected error is tweet exceeding 140 characters")
  }
}


func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T){
  //Initialization
  service.InitializeService()
  var tweet, secondTweet *domain.Tweet

  user := "danielacarrero"
  text1 := "My first tweet"
  text2 := "My second tweet"

  tweet = domain.NewTweet(user, text1)
  secondTweet = domain.NewTweet(user, text2)

  //Operation
  service.PublishTweet(tweet)
  service.PublishTweet(secondTweet)

  //Validation
  publishedTweets := service.GetTweets()
  if len(publishedTweets) != 2 {
    t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
    return
  }

  firstPublishedTweet := publishedTweets[0]
  secondPublishedTweet := publishedTweets[1]

  if !isValidTweet(t, firstPublishedTweet, "danielacarrero", "My first tweet"){
    return
  }
  if !isValidTweet(t, secondPublishedTweet, "danielacarrero", "My second tweet"){
    return
  }
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, user string, text string) bool{
  if tweet.User != user {
    return false
  }
  if tweet.Text != text{
    return false
  }
  if len(tweet.Text) > 140{
    return false
  }
  return true
}
