package service_test


import (
  "testing"
  "github.com/danielacarrero/Twitter/src/service"
)

//Se debe empezar con Test en la función para indicar que es test
// Debe recibir como parametro *testing.T, importando del package testing el tipo T, este tipo es el definido para test
// t es un puntero al tipo T
func TestPublishedTweetIsSaved(t *testing.T){
  var tweet string = "This is my first tweet"

  service.PublishTweet(tweet)

  if service.Tweet != tweet {
    t.Error("Expected tweet is", tweet)
  }
}
