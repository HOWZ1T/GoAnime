package types

type Episode struct {
	Number       int    `json:"number"`
	Link         string `json:"link"`
	DownloadLink string `json:"download_link"`
}

func NewEpisode(num int, link string, downloadLink string) Episode {
	return Episode{
		Number:       num,
		Link:         link,
		DownloadLink: downloadLink,
	}
}
