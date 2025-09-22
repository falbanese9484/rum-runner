package rum

import "net/http"

type (
	HandlerFunc  func(*RumContext) // Function to handle request
	HandlerChain []HandlerFunc     // Chain of handlers (middleware + final handler)
)

type RouteInfo struct {
	// HTTP method, path, and the handler name
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

type RoutesInfo []RouteInfo // Slice of RouteInfo

type RouterGroup struct {
	handlers HandlerChain
	basePath string
	engine   *Engine
}

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
	routeInfo := RouteInfo{
		Method:      method,
		Path:        fullPath,
		Handler:     "",
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

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routes := e.routes[r.Method]
	for _, route := range routes {
		if route.Path == r.URL.Path {
			ctx := NewRumContext(r, w)
			route.HandlerFunc(ctx)
			return
		}
	}

	http.NotFound(w, r)
}
