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

func ExtractContent(content string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("span").Remove()
	text := doc.Find("p").Text()
	return text
}

func ReplyNotification(client *mastodon.Client, notification *mastodon.Notification) {
	if notification.Type == "mention" {
		if notification.Status.InReplyToID != nil {
			statusID := mastodon.ID(fmt.Sprintf("%v", notification.Status.InReplyToID))
			status, err := client.GetStatus(context.Background(), statusID)
			if err != nil {
				log.Fatal(err)
			}
			text := ExtractContent(status.Content)
			result := Translate([]string{text}, "auto2zh")[0]
			result = "@" + notification.Status.Account.Acct + " " + result
			toot := &mastodon.Toot{InReplyToID: notification.Status.ID, Status: result}
			newStatus, err := client.PostStatus(context.Background(), toot)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(newStatus)
		}
	}
}

func ReplyNotifications(client *mastodon.Client) {
	notifications, err := client.GetNotifications(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := len(notifications) - 1; i >= 0; i-- {
		ReplyNotification(client, notifications[i])
	}

	err = client.ClearNotifications(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
