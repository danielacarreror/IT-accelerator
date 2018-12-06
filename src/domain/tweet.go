package domain

import "time"

type Tweet interface{
  PrintableTweet() string
  GetText() string
  GetUser() string
  GetId() int
  GetDate() *time.Time
}
