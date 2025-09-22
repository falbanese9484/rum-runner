package rum

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type contextKey string

const RequestIDKey contextKey = "rum-request-id"

type RumContext struct {
	ctx context.Context
	R   *http.Request
	W   http.ResponseWriter
}

func NewRumContext(r *http.Request, w http.ResponseWriter) *RumContext {
	uuid, err := NewUUID()
	if err != nil {
		log.Printf("Failed to generate a unique ID for request - %s", err.Error())
		// Kind of need to think about what to do here...
	}
	ctx := context.WithValue(context.Background(), RequestIDKey, uuid)
	return &RumContext{
		ctx: ctx,
		R:   r,
		W:   w,
	}
}

func (rc *RumContext) JSON(code int, payload any) {
	if rc.W == nil {
		log.Printf("ResponseWriter is nil, cannot write response")
		return
	}
	w := rc.W
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Printf("Failed to encode JSON response - %s", err.Error())
	}
}

func (rc *RumContext) String(code int, payload string) {
	if rc.W == nil {
		log.Printf("ResponseWriter is nil, cannot write response")
		return
	}
	w := rc.W
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	_, err := w.Write([]byte(payload))
	if err != nil {
		log.Printf("Failed to write string response - %s", err.Error())
	}
}

func (rc *RumContext) HTML(code int, payload string) {
	if rc.W == nil {
		log.Printf("ResponseWriter is nil, cannot write response")
		return
	}
	w := rc.W
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	_, err := w.Write([]byte(payload))
	if err != nil {
		log.Printf("Failed to write HTML response - %s", err.Error())
	}
}

func (rc *RumContext) Status(code int) {
	if rc.W == nil {
		log.Printf("ResponseWriter is nil, cannot write response")
		return
	}
	w := rc.W
	w.WriteHeader(code)
}
