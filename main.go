package main

import (
	"GoAnime/sources"
	"encoding/json"
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
	/*animes, err := searcher.Search("danganronpa", sources.GoGoAnime)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	for _, e := range animes {
		b, err := json.Marshal(e)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Info(string(b))
	}*/

	anime, err := sources.GoGoAnime.ParseAnime("https://gogoanime.io/category/-danganronpa-the-animation")
	if err != nil {
		log.Fatal(err)
		os.Exit(-2)
	}

	b, err := json.Marshal(anime)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	log.Info(string(b))
}
