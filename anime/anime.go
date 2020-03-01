package anime

import (
	"GoAnime/requests"
	"log"
	"regexp"
)

const UNKNOWN = -1

type Anime struct {
	title       string
	plotSummary string
	otherName   string
	typ         string
	genre       []string

	releaseYear int

	status   Status
	episodes []Episode
}

func NewEmptyAnime() Anime {
	return Anime{
		title:       "",
		plotSummary: "",
		otherName:   "",
		typ:         "",
		genre:       nil,
		releaseYear: 0,
		status:      0,
		episodes:    nil,
	}
}

func NewAnime(title string, plotSummary string, otherName string, genre []string,
	releaseYear int, typ string, status string, episodes []Episode) Anime {
	return Anime{
		title:       title,
		plotSummary: plotSummary,
		otherName:   otherName,
		genre:       genre,
		releaseYear: releaseYear,
		typ:         typ,
		status:      AsStatus(status),
		episodes:    episodes,
	}
}

func (Anime) Parse(uri string) error {
	contents, err := requests.Get(uri)
	if err != nil {
		return err
	}

	re := regexp.MustCompile("(<h1>.+?</h1>|<a.+?</a>|<p.+?</p>)")
	for i, e := range re.FindAllString(contents, -1) {
		log.Println(i, "| ", e)
	}
	return nil
}
