package main

func main() {
	client := GetAuthedClient()
	ReplyNotifications(client)
	Translate([]string{"Test", "Hello"})
}
