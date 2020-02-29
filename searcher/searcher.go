package searcher

import (
	"GoAnime/anime"
	"GoAnime/requests"
	"GoAnime/sources"
)

func Search(phrase string, source sources.Source) (anime.Anime, error) {
	uri := source.GetSearchUri(phrase)
	contents, err := requests.Get(uri)
	if err != nil {
		return anime.Anime{}, err
	}

	source.Parse(contents, source)
	// TODO parse DOM to construct anime object

	return anime.Anime{}, nil
}
