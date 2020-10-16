package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// PollerExample is the third routing template.
func (work *Worker) PollerExample(w http.ResponseWriter, r *http.Request) {
	type QueueItem struct {
		Task   string `json:"task"`
		Status string `json:"status"`
	}
	queueItem := QueueItem{Task: "Dishes", Status: "In progress"}
	queueItem2 := QueueItem{Task: "Washing", Status: "Todo"}
	queue := []QueueItem{queueItem, queueItem2}
	jsonMessage, jsonError := json.Marshal(queue)
	if jsonError != nil {
		work.logger.Error("Unable to marshal the JSON.")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonMessage)
}

const patience time.Duration = time.Second * 1

// Broker contains the clients and the message in a byte array.
type Broker struct {
	Notifier       chan []byte
	newClients     chan chan []byte
	closingClients chan chan []byte
	clients        map[chan []byte]bool
}

// NewServer declares the broker and runs the function listen in
// a go routine.
func NewServer() (broker *Broker) {
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}
	go broker.listen()
	return
}

func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	messageChan := make(chan []byte)
	broker.newClients <- messageChan
	defer func() {
		broker.closingClients <- messageChan
	}()
	notify := rw.(http.CloseNotifier).CloseNotify()

	for {
		select {
		case <-notify:
			return
		default:
			fmt.Fprintf(rw, "data: %s\n\n", <-messageChan)
			flusher.Flush()
		}
	}

}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:
			for clientMessageChan, _ := range broker.clients {
				select {
				case clientMessageChan <- event:
				case <-time.After(patience):
					log.Print("Skipping client.")
				}
			}
		}
	}
}

// StartStreamRequest is the third routing template.
func (work *Worker) StartStreamRequest(w http.ResponseWriter, r *http.Request) {
	if work.activeStream == true {
		w.WriteHeader(201)
	} else {
		w.WriteHeader(http.StatusOK)
		work.activeStream = true
	}
	work.logger.Info("An API request was made.", "API Path", r.URL.Path)
}

// StopStreamRequest is the second routing template.
func (work *Worker) StopStreamRequest(w http.ResponseWriter, r *http.Request) {
	if work.activeStream == false {
		w.WriteHeader(202)
	} else {
		w.WriteHeader(http.StatusOK)
		work.activeStream = false
	}
	work.logger.Info("An API request was made.", "API Path", r.URL.Path)
}
