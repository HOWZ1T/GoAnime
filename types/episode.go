package types

type Episode struct {
	Number int    `json:"number"`
	Link   string `json:"link"`
}

func NewEpisode(num int, link string) Episode {
	return Episode{
		Number: num,
		Link:   link,
	}
}
