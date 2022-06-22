package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	methods = [...]string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		echo.PROPFIND,
		http.MethodPut,
		http.MethodTrace,
		echo.REPORT,
	}
)

type group struct {
	prefix         string
	middleware     []echo.MiddlewareFunc
	middlewareRoot *middlewareRoot
}

func (g *group) Use(middleware ...echo.MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

func (g *group) Group(prefix string, middleware ...echo.MiddlewareFunc) (sg *group) {
	m := make([]echo.MiddlewareFunc, 0, len(g.middleware)+len(middleware))
	m = append(m, g.middleware...)
	m = append(m, middleware...)
	sg = g.middlewareRoot.Group(g.prefix+prefix, m...)
	return
}

func (g *group) Add(method, path string, middleware ...echo.MiddlewareFunc) {
	m := make([]echo.MiddlewareFunc, 0, len(g.middleware)+len(middleware))
	m = append(m, g.middleware...)
	m = append(m, middleware...)
	g.middlewareRoot.add(method, g.prefix+path, m...)
}

// CONNECT implements `Echo#CONNECT()` for sub-routes within the Group.
func (g *group) CONNECT(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodConnect, path, m...)
}

// DELETE implements `Echo#DELETE()` for sub-routes within the Group.
func (g *group) DELETE(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodDelete, path, m...)
}

// GET implements `Echo#GET()` for sub-routes within the Group.
func (g *group) GET(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodGet, path, m...)
}

// HEAD implements `Echo#HEAD()` for sub-routes within the Group.
func (g *group) HEAD(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodHead, path, m...)
}

// OPTIONS implements `Echo#OPTIONS()` for sub-routes within the Group.
func (g *group) OPTIONS(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodOptions, path, m...)
}

// PATCH implements `Echo#PATCH()` for sub-routes within the Group.
func (g *group) PATCH(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodPatch, path, m...)
}

// POST implements `Echo#POST()` for sub-routes within the Group.
func (g *group) POST(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodPost, path, m...)
}

// PUT implements `Echo#PUT()` for sub-routes within the Group.
func (g *group) PUT(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodPut, path, m...)
}

// TRACE implements `Echo#TRACE()` for sub-routes within the Group.
func (g *group) TRACE(path string, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodTrace, path, m...)
}

// Any implements `Echo#Any()` for sub-routes within the Group.
func (g *group) Any(path string, middleware ...echo.MiddlewareFunc) {
	for _, m := range methods {
		g.Add(m, path, middleware...)
	}
}

// Match implements `Echo#Match()` for sub-routes within the Group.
func (g *group) Match(methods []string, path string, middleware ...echo.MiddlewareFunc) {
	for _, m := range methods {
		g.Add(m, path, middleware...)
	}
}
