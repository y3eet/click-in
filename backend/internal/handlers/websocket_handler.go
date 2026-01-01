package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/y3eet/click-in/internal/auth"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	for {
		claims, b := auth.GetClaims(c)
		if !b {
			fmt.Println("unauthorized")
			return
		}
		var data map[string]any
		err := conn.ReadJSON(&data)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := conn.WriteJSON(gin.H{"message": "Message received", "data": data, "user": claims.User}); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
