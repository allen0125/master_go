package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	client := GetAuthedClient()
	notifications, err := client.GetNotifications(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := len(notifications) - 1; i >= 0; i-- {
		fmt.Println(notifications[i])
	}
}
