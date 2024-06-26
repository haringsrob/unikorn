// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

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

	"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",
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
