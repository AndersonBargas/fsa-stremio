package handlers

type Manifest struct {
	ID           string             `json:"id"`
	Version      string             `json:"version"`
	Name         string             `json:"name"`
	Description  string             `json:"description,omitempty"`
	Resources    []ManifestResource `json:"resources"` // or []string or []ManifestResource for granularity
	Types        []string           `json:"types"`     // movie, series, channel, tv
	Catalogs     []CatalogEntry     `json:"catalogs,omitempty"`
	Logo         string             `json:"logo,omitempty"`
	Background   string             `json:"background,omitempty"`
	ContactEmail string             `json:"contactEmail,omitempty"`
}

type CatalogEntry struct {
	Type  string      `json:"type"`
	ID    string      `json:"id"`
	Name  string      `json:"name,omitempty"`
	Extra []ExtraProp `json:"extra,omitempty"`
}

type ExtraProp struct {
	Name         string   `json:"name"`
	Options      []string `json:"options,omitempty"`
	IsRequired   bool     `json:"isRequired,omitempty"`
	OptionsLimit int      `json:"optionsLimit,omitempty"`
	Min          int      `json:"min,omitempty"`
	Max          int      `json:"max,omitempty"`
}

type ManifestResource struct {
	Name       string   `json:"name"`
	Types      []string `json:"types,omitempty"`
	ID         string   `json:"id,omitempty"`
	IDPrefixes []string `json:"idPrefixes,omitempty"`
}
