package domain

import "time"

var idIncremental int = 0

type TextTweet struct{
  Id int
  User string
  Text string
  Date *time.Time
}

func NewTextTweet(user string, text string) *TextTweet {
  date := time.Now()
  idIncremental++
  return &(TextTweet{idIncremental,user, text, &date})
}

func (tweet *TextTweet) PrintableTweet() string {
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

func (tweet *TextTweet) GetText() string{
  return tweet.Text
}

func (tweet *TextTweet) GetUser() string {
  return tweet.User
}

func (tweet *TextTweet) GetId() int {
  return tweet.Id
}

func (tweet *TextTweet) GetDate() *time.Time {
  return tweet.Date
}
