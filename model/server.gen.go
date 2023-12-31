// Package model provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package model

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /articles)
	GetArticles(ctx echo.Context, params GetArticlesParams) error

	// (POST /articles)
	PostArticle(ctx echo.Context) error

	// (DELETE /articles/{id})
	DeleteArticle(ctx echo.Context, id int64) error

	// (GET /articles/{id})
	GetArticle(ctx echo.Context, id int64) error

	// (GET /feeds)
	GetFeeds(ctx echo.Context) error

	// (POST /feeds)
	PostFeed(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArticles converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticles(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetArticlesParams
	// ------------- Optional query parameter "feedIDs" -------------

	err = runtime.BindQueryParameter("form", true, false, "feedIDs", ctx.QueryParams(), &params.FeedIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter feedIDs: %s", err))
	}

	// ------------- Optional query parameter "categories" -------------

	err = runtime.BindQueryParameter("form", true, false, "categories", ctx.QueryParams(), &params.Categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter categories: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArticles(ctx, params)
	return err
}

// PostArticle converts echo context to params.
func (w *ServerInterfaceWrapper) PostArticle(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostArticle(ctx)
	return err
}

// DeleteArticle converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteArticle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteArticle(ctx, id)
	return err
}

// GetArticle converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArticle(ctx, id)
	return err
}

// GetFeeds converts echo context to params.
func (w *ServerInterfaceWrapper) GetFeeds(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFeeds(ctx)
	return err
}

// PostFeed converts echo context to params.
func (w *ServerInterfaceWrapper) PostFeed(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostFeed(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/articles", wrapper.GetArticles)
	router.POST(baseURL+"/articles", wrapper.PostArticle)
	router.DELETE(baseURL+"/articles/:id", wrapper.DeleteArticle)
	router.GET(baseURL+"/articles/:id", wrapper.GetArticle)
	router.GET(baseURL+"/feeds", wrapper.GetFeeds)
	router.POST(baseURL+"/feeds", wrapper.PostFeed)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RX0W/bthP+Vw7sD/hhgGCnabEHvaXzshkYlmLIW2oMtHSy2EqkSp6SGIb+9+FIy5Ii",
	"JXa2dMuwF1uUjnff3X36jtqJxJSV0ajJiXgnXJJjKf3lhSWVFMiXsiiuMhHf7MT/LGYiFm/m3bb5fs/8",
	"V7xr9zTR06Z7u9+D42bVRGJ4K96JypoKLSn0aFTKv5mxpSQRC6Xp+/ciErStMCxxg1Y0TSQsfq2VxVTE",
	"N7xrdTAy68+YkGgi8aO1xo5jJCbFh1HenU9EiUSJzsmNt94/dGSV3owQeJ+d/RSaS8T0WVX2G46VmI36",
	"9e2tv21xezSIdwLvZVmFy0QSboz1MW/6q6i/WEUi8wUJf5Go6nWhXI7pQhJHGq4jUdtCxP63iR42tBdx",
	"JxRh6SY6dshBWiu3vObIy9NqMsI35V5RMf3EYz9GITZqnRywPVL5lkvHyr5921V9e87eQhn3+2aJKV+u",
	"nC+Q/zhdtlI6M7w/RZdYVZEyml+HN3Ch4eLjEiiXBAWSg6R2ZEq0DrhzViYEd4py+EXpLw6kTkEGzrpP",
	"+jpXDlxu6iKFNYKENRKhhSSv9RcwGaQmqUvUJDkgXH/4GeLvPulDj+JWyxyDEJG4ResCtLezs9kZJ24q",
	"1LJSIhbv/K1IVJJyX855i4QXG6Rxgr8h1VY7kEXRgy28V+tBMXnFT0gtEO/fyhIJrfMSM/TIpHJABjJV",
	"cKrrrYgE3leFV0SyNXKxRSy+1mj5mZYl7vctF+w+yM6AFye8Ow+Z4mjrC8hbvcINYXYE/BNYB4IzAfco",
	"jYfgVkxSVxntQqfOz87CDNGE2jdNVlWhEt+P+WfHCeym4p4wJ8domP3D2rREgBaV8CaZrAt6FrCn8ITB",
	"ORG91nhfYUKYAnY2lXET9P3BoiR0IEHjXUtgUBooR3BkLM5gUQeIbGaRmW7u/DQYUvyjcS3HRRANdPTB",
	"pNsXy7h/qBmnvX/EbJRpyn+HHERfxJiVzV8kzEk8eZQXr4cWTdRJ3Hyn0iYwpMAwPYcOwn3milN6U+CB",
	"LmvpMAUTSLNcgKsZ/QRFFt5DR5IndXC5YH2XXVf3uHoCk8nCtQrDmt0JjEpHPe8LzQly2CqMU354T2jM",
	"+3GJWrQBavoKXvunh5Ye9/DQ3eUiApX5ntYOLaQGHWhDkMtbBJkk6Fz7lslDSx8be8/udoaU5K+n2f9d",
	"ffDnkZPOP95y+vBz6Z38HWN6/0F2dEZ7RP+2AR1AnzKeAW+RzUy9ydl6C5U1a7kutvvjtP4/wRpn4E/Y",
	"igNcXy2uZpNj/TJ8/X2jmR46Nt2hf3yaPwaOuf6K3lPWMLS3rbaGD8icqIrn88F35Kr5IwAA//+un0ch",
	"ZxIAAA==",
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

