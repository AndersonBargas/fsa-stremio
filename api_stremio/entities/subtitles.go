package entities

type Subtitle struct {
	ID              string `json:"id"`
	Lang            string `json:"lang"`
	URL             string `json:"url"`
	Name            string `json:"name,omitempty"`
	HearingImpaired bool   `json:"hearingImpaired,omitempty"`
}

type SubtitlesResponse struct {
	Subtitles []Subtitle `json:"subtitles"`
}
