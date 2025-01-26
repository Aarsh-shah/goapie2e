package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"time"
)

// Upgrader to upgrade HTTP requests to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (customize for security)
	},
}

// Client represents a WebSocket connection
type Client struct {
	conn *websocket.Conn
	send chan []byte // Channel to send messages to the client
}
type Dependencies struct {
	grpc TardisFeatureServiceClient
}

// WebSocket handler
func handleConnections(deps Dependencies, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading:", err)
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte),
	}

	// Start a goroutine to handle outgoing messages
	go handleOutgoingMessages(client)

	// Read messages from the client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received message: %s", message)
		fmt.Println("printing here ")
		resp, _ := deps.grpc.GetFeatures(context.Background(), &TardisGetFeatureRequest{FeatureName: string(message)})
		fmt.Println("printing resp")
		fmt.Println(resp.Features)
		// You can broadcast or respond directly here
		client.send <- []byte("Server response: " + resp.Features)
	}
}

func handleOutgoingMessages(client *Client) {
	for msg := range client.send {
		err := client.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write error:", err)
			client.conn.Close()
			break
		}
	}
}

func main() {
	grpc_middleware.ChainStreamClient()
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println(err)

	}
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()))
	service := NewAarshServer()
	server.RegisterService(&TardisFeatureService_ServiceDesc, service)
	reflection.Register(server)
	fmt.Println("Serving on port 8085")

	go func() { server.Serve(listener) }()
	time.Sleep(5 * time.Second)
	client, _ := grpc.Dial("localhost:8085", grpc.WithInsecure())
	conn := NewTardisFeatureServiceClient(client)
	deps := &Dependencies{grpc: conn}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleConnections(*deps, w, r)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
