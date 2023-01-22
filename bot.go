package main

import (
	"context"
	"github.com/mattn/go-mastodon"
	"log"
	"os"
)

func GetAuthedClient() *mastodon.Client {
	client := mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("TOOTBOT_SERVER"),
		ClientID:     os.Getenv("TOOTBOT_CLIENT_ID"),
		ClientSecret: os.Getenv("TOOTBOT_CLIENT_SECRET"),
	})

	err := client.Authenticate(context.Background(), os.Getenv("TOOTBOT_EMAIL"), os.Getenv("TOOTBOT_PASSWD"))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
