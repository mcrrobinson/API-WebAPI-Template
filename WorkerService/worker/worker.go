package worker

import (
	"net/http"
	"os"

	log "github.com/inconshreveable/log15"
)

// Worker is the struct that organises the Clinic Service.
type Worker struct {
	logger       log.Logger
	apiServer    *http.Server
	address      string
	port         int
	activeStream bool
}

// NewWorkerStructure Defines an instance of the worker structure.
func NewWorkerStructure() *Worker {
	work := &Worker{}
	work.activeStream = false
	work.logger = log.New("module", "Worker")
	work.logger.SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat()))
	return work
}
