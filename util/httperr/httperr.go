package httperr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// HTTPErr conveys both an error message and an HTTP status code
type HTTPErr interface {
	// StatusCode is the HTTP status code
	StatusCode() int
	// Error produces a string representation of the error
	Error() string
	// WriteError writes an error payload response along with the HTTP status
	WriteError(w http.ResponseWriter)
}

// E is a convenient alias
type E = HTTPErr

// ErrorPayload is what gets written by WriteError
type ErrorPayload struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

type httpErr struct {
	msg     string
	details []string
	status  int
}

// New returns a new HTTPErr
func New(status int, msg string, details []string) HTTPErr {
	return &httpErr{status: status, msg: msg, details: details}
}

// New400 returns a new bad request error
func New400(msg string) HTTPErr {
	return New(http.StatusBadRequest, msg, []string{})
}

// New400Detailed returns a new bad request error with details
func New400Detailed(msg string, details []string) HTTPErr {
	return New(http.StatusBadRequest, msg, details)
}

// New404 returns a new not found error
func New404(msg string) HTTPErr {
	return New(http.StatusNotFound, msg, []string{})
}

// New404Detailed returns a new not found error with details
func New404Detailed(msg string, details []string) HTTPErr {
	return New(http.StatusNotFound, msg, details)
}

// New500 returns a new internal server error
func New500(msg string) HTTPErr {
	return New(http.StatusInternalServerError, msg, []string{})
}

// New500Detailed returns a new internal server error with details
func New500Detailed(msg string, details []string) HTTPErr {
	return New(http.StatusInternalServerError, msg, details)
}

func (e *httpErr) StatusCode() int {
	return e.status
}

func (e *httpErr) Error() string {
	if len(e.details) > 0 {
		return fmt.Sprintf("%s:\n%s", e.msg, strings.Join(e.details, "\n"))
	}

	return e.msg
}

func (e *httpErr) WriteError(w http.ResponseWriter) {
	p := ErrorPayload{Message: e.msg, Details: e.details}

	pJSON, err := json.Marshal(p)
	if err != nil {
		log.Printf("failed to marshal error payload: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(e.status)

	_, err = w.Write(pJSON)

	if err != nil {
		log.Printf("failed to write error payload: %v", err)
	}
}
