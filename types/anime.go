package types

import "strconv"

type Anime struct {
	Title       string
	PlotSummary string
	OtherName   string
	Typ         string
	Genre       []string

	ReleaseYear int

	Status   Status
	Episodes []Episode
}

func (a Anime) StatusAsStr() string { return a.Status.ToString() }
func (a Anime) EpisodesStr() string {
	out := "["
	for _, e := range a.Episodes {
		out += "(" + strconv.Itoa(e.number) + ", " + e.link + "), "
	}
	out = out[0:len(out)-2] + "]"
	return out
}

func NewEmptyAnime() Anime {
	return Anime{
		Title:       "",
		PlotSummary: "",
		OtherName:   "",
		Typ:         "",
		Genre:       nil,
		ReleaseYear: 0,
		Status:      0,
		Episodes:    nil,
	}
}

func NewAnime(title string, plotSummary string, otherName string, genre []string,
	releaseYear int, typ string, status string, episodes []Episode) Anime {
	return Anime{
		Title:       title,
		PlotSummary: plotSummary,
		OtherName:   otherName,
		Genre:       genre,
		ReleaseYear: releaseYear,
		Typ:         typ,
		Status:      AsStatus(status),
		Episodes:    episodes,
	}
}
