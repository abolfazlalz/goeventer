package main

import (
	"fmt"
	"github.com/abolfazlalz/goeventer/internal/websocket"
	"github.com/gin-gonic/gin"
	"github.com/goasali/toolkit/config"
	"github.com/goasali/toolkit/http/routers"
	"log"
	"time"
)

type User struct {
}

func (u User) Notify(client *websocket.Client) {
	tick := time.NewTicker(time.Second * 5)
	client.Send([]byte("hello"))

	go func() {
		ch := make(chan *websocket.Chat, 1)
		client.AddListener(ch)
		for {
			msg := <-ch
			fmt.Println(msg.Text)
			websocket.Broadcast(client.ID, fmt.Sprintf("received message: \"%s\" from \"%s\"", msg.Text, msg.SocketID))
		}
	}()

	go func() {
		for {
			select {
			case t := <-tick.C:
				client.Send([]byte(fmt.Sprintf("What the fuck ?: %v", t)))
			}
		}
	}()
}

func main() {
	if err := config.LoadEnvironments(); err != nil {
		log.Fatalf("Error in load enviroments: %v", err)
	}

	router := routes.SetupRouter()

	websocket.GetClientListener().Register(User{})

	router.GET("/ws", func(c *gin.Context) {
		websocket.HandleConnection(c.Writer, c.Request)
	})

	if err := router.Listen(); err != nil {
		log.Fatalf("Error in listenning gin http server: %v", err)
	}
}
