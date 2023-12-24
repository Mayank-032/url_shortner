package request

type ShortURL struct {
	LongURL string `json:"long_url"`
}

type RedirectURL struct {
	Key      string `json:"key"`
	ShortURL string `json:"short_url"`
}
