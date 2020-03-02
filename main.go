package main

import (
	"GoAnime/searcher"
	"GoAnime/sources"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// setting logger level
	if os.Getenv("GOLAND_DEBUG") == "1" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// setting logger formatting
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006/01/02 | 15:04:05"
	log.SetFormatter(customFormatter)
}
func main() {
	log.Info("running...")
	animes, err := searcher.Search("Danganronpa 3: The End of Kibougamine Gakuen - Mirai-hen", sources.GoGoAnime)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	for _, e := range animes {
		log.Info(e.Title)
	}
}
