package http

import (
	"github.com/labstack/echo/v4"
)

type ServerSentEventManager struct {
    channel chan string
}

func (sse *ServerSentEventManager) NewChannel() *ServerSentEventManager {
    sse.channel = make(chan string, 50)

    return sse
}

func (sse *ServerSentEventManager) SetHeadersForContext(c echo.Context) {
    c.Response().Header().Add("Content-Type", "text/event-stream")
    c.Response().Header().Set("Cache-Control", "no-cache")
    c.Response().Header().Set("Connection", "keep-alive")
}

func (sse *ServerSentEventManager) SendMessage(message string) {
    sse.channel <- message
}

func (sse *ServerSentEventManager) Serve(c echo.Context) {
    for {
        select {
        case message := <- sse.channel:
            c.Response().Writer.Write([]byte("data: " + message + "\n\n"))
            c.Response().Flush()
        case <- c.Request().Context().Done():
            // Connection closed
            return
        }
    }
}

