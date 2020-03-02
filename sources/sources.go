package sources

import (
	"GoAnime/requests"
	"GoAnime/types"
	"container/list"
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"regexp"
	"strconv"
	"strings"
)

type goGoAnime struct {
	baseUri   string
	searchUri string
	name      string
}

var GoGoAnime = goGoAnime{
	baseUri:   "https://gogoanime.io",
	searchUri: "https://ajax.apimovie.xyz/site/loadAjaxSearch?keyword=",
	name:      "GoGoAnime",
}

func (g goGoAnime) GetSearchUri(phrase string) string {
	return g.searchUri + strings.ReplaceAll(html.EscapeString(phrase), " ", "%20")
}

func (g goGoAnime) Name() string {
	return g.name
}

func (g goGoAnime) BaseUri() string {
	return g.baseUri
}

func (g goGoAnime) ParseAnime(uri string) (types.Anime, error) {
	a := types.NewEmptyAnime()
	log.Debug("engine: [goGoAnime] parsing anime from uri: " + uri)
	contents, err := requests.Get(uri)
	if err != nil {
		return a, err
	}
	log.Debugln(contents)

	// parse page into usable data
	// TODO: bug for some reason the regex in golang doesnt capture the summary for this specific example in main!
	// but in regex101 it does capture
	re := regexp.MustCompile("(<h1>.+?<\\/h1>|<a.+?<\\/a>|<p.+?<\\/p>)")
	parts := re.FindAllString(contents, -1)
	for i, e := range parts {
		if e[1] == 'h' {
			if a.Title == "" {
				a.Title = e[4 : len(e)-5]

				a.Typ = strings.Split(parts[i+2], ">")[1]
				a.Typ = a.Typ[0 : len(a.Typ)-3]

				log.Debug("\n\n\n")
				for i, e := range parts {
					log.Debug(i, ": ", e)
				}

				log.Debug("\n\n\n")
				for _, e := range strings.Split(contents, "\n") {
					log.Debug(e)
				}
				a.PlotSummary = strings.Split(parts[i+3], "</span>")[1]
				a.PlotSummary = a.PlotSummary[0 : len(a.PlotSummary)-4]

				// genres
				run := true
				j := i + 4
				genres := list.New()
				for run == true {
					ej := parts[j]
					if strings.Contains(ej, "genre") == true {
						rawGenre := strings.Split(ej, ">")[1]
						rawGenre = strings.ReplaceAll(rawGenre[0:len(rawGenre)-3], ", ", "")
						genres.PushBack(rawGenre)
						j++
					} else {
						run = false
					}
				}

				// convert genre list to arr
				genreArr := make([]string, genres.Len())
				c := 0
				for e := genres.Front(); e != nil; e = e.Next() {
					genreArr[c] = e.Value.(string)
					c++
				}
				a.Genre = genreArr

				// release year, status, other name and ep
				year, err := strconv.Atoi(strings.Split(strings.Split(parts[j], "</span>")[1], "<")[0])
				if err != nil {
					return a, err
				}
				a.ReleaseYear = year
				a.Status = types.AsStatus(strings.Split(strings.Split(parts[j+1], "</span>")[1], "<")[0])
				a.OtherName = strings.Split(strings.Split(parts[j+2], "</span>")[1], "<")[0]

				// episodes
				rawEp := strings.Split(parts[j+3], ">")[1]
				rawEp = rawEp[0 : len(rawEp)-3]
				rawEpRange := strings.Split(rawEp, "-")
				epStart, err := strconv.Atoi(rawEpRange[0])
				if err != nil {
					return a, err
				}
				epEnd, err := strconv.Atoi(rawEpRange[1])
				if err != nil {
					return a, err
				}

				episodes := list.New()
				rawNameParts := strings.Split(uri, "/")
				rawName := rawNameParts[len(rawNameParts)-1]
				// gogo anime bug on website, epStart may start at 0 but it's actually one!
				if epStart == 0 {
					epStart = 1
				}
				for i := epStart; i <= epEnd; i++ {
					eLink := g.BaseUri() + "/" + rawName + "-episode-" + strconv.Itoa(i)
					episodes.PushBack(types.NewEpisode(i, eLink))
				}

				episodeArr := make([]types.Episode, episodes.Len())
				c = 0
				for e := episodes.Front(); e != nil; e = e.Next() {
					episodeArr[c] = e.Value.(types.Episode)
					c++
				}
				a.Episodes = episodeArr
				return a, nil
			}
		}
	}

	return a, errors.New("malformed parsing")
}

func (g goGoAnime) Parse(body string) ([]types.Anime, error) {
	log.Debug("parsing anime with engine: " + g.Name())
	re := regexp.MustCompile("category\\\\//?.+?\\\\")
	links := re.FindAllString(body, -1)

	animeLst := list.New()
	// parse raw links into usable links and then query into anime objects
	for i, e := range links {
		links[i] = g.baseUri + "/category/" + e[10:len(e)-1]
		log.Debug(links[i])
		anim, err := g.ParseAnime(links[i])
		if err != nil {
			return nil, err
		}
		animeLst.PushBack(anim)
	}

	animeArr := make([]types.Anime, animeLst.Len())
	c := 0
	for e := animeLst.Front(); e != nil; e = e.Next() {
		animeArr[c] = e.Value.(types.Anime)
		c++
	}

	log.Debug("parsed " + strconv.Itoa(animeLst.Len()) + " animes.")
	return animeArr, nil
}
