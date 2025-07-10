package domain

import "time"

type MediaItem struct {
	ID          string // trocar
	Name        string
	Title       string
	Type        string // trocar
	IMDBRating  float32
	Poster      string // trocar
	ReleaseInfo time.Time
	Runtime     int // in minutes
	Description string
	Cast        []string
	Director    []string
	Writer      []string
	Genres      []string
	FilePath    string
}
