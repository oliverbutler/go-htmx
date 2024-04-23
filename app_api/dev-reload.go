package app_api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var versionString string

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// versionWsHandler is a WebSocket handler that sends a random string to the client.
// This is used to force the client to reload when the server is restarted.
func versionWsHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	pool, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upgrade WebSocket"})
		return
	}
	_ = pool.WriteMessage(websocket.TextMessage, []byte(versionString))
}

func InitDevReloadWebsocket(r *gin.Engine) {
	versionString = generateRandomString(10)

	r.GET("/ws", versionWsHandler)
}
