package domain

import (
  "time"
  //"fmt"
  //"strings"
)

var idIncremental int = 0

type Tweet struct{
  Id int
  User string
  Text string
  Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
  date := time.Now()
  idIncremental++
  return &(Tweet{idIncremental,user, text, &date})
}

func (tweet *Tweet) PrintableTweet() string {
  /*
  //Otra manera de hacerlo, más eficiente. Se usa cuando se necesita performance.
  var sb strings.Builder
  sb.WriteString("@")
  sb.WriteString(tweet.User)
  sb.WriteString(": ")
  sb.WriteString(tweet.Text)
  return sb.String()
  */
  //return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
  return "@" + tweet.User + ": " + tweet.Text
}
