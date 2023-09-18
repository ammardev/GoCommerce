package http

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type ServerSentEventManager struct {
    Channel <- chan *redis.Message
}

func (sse *ServerSentEventManager) SetHeadersForContext(c echo.Context) {
    c.Response().Header().Add("Content-Type", "text/event-stream")
    c.Response().Header().Set("Cache-Control", "no-cache")
    c.Response().Header().Set("Connection", "keep-alive")
}

func (sse *ServerSentEventManager) Serve(c echo.Context) {
    for {
        select {
        case message := <- sse.Channel:
            c.Response().Writer.Write([]byte("data: " + message.Payload + "\n\n"))
            c.Response().Flush()
        case <- c.Request().Context().Done():
            // Connection closed
            return
        }
    }
}

