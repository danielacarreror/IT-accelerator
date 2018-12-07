package domain

import "time"


type ImageTweet struct{
  textTweet *TextTweet
  urlImg string
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
  return &(ImageTweet{ NewTextTweet(user, text), url})
}

func (tweet *ImageTweet) PrintableTweet() string {
  return "@" + tweet.textTweet.User + ": " + tweet.textTweet.Text + "\n" + tweet.UrlImg
}

func (tweet *ImageTweet) GetText() string{
  return tweet.textTweet.Text
}

func (tweet *ImageTweet) GetUser() string {
  return tweet.textTweet.User
}

func (tweet *ImageTweet) GetId() int {
  return tweet.textTweet.Id
}

func (tweet *ImageTweet) GetDate() *time.Time {
  return tweet.textTweet.Date
}

func (tweet *ImageTweet) GetUrlImg() string{
  return tweet.UrlImg
}

func (tweet * ImageTweet) SetUrlImg(url string){
  tweet.urlImg = url
}
