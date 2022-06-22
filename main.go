package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Tim0401/oapi-codegen-demo/middleware"
	"github.com/Tim0401/oapi-codegen-demo/openapi"
	"github.com/Tim0401/oapi-codegen-demo/router"
	oapiMw "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

const (
	// OpenAPI定義以前にprefixが必要な場合
	apiPathPrefix = "/v0"
	// Auth Header
	HeaderAuthorization      = "Authorization"
	HeaderAuthorizationAdmin = "Authorization-Admin"
)

type Handler struct {
	router.ItemRouterInterface
}

func main() {

	swagger, err := openapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	newPath := openapi3.Paths{}
	for k := range swagger.Paths {
		newPath[apiPathPrefix+k] = swagger.Paths[k]
	}
	swagger.Paths = newPath

	validatorOpts := &oapiMw.Options{}
	// Authorization ヘッダーが必要なルーターに対する存在チェック
	validatorOpts.Options.AuthenticationFunc = func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		h := input.RequestValidationInput.Request.Header[input.SecuritySchemeName]
		if len(h) == 0 {
			return errors.New("HeaderAuthorizationNotFound")
		}
		// ユーザーと管理者のトークンチェック
		switch input.SecuritySchemeName {
		case HeaderAuthorization:
			if err := checkToken(h[0]); err != nil {
				return errors.New("HeaderAuthorization")
			}
		case HeaderAuthorizationAdmin:
			if err := checkTokenAdmin(h[0]); err != nil {
				return errors.New("HeaderAuthorizationAdmin")
			}
		default:
			return errors.New("HeaderAuthorizationNotFound")
		}
		return nil
	}

	// custom error message
	validatorOpts.ErrorHandler = func(c echo.Context, err *echo.HTTPError) error {
		if rerr, ok := err.Internal.(*openapi3filter.RequestError); ok {
			if rerr.Parameter != nil {
				fmt.Print(rerr.Parameter.Name)
				err.Message = rerr.Parameter.Name + "が不正な値です。"
			}
		}
		return err
	}

	// init echo
	e := echo.New()
	g := e.Group(apiPathPrefix)

	g.Use(echoMw.Recover())
	g.Use(echoMw.Logger())

	// バリデーション
	g.Use(oapiMw.OapiRequestValidatorWithOptions(swagger, validatorOpts))

	// カスタムMiddlewareの定義

	// 各ルートに適用
	mwRoot := middleware.NewMiddlewareRoot()
	userAuthRoute := mwRoot.Group(apiPathPrefix)
	userAuthRoute.Use(authUser)
	{
		userAuthRoute.POST("/items")
	}
	adminAuthRoute := mwRoot.Group(apiPathPrefix)
	adminAuthRoute.Use(authAdmin)
	{
		adminAuthRoute.PUT("/items/:id")
		adminAuthRoute.DELETE("/items/:id")
	}

	// カスタムMiddlewareの適用
	g.Use(mwRoot.Exec)

	// 定義した struct を登録
	handler := Handler{
		ItemRouterInterface: &router.ItemRouter{},
	}
	openapi.RegisterHandlers(g, handler)

	// Start server
	port := "9000"
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func checkToken(s string) error {
	fmt.Println("checkToken: do something")
	return nil
}

func checkTokenAdmin(s string) error {
	fmt.Println("checkTokenAdmin: do something")
	return nil
}

func authUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("authUser: do something")
		return next(c)
	}
}

func authAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("authAdmin: do something")
		return next(c)
	}
}
