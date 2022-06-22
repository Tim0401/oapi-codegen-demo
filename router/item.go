package router

import (
	"github.com/Tim0401/oapi-codegen-demo/openapi"
	"github.com/labstack/echo/v4"
)

type ItemRouterInterface interface {
	// (GET /items)
	GetItems(ctx echo.Context, params openapi.GetItemsParams) error

	// (POST /items)
	PostItems(ctx echo.Context) error

	// (DELETE /items/{id})
	DeleteItem(ctx echo.Context, id string) error
	// Your GET endpoint
	// (GET /items/{id})
	GetItem(ctx echo.Context, id string) error

	// (PUT /items/{id})
	PutItem(ctx echo.Context, id string) error
}

type ItemRouter struct{}

func (r *ItemRouter) GetItems(ctx echo.Context, params openapi.GetItemsParams) error {
	return nil
}
func (r *ItemRouter) PostItems(ctx echo.Context) error {
	return nil
}
func (r *ItemRouter) DeleteItem(ctx echo.Context, id string) error {
	return nil
}
func (r *ItemRouter) GetItem(ctx echo.Context, id string) error {
	return nil
}
func (r *ItemRouter) PutItem(ctx echo.Context, id string) error {
	return nil
}
