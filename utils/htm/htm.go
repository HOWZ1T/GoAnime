package htm

import (
	"container/list"
	"golang.org/x/net/html"
	"strings"
)

type HtmlEntry struct {
	Toke html.Token
	Text string
}

// returns a list of HtmlEntry
func GetTags(dom *html.Tokenizer, tags []string) *list.List {
	tknLst := list.New()
	var prevToke html.Token
	for {
		tokeType := dom.Next()
		switch tokeType {
		case html.ErrorToken:
			// eof
			return tknLst

		case html.StartTagToken:
			toke := dom.Token()
			prevToke = toke

		case html.TextToken:
			push := false
			txt := strings.TrimSpace(string(dom.Text()))
			if len(txt) <= 0 {
				continue
			} else if tags == nil {
				push = true
			} else if hasTag(prevToke, tags) == true {
				push = true
			}

			if push == true {
				tknLst.PushBack(HtmlEntry{
					Toke: prevToke,
					Text: txt,
				})
			}
		}
	}
}

func hasTag(token html.Token, tags []string) bool {
	for _, t := range tags {
		if strings.ToLower(t) == token.Data {
			return true
		}
	}

	return false
}
