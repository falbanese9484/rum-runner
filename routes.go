package rum

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	Segments    []string
	Params      []string
	HandlerFunc HandlerFunc
}

type (
	RoutesInfo  []RouteInfo
	RouterGroup struct {
		handlers HandlerChain
		basePath string
		engine   *Engine
	}
)

func (rg *RouterGroup) GET(path string, handler HandlerFunc) {
	fullPath := rg.basePath + path
	rg.engine.addRoute("GET", fullPath, handler)
}

func (rg *RouterGroup) POST(path string, handler HandlerFunc) {
	fullPath := rg.basePath + path
	rg.engine.addRoute("POST", fullPath, handler)
}

func (rg *RouterGroup) PUT(path string, handler HandlerFunc) {
	fullPath := rg.basePath + path
	rg.engine.addRoute("PUT", fullPath, handler)
}

func (rg *RouterGroup) DELETE(path string, handler HandlerFunc) {
	fullPath := rg.basePath + path
	rg.engine.addRoute("DELETE", fullPath, handler)
}

func (rg *RouterGroup) PATCH(path string, handler HandlerFunc) {
	fullPath := rg.basePath + path
	rg.engine.addRoute("PATCH", fullPath, handler)
}
