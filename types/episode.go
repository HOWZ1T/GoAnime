package types

type Episode struct {
	number int
	link   string
}

func NewEpisode(num int, link string) Episode {
	return Episode{
		number: num,
		link:   link,
	}
}
