package domain

type URL struct {
	Key      string `json:"key"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
