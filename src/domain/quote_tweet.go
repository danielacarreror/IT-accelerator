package domain

import "time"


type QuoteTweet struct{
  textTweet *TextTweet
  QuotedTweet *TextTweet
}

func NewQuoteTweet(user string, text string, quotedTweet *TextTweet) *QuoteTweet {
  return &(QuoteTweet{ NewTextTweet(user, text), quotedTweet})
}

func (tweet *QuoteTweet) PrintableTweet() string {
  return "@" + tweet.textTweet.User + ": " + tweet.textTweet.Text + ` "@` + tweet.QuotedTweet.User + ": " + tweet.QuotedTweet.Text + `"`
}

func (tweet *QuoteTweet) GetText() string{
  return tweet.textTweet.Text
}

func (tweet *QuoteTweet) GetUser() string {
  return tweet.textTweet.User
}

func (tweet *QuoteTweet) GetId() int {
  return tweet.textTweet.Id
}

func (tweet *QuoteTweet) GetDate() *time.Time {
  return tweet.textTweet.Date
}
