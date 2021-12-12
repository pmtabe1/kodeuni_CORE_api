package websocket_client

import (
	"fmt"
	"log"
	"os"

	"github.com/rgamba/evtwebsocket"
)

type IWebSocketClient interface {
}

type WebsocketClient struct {
	Connection WebsocketClientConnection
}

func NewWebsocketClient(endpointUrl string, payload string) (websocketClient *WebsocketClient) {

	if len(os.Getenv("TRA_API_URL")) == 0 || len(endpointUrl) == 0 {
		traBaseUrl := "wss://localhost:9001"
		endpointUrl = traBaseUrl + "/tra/api/v1/ws/instant/"
	}

	//Declare closures

	// type onConnected func(w *evtwebsocket.Conn)
	// type onError func(w *evtwebsocket.Conn)
	// type onMessage func(w *evtwebsocket.Conn)
	// type pingMsg func([]byte)

	//type onConnected websocketClient.Connection.Connection.OnConnected

	c := evtwebsocket.Conn{
		// Fires when the connection is established

		OnConnected: func(w *evtwebsocket.Conn) {
			log.Println("Connected!")

		},
		// Fires when a new message arrives from the server
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			fmt.Printf("New message: %s\n", msg)
		},
		// Fires when an error occurs and connection is closed
		OnError: func(err error) {
			message := fmt.Sprintf("Error: %s\n", err.Error())
			log.Println(message)
			os.Exit(1)
		},
		// Ping interval in secs (optional)
		PingIntervalSecs: 5,
		// Ping message to send (optional)
		PingMsg: []byte("PING"),

		// websocketClient.Connection.Connection.OnConnected =onConnected

	}
	// Connect
	err := c.Dial(endpointUrl, payload)

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err.Error())
	}
	return &WebsocketClient{
		Connection: WebsocketClientConnection{
			Connection: &c,
		},
	}
}

func NewWebsocketRequest(endpointUrl string, payload string) {
	if len(os.Getenv("TRA_API_URL")) == 0 || len(endpointUrl) == 0 {
		traBaseUrl := "wss://localhost:9001"
		endpointUrl = traBaseUrl + "/tra/api/v1/ws/instant/"
	}

	if payload == "" {
		payload = "test payload ..."
	}
	c := evtwebsocket.Conn{
		// Fires when the connection is established
		OnConnected: func(w *evtwebsocket.Conn) {
			fmt.Println("Connected!")
		},
		// Fires when a new message arrives from the server
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			fmt.Printf("New message: %s\n", msg)
		},
		// Fires when an error occurs and connection is closed
		OnError: func(err error) {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		},
		// Ping interval in secs (optional)
		PingIntervalSecs: 5,
		// Ping message to send (optional)
		PingMsg: []byte("PING"),
	}
	// Connect
	err := c.Dial(endpointUrl, "")

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err.Error())
	}

	msg := evtwebsocket.Msg{
		Body: []byte(payload),
		Callback: func(resp []byte, w *evtwebsocket.Conn) {
			// This function executes when the server responds
			fmt.Printf("Got response: %s\n", resp)
		},
	}
	c.Send(msg)
}
