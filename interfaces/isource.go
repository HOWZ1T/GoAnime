package interfaces

import "GoAnime/types"

type Source interface {
	GetSearchUri(phrase string) string
	Name() string
	Parse(body string) ([]types.Anime, error)
	ParseAnime(uri string) (types.Anime, error)
	BaseUri() string
}
