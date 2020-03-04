package sources

import (
	"GoAnime/requests"
	"GoAnime/types"
	"GoAnime/utils/htm"
	"container/list"
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

func (g goGoAnime) parseTokens(tags *list.List, a *types.Anime, uri string) (*list.List, *list.List) {
	var prev htm.HtmlEntry
	passGenre := false
	genreLst := list.New()
	episodeLst := list.New()
	for e := tags.Front(); e != nil; e = e.Next() {
		toke := e.Value.(htm.HtmlEntry).Toke
		txt := e.Value.(htm.HtmlEntry).Text

		switch toke.Data {
		case "h1":
			a.Title = txt

		case "span":
			if prev.Toke.Data == "span" {
				switch {
				case strings.Contains(prev.Text, "Type"):
					a.Typ = txt

				case strings.Contains(prev.Text, "Plot Summary"):
					a.PlotSummary = html.UnescapeString(txt)

				case strings.Contains(prev.Text, "Released"):
					yr, err := strconv.Atoi(txt)
					if err == nil {
						a.ReleaseYear = yr
					} else {
						log.Error("couldn't pass release year: ", err)
					}

				case strings.Contains(prev.Text, "Status"):
					a.Status = types.AsStatus(txt)

				case strings.Contains(prev.Text, "Other name"):
					a.OtherName = txt
				}
			}

			if strings.Contains(txt, "Genre") {
				passGenre = true
			}

			if passGenre == true && strings.Contains(txt, "Released") {
				passGenre = false
			}

		case "a":
			if passGenre == true {
				genreLst.PushBack(strings.ReplaceAll(txt, ", ", ""))
			}

			if strings.Contains(txt, "-") {
				parts := strings.Split(txt, "-")
				start, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Error("couldn't parse episode start: ", err)
					start = -1
				}

				end, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Error("couldn't parse episode end: ", err)
					end = -1
				}

				if end != -1 && start != -1 {
					rawNameParts := strings.Split(uri, "/")
					rawName := rawNameParts[len(rawNameParts)-1]
					if start == 0 {
						start = 1
					}
					for i := start; i <= end; i++ {
						eLink := g.baseUri + "/" + rawName + "-episode-" + strconv.Itoa(i)
						episodeLst.PushBack(types.NewEpisode(i, eLink))
					}
				}
			}
		}

		prev = e.Value.(htm.HtmlEntry)
	}

	return genreLst, episodeLst
}

func (goGoAnime) genreEpisodesToArr(genreLst *list.List, episodeLst *list.List) ([]string, []types.Episode) {
	if genreLst.Len() > 0 && episodeLst.Len() > 0 {
		genres := make([]string, genreLst.Len())
		c := 0
		for e := genreLst.Front(); e != nil; e = e.Next() {
			genres[c] = e.Value.(string)
			c++
		}

		episodes := make([]types.Episode, episodeLst.Len())
		c = 0
		for e := episodeLst.Front(); e != nil; e = e.Next() {
			episodes[c] = e.Value.(types.Episode)
			c++
		}

		return genres, episodes
	}

	return nil, nil
}

func (g goGoAnime) ParseAnime(uri string) (types.Anime, error) {
	a := types.NewEmptyAnime()
	log.Debug("engine: [goGoAnime] parsing anime from uri: " + uri)
	contents, err := requests.GetRaw(uri)
	if err != nil {
		return a, err
	}

	dom := html.NewTokenizer(strings.NewReader(string(contents)))
	tags := htm.GetTags(dom, []string{"h1", "a", "span"})
	tags = g.trimUnnessescaryTags(tags)

	genreLst, episodeLst := g.parseTokens(tags, &a, uri)
	genre, episodes := g.genreEpisodesToArr(genreLst, episodeLst)

	a.Genre = genre
	a.Episodes = episodes

	return a, nil
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

// returns a new list of htm.HtmlEntry with the necessary tags for parsing the anime
func (g goGoAnime) trimUnnessescaryTags(tags *list.List) *list.List {
	tgs := list.New()
	// usable from h1
	hit := false
	for e := tags.Front(); e != nil; e = e.Next() {
		if hit == false && e.Value.(htm.HtmlEntry).Toke.Data == "h1" {
			hit = true
		} else if hit == true && e.Value.(htm.HtmlEntry).Toke.Data == "span" && strings.Contains(e.Value.(htm.HtmlEntry).Text, "Show") {
			hit = false
			return tgs
		}

		if hit == true {
			tgs.PushBack(e.Value.(htm.HtmlEntry))
		}
	}

	return nil
}
