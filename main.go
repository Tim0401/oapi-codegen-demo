package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Tim0401/oapi-codegen-demo/middleware"
	"github.com/Tim0401/oapi-codegen-demo/openapi"
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
	}
	adminAuthRoute := mwRoot.Group(apiPathPrefix)
	adminAuthRoute.Use(authAdmin)
	{
	}

	// カスタムMiddlewareの適用
	g.Use(mwRoot.Exec)

	// 定義した struct を登録
	// openapi.RegisterHandlers(g, handler)

	// 起動
	port := "9000"
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func checkToken(s string) error {
	fmt.Print("checkToken: do something")
	return nil
}

func checkTokenAdmin(s string) error {
	fmt.Print("checkTokenAdmin: do something")
	return nil
}

func authUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Print("authUser: do something")
		return next(c)
	}
}

func authAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Print("authAdmin: do something")
		return next(c)
	}
}
