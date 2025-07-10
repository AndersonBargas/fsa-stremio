package entities

type VideoItem struct {
	ID        string   `json:"id"`
	Released  string   `json:"released"`
	Title     string   `json:"title"`
	Available bool     `json:"available"`
	Streams   []Stream `json:"streams,omitempty"`
}

type MetaItem struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Poster      string    `json:"poster,omitempty"`
	Background  string    `json:"background,omitempty"`
	Logo        string    `json:"logo,omitempty"`
	Description string    `json:"description,omitempty"`
	Genres      []string  `json:"genres,omitempty"`
	Released    string    `json:"released,omitempty"` // "YYYY-MM-DD"
	Runtime     string    `json:"runtime,omitempty"`  // "120 min"
	Director    string    `json:"director,omitempty"`
	Cast        []string  `json:"cast,omitempty"`
	TMDB        int       `json:"tmdb,omitempty"`
	Country     string    `json:"country,omitempty"`
	Language    string    `json:"language,omitempty"`
	PosterShape string    `json:"posterShape,omitempty"` // square, poster, landscape
	IMDBRating  string    `json:"imdbRating,omitempty"`
	Trailer     string    `json:"trailer,omitempty"`
	Website     string    `json:"website,omitempty"`
	Video       VideoItem `json:"video,omitempty"`
}

type MetaResponse struct {
	Meta MetaItem `json:"meta"`
}
