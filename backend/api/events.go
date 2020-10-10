package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const heartbeatInterval = 1 * time.Second

type event struct {
	category string
	content  []byte
}

func (e *event) String() string {
	if string(e.content) == ":" {
		return ":\n"
	}
	return fmt.Sprintf("event: %s\ndata: %s\n\n", e.category, e.content)
}

type client struct {
	clientId string
	events   chan event
}

type broker struct {
	clients map[string]client
}

func NewBroker() broker {
	return broker{clients: make(map[string]client)}
}

func (c client) listen(ctx context.Context, w http.ResponseWriter, f http.Flusher) {
	done := ctx.Done()
	heartbeat := time.NewTicker(heartbeatInterval)
	defer heartbeat.Stop()

	for {
		select {
		case <-done:
			log.Printf("close connection for %s", c.clientId)
			return
		case event := <-c.events:
			log.Printf("deliver event for %s: %s", c.clientId, event)
			num, err := fmt.Fprint(w, event.String())
			if err != nil {
				log.Printf("error sending message for %s: %v", c.clientId, err)
				continue
			}
			log.Printf("flush %d bytes for %s", num, c.clientId)
			f.Flush()
			log.Printf("reset heartbeat timer for %s", c.clientId)
			heartbeat.Reset(heartbeatInterval)
			log.Printf("delivered event for %s", c.clientId)
		case <-heartbeat.C:
			log.Printf("tick for %s", c.clientId)
			go func() {
				log.Printf("send heartbeat for %s", c.clientId)
				c.events <- event{content: []byte(":")}
				log.Printf("sent heartbeat for %s", c.clientId)
			}()
			log.Printf("finished tick for %s", c.clientId)
		}
	}
}

func (c client) disconnect() {
	close(c.events)
}

func subscribe(b broker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*") // TODO: revert

		f, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		}

		vars := mux.Vars(req)
		clientId, ok := vars["clientId"]
		if !ok {
			http.Error(w, "missing required ID", http.StatusBadRequest)
		}

		// TODO: validate clientId

		log.Printf("new client connection %s", clientId)
		client := client{clientId: clientId, events: make(chan event)}
		defer client.disconnect()
		b.clients[clientId] = client

		client.listen(req.Context(), w, f)
		log.Printf("closed connection for %s", clientId)
	}
}
