package main

import (
    "fmt"
    "net/http"
    "net/url"
    "os"

    "github.com/ChimeraCoder/anaconda"
    "github.com/BillyZac/mumble-bot/config"
)

func postToTwitter(w http.ResponseWriter, r *http.Request) {
   config, err := config.NewConfig()
   if err != nil {
     fmt.Printf("ERROR - Configuration error: %s\n", err.Error())   
     os.Exit(1)
   }

   tweet := r.URL.Path[1:]

   anaconda.SetConsumerKey(config.TwitterConsumerKey)
   anaconda.SetConsumerSecret(config.TwitterConsumerSecret)
   api := anaconda.NewTwitterApi(config.TwitterAccessToken, config.TwitterAccessTokenSecret)

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
