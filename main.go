package main

import (
	"flag"
	"log"

	types "github.com/timmo001/letmeknow-types-go"
)

func main() {
	var host string = "localhost"
	var port int = 8080

	// Parse the command line arguments
	flag.StringVar(&host, "host", "localhost", "The host to connect to")
	flag.IntVar(&port, "port", 8080, "The port to connect to")
	flag.Parse()

	// Create a new client
	c := Client{
		LMKHost:       host,
		LMKPort:       port,
		LMKClientType: types.ClientTypeHeadless,
		LMKUserID:     GenerateUserID(types.ClientTypeHeadless, nil),
	}

	log.Println("Connecting to server:", c.LMKHost, c.LMKPort)
	err := c.Connect()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to server")

	err = c.RegisterClient()
	if err != nil {
		panic(err)
	}
	log.Println("Client registered")

	title := "Hello, World!"
	subtitle := "This is a test notification"
	content := "This is a test notification from the Go client"

	notification := types.Notification{
		Type:     "notification",
		Title:    &title,
		Subtitle: &subtitle,
		Content:  &content,
		Image: &types.Image{
			URL: "https://picsum.photos/400",
		},
	}

	err = c.SendNotification(notification)
	if err != nil {
		panic(err)
	}
	log.Println("Notification sent successfully")

	err = c.Disconnect()
	if err != nil {
		panic(err)
	}

}
