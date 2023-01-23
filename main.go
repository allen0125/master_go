package main

func main() {
	client := GetAuthedClient()
	ReplyNotifications(client)
}
