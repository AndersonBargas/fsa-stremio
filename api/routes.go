package api

import (
	"FSA/api/handlers"
	"errors"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	instance *echo.Echo
}

func NewRoutes(wsInstance *echo.Echo) (*Routes, error) {
	if wsInstance == nil {
		return nil, errors.New("The WS Instance should not be null")
	}
	return &Routes{
		instance: wsInstance,
	}, nil
}

func (r *Routes) RegisterEndpoints() {
	r.registerAdminEndpoints()
	r.registerStremioEndpoints()
}

func (r *Routes) registerAdminEndpoints() {
	group := r.instance.Group("/admin")
	group.GET("/", func(c echo.Context) error { return nil })

}

func (r *Routes) registerStremioEndpoints() {
	group := r.instance.Group("/stremio")
	group.GET("/manifest.json", handlers.GetManifest)
	group.GET("/catalog/:type/:id.json", handlers.GetCatalog)
	group.GET("/stream/:type/:id.json", handlers.GetStream)
}
