// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

const (
	AuthorizationScopes       = "Authorization.Scopes"
	Authorization_AdminScopes = "Authorization_Admin.Scopes"
)

// アイテム
type Item struct {
	// アイテム説明
	Description *string `json:"description,omitempty"`
	Id          string  `json:"id"`

	// アイテム名
	Name string `json:"name"`

	// 価格
	Price int `json:"price"`
}

// GetItemRes defines model for GetItemRes.
type GetItemRes struct {
	// アイテム
	Item Item `json:"item"`
}

// GetItemsRes defines model for GetItemsRes.
type GetItemsRes struct {
	Items []Item `json:"items"`
}

// GetItemsParams defines parameters for GetItems.
type GetItemsParams struct {
	// 取得数
	Top *int `form:"$top,omitempty" json:"$top,omitempty"`
}

// PostItemsJSONBody defines parameters for PostItems.
type PostItemsJSONBody = Item

// PutItemJSONBody defines parameters for PutItem.
type PutItemJSONBody = Item

// PostItemsJSONRequestBody defines body for PostItems for application/json ContentType.
type PostItemsJSONRequestBody = PostItemsJSONBody

// PutItemJSONRequestBody defines body for PutItem for application/json ContentType.
type PutItemJSONRequestBody = PutItemJSONBody
