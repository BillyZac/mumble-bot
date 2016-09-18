package config

import (
	"fmt"
	"os"
)

type Config struct {
	TwitterConsumerKey       string
	TwitterConsumerSecret    string
	TwitterAccessToken       string
	TwitterAccessTokenSecret string
}

func NewConfig() (*Config, error) {
	config := &Config{}
	config.TwitterConsumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	config.TwitterConsumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	config.TwitterAccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	config.TwitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	err := config.validate()
	return config, err
}

func (c *Config) validate() error {
	if len(c.TwitterConsumerKey) < 1 {
		return fmt.Errorf("Twitter consumer key config not valid")
	}
	if len(c.TwitterConsumerSecret) < 1 {
		return fmt.Errorf("Twitter consumer secret config not valid")
	}
	if len(c.TwitterAccessToken) < 1 {
		return fmt.Errorf("Twitter access token config not valid")
	}
	if len(c.TwitterAccessTokenSecret) < 1 {
		return fmt.Errorf("Twitter access token secret config not valid")
	}

	return nil
}
