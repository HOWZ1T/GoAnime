package anime

const UNKNOWN = -1

type Anime struct {
	title       string
	plotSummary string
	otherName   string
	genre       []string

	releaseYear int

	typ      Type
	status   Status
	episodes []Episode
}

func (Anime) NewEmpty() Anime {
	return Anime{
		title:       "",
		plotSummary: "",
		otherName:   "",
		genre:       nil,
		releaseYear: 0,
		typ:         0,
		status:      0,
		episodes:    nil,
	}
}

func (Anime) New(title string, plotSummary string, otherName string, genre []string,
	releaseYear int, typ string, status string, episodes []Episode) Anime {
	return Anime{
		title:       title,
		plotSummary: plotSummary,
		otherName:   otherName,
		genre:       genre,
		releaseYear: releaseYear,
		typ:         AsType(typ),
		status:      AsStatus(status),
		episodes:    episodes,
	}
}
