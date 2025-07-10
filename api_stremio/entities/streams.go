package entities

type Stream struct {
	Title         string                 `json:"title,omitempty"`
	Name          string                 `json:"name,omitempty"`
	URL           string                 `json:"url"`
	ExternalURL   string                 `json:"externalUrl,omitempty"`
	InfoHash      string                 `json:"infoHash,omitempty"`
	FileIdx       int                    `json:"fileIdx,omitempty"`
	YTDL          string                 `json:"ytDl,omitempty"`
	Subtitle      string                 `json:"subtitle,omitempty"`
	BehaviorHints map[string]interface{} `json:"behaviorHints,omitempty"`
}

type StreamResponse struct {
	Streams []Stream `json:"streams"`
}
