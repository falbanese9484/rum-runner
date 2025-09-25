package rum

import (
	"net/http"
	"strings"
)

type (
	HandlerFunc  func(*RumContext) // Function to handle request
	HandlerChain []HandlerFunc     // Chain of handlers (middleware + final handler)
)

type Engine struct {
	RouterGroup
	routes map[string][]RouteInfo // Map of method to slice of RouteInfo
}

func New() *Engine {
	engine := &Engine{
		routes: make(map[string][]RouteInfo),
	}
	engine.engine = engine
	return engine
}

func (e *Engine) NewGroup(basePath string) *RouterGroup {
	return &RouterGroup{
		basePath: basePath,
		engine:   e,
	}
}

func (e *Engine) addRoute(method, path string, handler HandlerFunc) {
	fullPath := e.basePath + path
	if _, exists := e.routes[method]; !exists {
		e.routes[method] = []RouteInfo{}
	}
	parts := strings.Split(fullPath, "/")
	params := make([]string, 0)
	if len(parts) > 1 {
		for _, p := range parts {
			if len(p) > 0 && p[0] == ':' {
				params = append(params, p[1:])
			}
		}
	}
	routeInfo := RouteInfo{
		Method:      method,
		Path:        fullPath,
		Handler:     "",
		Segments:    parts,
		Params:      params,
		HandlerFunc: handler,
	}
	e.routes[method] = append(e.routes[method], routeInfo)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRoute("POST", path, handler)
}

func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.addRoute("PUT", path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.addRoute("DELETE", path, handler)
}

func (e *Engine) PATCH(path string, handler HandlerFunc) {
	e.addRoute("PATCH", path, handler)
}

func (e *Engine) Use(handler HandlerFunc) {
	e.handlers = append(e.handlers, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routes := e.routes[r.Method]
	for _, route := range routes {
		params := make(map[string]string, 0)
		parts := strings.Split(r.URL.Path, "/")
		match := true
		if len(parts) == len(route.Segments) {
			for i, p := range parts {
				if p != route.Segments[i] && route.Segments[i][0] == ':' {
					params[route.Segments[i][1:]] = p
				} else if p != route.Segments[i] {
					match = false
					break
				}
			}
			if match {
				chain := append(e.handlers, route.HandlerFunc)
				ctx := NewRumContext(r, w, chain, params, r.URL.Path)
				exec := NewHandlerChainExecutor(ctx, chain)
				exec.Begin()
				exec.Complete()

				return
			}
		}
	}

	http.NotFound(w, r)
}
