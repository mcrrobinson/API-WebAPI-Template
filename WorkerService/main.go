package main

import "web-app-template/worker"

func main() {

	// Declare the API port.
	var apiPort int = 8080

	// Make a bool to tell the thread when to stop.
	stop := make(chan bool)

	// Declare the new structure.
	c := worker.NewWorkerStructure()

	// Create the API server, declare endpoints etc.
	c.CreateAPIServer(apiPort)

	// Start listening for API requests.
	c.ListenAPIServer(stop)

	<-stop

	// Stop the API.
	c.APIShutdown()
}
