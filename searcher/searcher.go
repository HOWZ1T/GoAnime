package searcher

import (
	"GoAnime/interfaces"
	"GoAnime/requests"
	"GoAnime/types"
	log "github.com/sirupsen/logrus"
)

func Search(phrase string, source interfaces.Source) ([]types.Anime, error) {
	log.Debug("running search with phrase: " + phrase + " and source: " + source.Name())
	uri := source.GetSearchUri(phrase)
	contents, err := requests.Get(uri)
	if err != nil {
		return nil, err
	}

	animes, err := source.Parse(contents)
	if err != nil {
		return nil, err
	}

	return animes, nil
}
