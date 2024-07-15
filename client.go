package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	types "github.com/timmo001/letmeknow-types-go"
)

type Client struct {
	LMKHost       string
	LMKPort       int
	LMKClientType types.ClientType
	LMKUserID     string

	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin
		return true
	},
}

func (c *Client) Connect() error {
	// Connect to the server
	conn, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf("ws://%s:%d/websocket", c.LMKHost, c.LMKPort),
		http.Header{
			"User-Agent": []string{fmt.Sprintf("LMKClientGo/%s", c.LMKClientType)},
		},
	)
	if err != nil {
		return err
	}
	c.conn = conn

	return nil
}

func (c *Client) Disconnect() error {
	// Disconnect from the server
	err := c.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) IsConnected() bool {
	// Check if the client is connected to the server
	return c.conn != nil
}

func (c *Client) RegisterClient() error {
	r := types.RequestRegister{
		Type:   "register",
		UserID: c.LMKUserID,
	}
	log.Printf("Registering client: %s", c.LMKUserID)

	// Register the client with the server
	return c.conn.WriteJSON(r)
}

func (c *Client) SendNotification(notification types.Notification) error {
	log.Printf("Sending notification:\n%s", notification.Display())
	// Send a notification to the server
	return c.conn.WriteJSON(notification)
}
