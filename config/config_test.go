package config

import (
  "testing"
)

func TestValidateSuccess(t *testing.T) {
  config := &Config{
    TwitterConsumerKey: "consumerkey",
    TwitterConsumerSecret: "consumersecret",
    TwitterAccessToken: "accesstoken",
    TwitterAccessTokenSecret: "accesstokensecret",
  }  
  
  err := config.validate()
  if err != nil {
    t.Fail()   
  }
}
