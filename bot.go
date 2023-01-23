package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/go-mastodon"
	"log"
	"os"
	"strings"
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

func ExtractContent(content string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("span").Remove()
	text := doc.Find("p").Text()
	fmt.Println(text)
}

func ReplyNotification(notification *mastodon.Notification) {
	if notification.Type == "mention" {
		content := notification.Status.Content
		ExtractContent(content)
	}
}

func ReplyNotifications(client *mastodon.Client) {
	notifications, err := client.GetNotifications(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := len(notifications) - 1; i >= 0; i-- {
		ReplyNotification(notifications[i])
	}
}
