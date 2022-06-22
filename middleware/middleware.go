package middleware

import "github.com/labstack/echo/v4"

type middlewareRoot struct {
	// [method][path]
	middlewares map[string]map[string][]echo.MiddlewareFunc
	router      *echo.Router
	echo        *echo.Echo
}

func NewMiddlewareRoot() middlewareRoot {
	e := echo.New()
	return middlewareRoot{
		middlewares: map[string]map[string][]echo.MiddlewareFunc{},
		router:      echo.NewRouter(e),
		echo:        e,
	}
}

func (mwr *middlewareRoot) Group(prefix string, m ...echo.MiddlewareFunc) (g *group) {
	g = &group{prefix: prefix, middlewareRoot: mwr}
	g.Use(m...)
	return
}

func (mwr *middlewareRoot) add(method string, path string, m ...echo.MiddlewareFunc) {
	if mwr.middlewares == nil {
		mwr.middlewares = make(map[string]map[string][]echo.MiddlewareFunc)
	}
	if _, ok := mwr.middlewares[method]; !ok {
		mwr.middlewares[method] = make(map[string][]echo.MiddlewareFunc)
	}
	if _, ok := mwr.middlewares[method][path]; !ok {
		mwr.middlewares[method][path] = make([]echo.MiddlewareFunc, 0, len(m))
	}
	mwr.middlewares[method][path] = append(mwr.middlewares[method][path], m...)
	mwr.router.Add(method, path, func(c echo.Context) error { return nil })
}

func (mwr *middlewareRoot) Exec(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := c.Request().Method
		path := c.Request().URL.Path
		if _, ok := mwr.middlewares[method]; !ok {
			return next(c)
		}

		// パスの検索
		mwc := mwr.echo.NewContext(c.Request(), c.Response())
		mwr.router.Find(method, path, mwc)
		routePath := mwc.Path()

		if _, ok := mwr.middlewares[method][routePath]; !ok {
			return next(c)
		}

		middleware := mwr.middlewares[method][routePath]

		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next(c)
	}
}
