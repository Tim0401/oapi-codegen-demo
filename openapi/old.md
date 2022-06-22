`oapi-codegen@v1.11`以前の生成コマンド  
`--old-config-style`を使うことで`v1.11`以降でも生成可能  
https://github.com/deepmap/oapi-codegen/releases/tag/v1.11.0  

```go
//go:generate oapi-codegen --old-config-style -generate "server" -package openapi -o "server.gen.go" openapi.yml
//go:generate oapi-codegen --old-config-style -generate "types" -package openapi -o "types.gen.go" .openapi.yml
//go:generate oapi-codegen --old-config-style -generate "spec" -package openapi -o "spec.gen.go" openapi.yml
```