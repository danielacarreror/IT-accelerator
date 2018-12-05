package service

import "fmt"

var Tweet string


func PublishTweet(tweet string){
  Tweet = tweet
  fmt.Println(tweet);
}

func main() {
  
}
