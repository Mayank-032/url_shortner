package response

type ShortURL struct {
	ShortURL    string `json:"short_url"`
	Key         string `json:"key"`
	IsKeySigned bool   `json:"is_key_signed"`
}
