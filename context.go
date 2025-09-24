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

	handlers HandlerChain
	index    int8
}

func NewRumContext(r *http.Request, w http.ResponseWriter, handlers HandlerChain) *RumContext {
	uuid, err := NewUUID()
	var ctx context.Context
	if err != nil {
		log.Printf("Failed to generate a unique ID for request - %s", err.Error())
		ctx = context.Background()
	} else {
		ctx = context.WithValue(context.Background(), RequestIDKey, uuid)
	}
	return &RumContext{
		ctx: ctx,
		R:   r,
		W:   w,

		handlers: handlers,
		index:    0,
	}
}

func (rc *RumContext) Next() {
	rc.index++
	if int(rc.index) < len(rc.handlers) {
		handler := rc.handlers[rc.index]
		handler(rc)
	}
}

func (rc *RumContext) RequestId() string {
	if rc.ctx == nil {
		return ""
	}
	reqID, ok := rc.ctx.Value(RequestIDKey).(UUID)
	if !ok {
		return ""
	}
	return reqID.String()
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
