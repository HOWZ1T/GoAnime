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
	base_uri:   "https://gogoanime.video",
	search_uri: "https://ajax.apimovie.xyz/site/loadAjaxSearch?keyword=",
}

func (g goGoAnime) GetSearchUri(phrase string) string {
	return g.search_uri + strings.ReplaceAll(html.EscapeString(phrase), " ", "%20")
}

func (g goGoAnime) Parse(body string, source Source) []anime.Anime {
	log.Println(body)
	re := regexp.MustCompile("href=.+?\" ")
	links := re.FindAllString(body, -1)

	// parse raw links into usable links
	for i, e := range links {
		// TODO: Faster way to parse this string ?
		links[i] = strings.ReplaceAll(e[7:], "\\/", "\\")
		links[i] = g.base_uri + "/" + strings.ReplaceAll(links[i][:len(links[i])-3], "\\", "/")
		log.Println(links[i])
	}

	// TODO request each link and parse into anime.Anime
	return nil
}
