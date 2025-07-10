package handlers

import (
	"FSA/api_stremio/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetManifest(c echo.Context) error {
	catalog := CatalogEntry{
		Type: "movie",
		ID:   "1245",
		Name: "Local Movie Catalog",
	}

	resource1 := ManifestResource{
		Name: "stream",
		Types: []string{
			"movie", "series", "anime",
		},
		IDPrefixes: []string{"tt"},
	}

	resource2 := ManifestResource{
		Name: "catalog",
		Types: []string{
			"movie", "series", "anime",
		},
	}

	manifest := &Manifest{
		ID:           "aewww",
		Version:      "0.1.3",
		Name:         "FSA - File Server Addon",
		Description:  "Serves local file through network",
		ContactEmail: "Anderson@AndersonBargas.com",
		Types:        []string{"movie"},
		Resources:    []ManifestResource{resource1, resource2},
		Catalogs:     []CatalogEntry{catalog},
	}

	return c.JSONPretty(http.StatusOK, manifest, " ")
}

func GetCatalog(c echo.Context) error {
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
}

func GetStream(c echo.Context) error {
	cType := c.Param("type")
	ID := c.Param("id")

	println(cType, ID)

	stream := entities.Stream{
		Title: "The Matrix",
		URL:   "http://192.168.1.39/Media/MOVIES/The%20Matrix/The%20Matrix.mkv",
	}

	streamResponse := entities.StreamResponse{
		Streams: []entities.Stream{stream},
	}

	return c.JSONPretty(http.StatusOK, streamResponse, " ")
}
