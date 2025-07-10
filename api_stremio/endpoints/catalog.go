package endpoints

import (
	"FSA/api_stremio/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterCatalogEndpoint(e *echo.Echo) {
	e.GET("/catalog/:type/:id.json", func(c echo.Context) error {
		cType := c.Param("type")
		ID := c.Param("id")

		println(cType, ID)

		stream := entities.Stream{
			URL: "https://www.w3schools.com/html/mov_bbb.mp4",
		}

		metaItem := entities.MetaItem{
			ID:     "tt0133093",
			Type:   "movie",
			Name:   "The Matrix",
			Poster: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ0ErtaYA0YL18ugo2CR2lf8TJKcYpAhvA-5A&s",
			Genres: []string{"Drama", "Sci-Fi"},
			TMDB:   603,
			Video: entities.VideoItem{
				ID:        "fsa_1",
				Title:     "The Matrix",
				Available: true,
				Released:  "2010-12-06T05:00:00.000Z",
				Streams:   []entities.Stream{stream},
			},
		}

		catalogResponse := entities.CatalogResponse{
			Metas: []entities.MetaItem{metaItem},
		}

		return c.JSONPretty(http.StatusOK, catalogResponse, " ")
	})
}
