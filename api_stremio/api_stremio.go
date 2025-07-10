package apistremio

import (
	"FSA/api_stremio/endpoints"

	"github.com/labstack/echo/v4"
)

func RegisterEverything(e *echo.Echo) {
	endpoints.RegisterManifestEndpoint(e)
	endpoints.RegisterCatalogEndpoint(e)
	endpoints.RegisterStreamEndpoint(e)
}
