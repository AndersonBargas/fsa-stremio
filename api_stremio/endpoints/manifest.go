package endpoints

import (
	"FSA/api_stremio/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterManifestEndpoint(e *echo.Echo) {
	e.GET("/manifest.json", func(c echo.Context) error {

		catalog := entities.CatalogEntry{
			Type: "movie",
			ID:   "1245",
			Name: "Local Movie Catalog",
		}

		resource1 := entities.ManifestResource{
			Name: "stream",
			Types: []string{
				"movie", "series", "anime",
			},
			IDPrefixes: []string{"tt"},
		}

		resource2 := entities.ManifestResource{
			Name: "catalog",
			Types: []string{
				"movie", "series", "anime",
			},
		}

		manifest := &entities.Manifest{
			ID:           "aewww",
			Version:      "0.1.3",
			Name:         "FSA - File Server Addon",
			Description:  "Serves local file through network",
			ContactEmail: "Anderson@AndersonBargas.com",
			Types:        []string{"movie"},
			Resources:    []entities.ManifestResource{resource1, resource2},
			Catalogs:     []entities.CatalogEntry{catalog},
		}
		return c.JSONPretty(http.StatusOK, manifest, " ")
	})
}
