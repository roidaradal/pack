package my

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/zeroibot/pack/db"
	"github.com/zeroibot/pack/dict"
	"github.com/zeroibot/pack/qb"
)

var errNoDBConn = errors.New("no db connection")

// Task contains the Action and Target of the task
type Task struct {
	Action string
	Target string
}

// Request is an application request that holds a DB connection, a transaction,
// result checker, transaction queries, request start time, and logs
type Request struct {
	Task
	Name    string
	Params  dict.Object
	DB      db.Conn
	Tx      db.Tx
	Checker qb.ResultChecker
	Status  int
	Now     time.Time
	// Private fields
	start   time.Time
	txSteps []qb.Query
	// Logs
	mu   sync.RWMutex
	logs []string
}

// NewRequest creates a new Request
func (i *Instance) NewRequest(name string, args ...any) (*Request, error) {
	if len(args) > 0 {
		name = fmt.Sprintf(name, args...)
	}
	rq := newRequest(name)
	if i.dbConn == nil {
		rq.Status = Err500
		return nil, errNoDBConn
	}
	rq.DB = i.dbConn
	return rq, nil
}

// NewRequestAt creates a new Request at custom db
func (i *Instance) NewRequestAt(key, name string, args ...any) (*Request, error) {
	if len(args) > 0 {
		name = fmt.Sprintf(name, args...)
	}
	rq := newRequest(name)
	conn, ok := i.dbConnMap[key]
	if !ok || conn == nil {
		rq.Status = Err500
		return nil, errNoDBConn
	}
	rq.DB = conn
	return rq, nil
}

// Create a new Request object
func newRequest(name string) *Request {
	return new(Request{
		Name:   name,
		Params: make(dict.Object),
		Status: OK200,
		start:  time.Now(),
		logs:   make([]string, 0),
	})
}
