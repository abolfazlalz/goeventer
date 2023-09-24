package main

import (
	"fmt"
	"github.com/abolfazlalz/goeventer/internal/interpreter"
	"github.com/abolfazlalz/goeventer/internal/websocket"
	"github.com/gin-gonic/gin"
	"github.com/goasali/toolkit/config"
	"github.com/goasali/toolkit/http/routers"
	"log"
	"os"
)

type User struct {
}

func (u User) Notify(client *websocket.Client) {
	go func() {
		ch := make(chan *websocket.Chat, 1)
		client.AddListener(ch)
		for {
			msg := <-ch
			fmt.Println(msg.Text)
			websocket.Broadcast(client.ID, fmt.Sprintf("received message: \"%s\" from \"%s\"", msg.Text, msg.SocketID))
		}
	}()
}

func main() {
	if err := config.LoadEnvironments(); err != nil {
		log.Fatalf("Error in load enviroments: %v", err)
	}

	router := routes.SetupRouter()

	router.GET("/ws", func(c *gin.Context) {
		websocket.HandleConnection(c.Writer, c.Request)
	})

	go func() {
		interpreter.File(os.Args[1])
	}()

	if err := router.Listen(); err != nil {
		log.Fatalf("Error in listenning gin http server: %v", err)
	}
}
