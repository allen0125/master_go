package main

import "time"

func main() {
	client := GetAuthedClient()
	for true {
		ReplyNotifications(client)
		time.Sleep(time.Second * 60)
	}
}
