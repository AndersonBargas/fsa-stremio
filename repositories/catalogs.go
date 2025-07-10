package repositories

type Catalog struct {
	ID          string // unique ID
	Type        string // "movie", "series", etc.
	Name        string // name shown in the catalog
	Poster      string // poster URL (optional)
	Background  string // background URL (optional)
	PosterShape string // "poster", "square", "landscape" (optional)
}

type Catalogs interface {
	FindByID(string) *Catalog
	GetAll() []Catalog
}
