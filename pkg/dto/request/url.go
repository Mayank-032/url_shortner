package request

type ShortURL struct {
	LongURL string `json:"long_url,required"`
}

type RedirectURL struct {
	Key         string `json:"key,required"`
	IsKeySigned bool   `json:"is_key_signed,required"`
}
