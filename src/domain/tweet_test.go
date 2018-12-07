package domain_test

import "testing"
import "github.com/danielacarrero/Twitter/src/domain"

func TestTextTweetPrintsUserAndText(t *testing.T) {

    // Initialization
    tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

    // Operation
    text := tweet.PrintableTweet()

    // Validation
    expectedText := "@grupoesfera: This is my tweet"
    if text != expectedText {
        t.Errorf("The expected text is %s but was %s", expectedText, text)
    }

}

func TestTextTweetCreatedWithUserAndText(t *testing.T){
  //Initialization
  user := "grupoesfera"
  text := "This is my tweet"
  //Operation
  tweet := domain.NewTextTweet(user, text)

  //Validation
  if user != tweet.GetUser(){
    t.Errorf("The expected user is %s but was %s", user, tweet.GetUser())
  }
  if text != tweet.GetText(){
    t.Errorf("The expected text is %s but was %s", text, tweet.GetText())
  }
}

func TestImageTweetCreatedWithUserTextAndImaURL(t *testing.T){
  //Initialization
  user := "grupoesfera"
  text := "This is my image"
  url := "http://www.grupoesfera.com.ar/common/img/grupoesfera.png"

  //Operation
  tweet := domain.NewImageTweet(user, text, url)

  //Validation
  if user != tweet.GetUser(){
    t.Errorf("The expected user is %s but was %s", user, tweet.GetUser())
  }
  if text != tweet.GetText(){
    t.Errorf("The expected text is %s but was %s", text, tweet.GetText())
  }
  if url != tweet.GetUrlImg(){
    t.Errorf("The expected text is %s \n but was %s", url, tweet.GetUrlImg())
  }
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

    // Initialization
    tweet := domain.NewImageTweet("grupoesfera", "This is my image",
                "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
    // Operation
    text := tweet.PrintableTweet()
    // Validation
    expectedText := "@grupoesfera: This is my image\thttp://www.grupoesfera.com.ar/common/img/grupoesfera.png"
    if text != expectedText {
      t.Errorf("The expected text is %s but was %s", expectedText, text)
    }

}

func TestQuoteTweetCreatedWithUserTextAndQuoteTweet(t *testing.T){
  //Initialization
  myUser := "grupoesfera"
  myText := "This is my tweet"
  quotedUser := "nick"
  quotedText := "This is nick's tweet"

  //Operation
  quoteTweet := domain.NewTextTweet(quotedUser, quotedText)
  tweet := domain.NewQuoteTweet(myUser, myText, quoteTweet)

  //Validation
  if myUser != tweet.GetUser(){
    t.Errorf("The expected user is %s but was %s", myUser, tweet.GetUser())
  }
  if myText != tweet.GetText(){
    t.Errorf("The expected text is %s but was %s", myText, tweet.GetText())
  }
  if quotedUser != tweet.QuotedTweet.GetUser() {
    t.Errorf("The expected user is %s \n but was %s", quotedUser, tweet.QuotedTweet.GetUser())
  }
  if quotedText != tweet.QuotedTweet.GetText() {
    t.Errorf("The expected text is %s \n but was %s", quotedText, tweet.QuotedTweet.GetText())
  }
}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
    // Initialization
    quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
    tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)
    //Operation
    text := tweet.PrintableTweet()
    // Validation
    expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
    if text != expectedText {
      t.Errorf("The expected text is %s but was %s", expectedText, text)
    }
}
