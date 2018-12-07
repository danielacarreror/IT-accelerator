package service

import (
  "os"
  "github.com/danielacarrero/Twitter/src/domain"
)

type TweetWriter interface {
  WriteTweet(domain.Tweet)
}

type FileTweetWriter struct {
  file *os.File
}

type MemoryTweetWriter struct {
  lastTweet domain.Tweet
}

func NewFileTweetWriter() *FileTweetWriter{
  file, _ := os.OpenFile(
    "tweets.txt",
    os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
    0666,
  )

  writer := new(FileTweetWriter)
  writer.file = file

  return writer
}

func (tweetWriter * FileTweetWriter) WriteTweet (tweet domain.Tweet) {
  go func() {
    if tweetWriter.file != nil{
      byteSlice := []byte(tweet.PrintableTweet() + "\n")
      tweetWriter.file.Write(byteSlice)
    }
  }()
}

func NewMemoryTweetWriter() *MemoryTweetWriter{
  return &MemoryTweetWriter{}
}

func (tweetWriter * MemoryTweetWriter) WriteTweet(tweet domain.Tweet) {
  tweetWriter.lastTweet = tweet
}

func (tweetWriter * MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
  return tweetWriter.lastTweet
}
