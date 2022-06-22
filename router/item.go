package router

import (
	"fmt"
	"net/http"

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
	if params.Top != nil {
		fmt.Printf("%v", *params.Top)
	}
	return ctx.JSON(http.StatusOK, openapi.GetItemsRes{
		Items: []openapi.Item{
			{},
		},
	})
}
func (r *ItemRouter) PostItems(ctx echo.Context) error {
	body := openapi.PostItemsJSONRequestBody{}
	if err := ctx.Bind(&body); err != nil {
		fmt.Printf("%v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusCreated)
}
func (r *ItemRouter) DeleteItem(ctx echo.Context, id string) error {
	fmt.Printf("%v", id)
	return ctx.NoContent(http.StatusNoContent)
}
func (r *ItemRouter) GetItem(ctx echo.Context, id string) error {
	fmt.Printf("%v", id)
	return ctx.JSON(http.StatusOK, openapi.GetItemRes{
		Item: openapi.Item{},
	})
}
func (r *ItemRouter) PutItem(ctx echo.Context, id string) error {
	fmt.Printf("%v", id)
	body := openapi.PutItemJSONRequestBody{}
	if err := ctx.Bind(&body); err != nil {
		fmt.Printf("%v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	fmt.Printf("%v", body)
	return ctx.NoContent(http.StatusNoContent)
}
