package main

import (
    "fmt"
    "net/http"
    "net/url"
    "os"

    "github.com/ChimeraCoder/anaconda"
)

func postToTwitter(w http.ResponseWriter, r *http.Request) {
   tweet := r.URL.Path[1:]

   anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
   anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
   api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

   fmt.Fprintf(w, "Going to post")

   result, err := api.PostTweet(tweet, url.Values{})
   if err != nil {
     fmt.Printf("Error: %s\n", err.Error())
   } else {
     fmt.Printf("Result: %v+\n", result)
   }
}

func main() {
    http.HandleFunc("/", postToTwitter)
    http.ListenAndServe(":8080", nil)
}
