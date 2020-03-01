package sources

import (
	"GoAnime/anime"
	"log"
	"regexp"
	"strings"
)
import "golang.org/x/net/html"

type Source interface {
	GetSearchUri(phrase string) string
	Parse(body string, source Source) []anime.Anime
}

type goGoAnime struct {
	base_uri   string
	search_uri string
}

var GoGoAnime = goGoAnime{
	base_uri:   "https://gogoanime.io",
	search_uri: "https://ajax.apimovie.xyz/site/loadAjaxSearch?keyword=",
}

func (g goGoAnime) GetSearchUri(phrase string) string {
	return g.search_uri + strings.ReplaceAll(html.EscapeString(phrase), " ", "%20")
}

func (g goGoAnime) Parse(body string, source Source) []anime.Anime {
	re := regexp.MustCompile("category\\\\//?.+?\\\\")
	links := re.FindAllString(body, -1)

	// parse raw links into usable links and then query into anime objects
	for i, e := range links {
		links[i] = g.base_uri + "/category/" + e[10:len(e)-1]
		log.Println(links[i])
		anime.Anime{}.Parse(links[i])
		// TODO request each link and parse into anime.Anime
		break
	}

	return nil
}
