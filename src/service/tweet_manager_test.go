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
  tweetManager := service.NewTweetManager()
  var tweet *domain.TextTweet
  user := "danielacarrero"
  text := "This is my first tweet"
  tweet = domain.NewTextTweet(user, text)

  //Operation
  tweetManager.PublishTweet(tweet)

  //Validation

  /*if service.Tweet != tweet {
    t.Error("Expected tweet is", tweet)
  }*/

  publishedTweet := tweetManager.GetLastTweet()
  if publishedTweet.GetUser() != user || publishedTweet.GetText() != text {
    t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.GetUser(), publishedTweet.GetText())
  }

  if publishedTweet.GetDate() == nil {
    t.Error("Expected date can't be nil")
  }
}

func TestTweetWithoutUserIsNotPublished(t *testing.T){
  //Initialization
  tweetManager := service.NewTweetManager()
  var tweet *domain.TextTweet

  var user string
  text := "This is my first tweet"

  tweet = domain.NewTextTweet(user, text)

  //Operation
  var err error
  _ , err = tweetManager.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "user is required" {
    t.Error("Expected error is user is required")
  }
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
  //Initialization
  tweetManager := service.NewTweetManager()
  var tweet *domain.TextTweet

  var text string
  user := "danielacarrero"

  tweet = domain.NewTextTweet(user, text)

  //Operation
  var err error
  _ , err = tweetManager.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "text is required" {
    t.Error("Expected error is text is required")
  }
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T){
  //Initialization
  tweetManager := service.NewTweetManager()
  var tweet *domain.TextTweet

  text := `asdasdasdslajflkdkkkkkkaslkdjalkdjlaskjdlkajdl
            kjsdljdlkajdlkasjlksajflksjflksajflkjlkajsdlk
            jlkajsdlkjaflkjlkjdlkajdlajdlajdlakjdlkajdlak
            jflkasjdlkjsdlkjsadlkafjlakdjlkasjflkasjdljflkdljlaa`
  user := "danielacarrero"

  tweet = domain.NewTextTweet(user, text)

  //Operation
  var err error
  _ , err = tweetManager.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "tweet exceeding 140 characters" {
    t.Error("Expected error is tweet exceeding 140 characters")
  }
}


func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T){
  //Initialization
  tweetManager := service.NewTweetManager()
  var tweet, secondTweet *domain.TextTweet

  user := "danielacarrero"
  text1 := "My first tweet"
  text2 := "My second tweet"

  tweet = domain.NewTextTweet(user, text1)
  secondTweet = domain.NewTextTweet(user, text2)

  //Operation
  tweetManager.PublishTweet(tweet)
  tweetManager.PublishTweet(secondTweet)

  //Validation
  publishedTweets := tweetManager.GetTweets()
  if len(publishedTweets) != 2 {
    t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
    return
  }

  firstPublishedTweet := publishedTweets[0]
  secondPublishedTweet := publishedTweets[1]

  if !isValidTweet(t, firstPublishedTweet, "danielacarrero", "My first tweet"){
    t.Errorf("First tweet is not valid")
    return
  }
  if !isValidTweet(t, secondPublishedTweet, "danielacarrero", "My second tweet"){
    t.Errorf("Second tweet is not valid")
    return
  }
}

func isValidTweet(t *testing.T, tweet domain.Tweet, user string, text string) bool{
  if tweet == nil {
    return false
  }
  if tweet.GetUser() != user {
    return false
  }
  if tweet.GetText() != text{
    return false
  }
  if len(tweet.GetText()) > 140{
    return false
  }
  return true
}

func TestCanRetrieveTweetById(t *testing.T){
  //Initialization
  tweetManager := service.NewTweetManager()

  var tweet *domain.TextTweet
  var id int

  user := "grupoesfera"
  text := "This is my first tweet"

  tweet = domain.NewTextTweet(user, text)

  //Operation
  id, _ = tweetManager.PublishTweet(tweet)

  //Validation
  publishedTweet := tweetManager.GetTweetById(id)

   if !isValidTweet(t, publishedTweet, user, text) {
     t.Errorf("Id is not valid")
     return
   }

   publishedTweet = tweetManager.GetTweetById(1234)

   if publishedTweet != nil{
     t.Errorf("Tweet not existing must return nil")
   }
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
    // Initialization
    tweetManager := service.NewTweetManager()

    var tweet, secondTweet, thirdTweet *domain.TextTweet
    user := "grupoesfera"
    anotherUser := "nick"
    text := "This is my first tweet"
    secondText := "This is my second tweet"

    tweet = domain.NewTextTweet(user, text)
    secondTweet = domain.NewTextTweet(user, secondText)
    thirdTweet = domain.NewTextTweet(anotherUser, text)

    tweetManager.PublishTweet(tweet)
    tweetManager.PublishTweet(secondTweet)
    tweetManager.PublishTweet(thirdTweet)

    // Operation
    count := tweetManager.CountTweetsByUser(user)
    countAnotherUser := tweetManager.CountTweetsByUser(anotherUser)

    // Validation
    if count != 2 {
        t.Errorf("Expected count is 2 but was %d", count)
    }
    if countAnotherUser != 1 {
        t.Errorf("Expected count is 1 but was %d", countAnotherUser)
    }
}


func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
    // Initialization
    tweetManager := service.NewTweetManager()
    var tweet, secondTweet, thirdTweet *domain.TextTweet
    user := "grupoesfera"
    anotherUser := "nick"
    text := "This is my first tweet"
    secondText := "This is my second tweet"
    tweet = domain.NewTextTweet(user, text)
    secondTweet = domain.NewTextTweet(user, secondText)
    thirdTweet = domain.NewTextTweet(anotherUser, text)
    // publish the 3 tweets

    tweetManager.PublishTweet(tweet)
    tweetManager.PublishTweet(secondTweet)
    tweetManager.PublishTweet(thirdTweet)

    // Operation
    tweets := tweetManager.GetTweetsByUser(user)
    tweetsNotExisting := tweetManager.GetTweetsByUser("dani")

    // Validation
    if len(tweets) != 2 {
      t.Errorf("Expected 2 tweets but was %d", len(tweets))
      return
    }
    if tweetsNotExisting != nil {
      t.Errorf("Expected for user without tweets is nil")
    }
    firstPublishedTweet := tweets[0]
    secondPublishedTweet := tweets[1]
    // check if isValidTweet for firstPublishedTweet and secondPublishedTweet
    if !isValidTweet(t, firstPublishedTweet, user, text) {
      t.Errorf("First tweet obtained is not equal to published tweet")
      return
    }
    if !isValidTweet(t, secondPublishedTweet, user, secondText) {
      t.Errorf("Second tweet obtained is not equal to published tweet")
      return
    }
}
