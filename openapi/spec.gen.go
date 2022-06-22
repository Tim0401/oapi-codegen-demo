// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUT2sUMRz9KuWnx7g7rQVxjoqWIqgHLyI9pDO/blM2f5pkpMMSaLfgyX/0YO2tiiiI",
	"Jw9SqPTLzG6t30KSme6/WXcreErYzXt5v/fepAOJ5EoKFNZA3AGTbCKnYcsscr+maBLNlGVSQAxF92PR",
	"/VTsvyj2j4GA0lKhtgxN7egs5MXXb+fvXwMBmyuEGIzVTLTAEWCpR9Z+FpTjbMr+21fT+JRmyRRk7+zD",
	"+fFPIMCZYDzjEEcDMBMWW6jBOQIatzOmMYX4mZdWCbmkXSNgmW0HkHdrQCHXtzCxQGDnhrFStVlr0wZP",
	"U4ihnd9az5Z2blO+bSk45+9hYkOGwSs6qVBQxRo5bwOB56hNqXuxEfmpqr8hhpuNqBF5QdRuhgyaXknY",
	"tdDOtqx3snvx+UvRPei/edc/Oyz2Dou9s2LvCMIFmnrIqle8gnY1sHpDjJLClHkvRZFfEiksinAZVarN",
	"koBsbpmyBmWn/G68LQOh8xQCGZ69rnEDYrjWHPa2WZU2TO7dqUKgWtO8nmJgWpuMKqQwLuTRgwA2GedU",
	"5xDDU5nphZV7TxZQpEoy4QO2tGVGaH3jpJnje9E96J3s9k5//Do6/f3y+999fyzNiPHbGRp7R6b5P3k+",
	"3zBXWjSW62J9grsaqcU0eDI5tSNV75odlrqR8k3t0f+v0RUHrTfhykUgsBwt1015KO3CfZmJdG5VfDOo",
	"phwtau9dB5gn8N/t5bMSl0/MUKPVGZKRwSdeN7fmnPsTAAD//7A6QG/CBQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}