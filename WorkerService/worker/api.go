package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CreateAPIServer creates the mux instance and sets the endpoints.
func (work *Worker) CreateAPIServer(port int) {
	router := mux.NewRouter()

	// Setup all of the endpoints
	router.HandleFunc("/startPolling", work.PollerExample)
	router.HandleFunc("/startStream", work.StartStreamRequest)
	router.HandleFunc("/stopStream", work.StopStreamRequest)
	http.Handle("/", router)

	work.apiServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handlers.CORS()(router),
	}
}

func (work *Worker) greet() {
	broker := NewServer()
	type QueueItem struct {
		Name      string `json:"name"`
		Position  string `json:"position"`
		Office    string `json:"office"`
		Age       int    `json:"age"`
		StartDate string `json:"startdate"`
		Salary    int    `json:"salary"`
	}
	queueItem := QueueItem{Name: "Harold", Position: "Officer", Office: "Americano", Age: 51, StartDate: "October 21st", Salary: 50320}
	jsonMessage, jsonError := json.Marshal(queueItem)
	if jsonError != nil {
		work.logger.Error("Unable to marshal the JSON.")
	}

	go func() {
		for {
			// Simulate your application sending data.
			if work.activeStream == true {
				time.Sleep(time.Millisecond * 1000)
				log.Println("Sending event")
				broker.Notifier <- []byte(jsonMessage)
			}
		}
	}()
	work.logger.Crit("HTTP server error: ", http.ListenAndServe("localhost:3000", broker))
}

// ListenAPIServer starts the API server on a new thread as ListenAndServe blocks.
func (work *Worker) ListenAPIServer(stop chan bool) {
	work.logger.Info(
		"Starting API Server",
		"address", work.apiServer.Addr,
	)

	go func() {
		err := work.apiServer.ListenAndServe()
		if err != nil {
			work.logger.Crit(
				"Error with API Server",
				"error", err,
			)
			stop <- true
		}
	}()
	work.greet()
}

// APIShutdown stops the API.
func (work *Worker) APIShutdown() error {
	work.logger.Info("API Server Shutting down")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	return work.apiServer.Shutdown(ctx)
}

// KNOWN ISSUES;
// 	WONT UNREGISTER CLIENT UNTIL EVENT STREAM IS ACTIVE.
// 	SETUP TWO ROUTINES ONE FOR SENDING DATA AND ONE FOR
// 	CHECKING TO SEE IF A USER IS REGISTERED.

// COULD ALSO DO WITH A BUFFER LATER ON.
