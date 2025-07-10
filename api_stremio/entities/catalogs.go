package entities

// type CatalogItem struct {
// 	ID          string `json:"id"`
// 	Type        string `json:"type"`
// 	Name        string `json:"name"`
// 	Poster      string `json:"poster,omitempty"`
// 	Background  string `json:"background,omitempty"`
// 	PosterShape string `json:"posterShape,omitempty"`
// }

type CatalogResponse struct {
	Metas []MetaItem `json:"metas"`
}
