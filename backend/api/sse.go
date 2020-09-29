package api

import (
	"fmt"
	"log"
	"net/http"
)

type client struct {
	messages chan []byte
	done     <-chan struct{}
}

func Subscribe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
	}

	log.Println("new client connection")
	client := client{
		messages: make(chan []byte),
		done:     req.Context().Done(),
	}
	defer close(client.messages)

	for {
		select {
		case <-client.done:
			log.Println("closing connection")
			return
		case msg := <-client.messages:
			// TODO: make messages more complete
			fmt.Fprintf(w, "data:%s\n\n", msg)
			f.Flush()
		}
	}
}
