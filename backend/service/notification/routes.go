package notification

import (
	"log"
	"net/http"

	"github.com/Jeno7u/studybud/service/auth"
	"github.com/Jeno7u/studybud/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true }, // разрешаем все источники
}

type Handler struct {
	store types.NotificationStore
}

func NewHandler(store types.NotificationStore) *Handler {
	return &Handler{store: store}
}


type Client struct {
    UserID int
    Conn   *websocket.Conn
}

func NewClient(userID int, conn *websocket.Conn) *Client {
	return &Client{
		UserID: userID,
		Conn: conn,
	}
}

func (h *Handler) NotificationRoutes(router gin.IRouter) {
	router.GET("/notifications", auth.AuthMiddleware(), h.wsHandler)
}

var clients = make(map[int]*Client)

func (h *Handler) wsHandler(c *gin.Context) {
    userID := c.GetInt("user_id")

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }

    client := NewClient(userID, conn)
    clients[userID] = client
    log.Printf("User %d connected\n", userID)

    defer func() {
        conn.Close()
        delete(clients, userID)
        log.Printf("User %d disconnected\n", userID)
    }()

    // Цикл чтения, чтобы соединение держалось
    for {
        _, _, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }
	}
}
