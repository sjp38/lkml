package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type rssItem struct {
	title  string
	author string
	link   string
}

func isElemOf(element string, tagname string) bool {
	starttag := fmt.Sprintf("<%s>", tagname)
	endtag := fmt.Sprintf("</%s>", tagname)

	return strings.HasPrefix(element, starttag) &&
		strings.HasSuffix(element, endtag)
}

func unEscape(txt string) string {
	txt = strings.Replace(txt, "&lt;", "<", -1)
	return strings.Replace(txt, "&gt;", ">", -1)
}

func contentOf(element string, tagname string) string {
	starttag := fmt.Sprintf("<%s>", tagname)
	endtag := fmt.Sprintf("</%s>", tagname)

	sidx := len(starttag)
	eidx := len(element) - len(endtag)
	return unEscape(element[sidx:eidx])
}

func parseRSS(rssText string) []rssItem {
	// Assumes that each element is splitted by newline
	rsslines := strings.Split(rssText, "\n")

	var items []rssItem
	var item rssItem
	for _, line := range rsslines {
		txt := strings.TrimSpace(line)
		if txt == "<item>" {
			item = rssItem{}
			continue
		}
		if isElemOf(txt, "title") {
			item.title = contentOf(txt, "title")
			continue
		}
		if isElemOf(txt, "author") {
			item.author = contentOf(txt, "author")
			continue
		}
		if isElemOf(txt, "link") {
			item.link = contentOf(txt, "link")
			continue
		}
		if txt == "</item>" {
			items = append(items, item)
		}
	}

	return items
}

func main() {
	resp, err := http.Get("https://lkml.org/rss.php")
	if err != nil {
		panic(fmt.Sprintf("failed to get rss: %s", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("failed to read body: %s", err))
	}

	items := parseRSS(string(body))

	// 0th index is rss channel title. So, skip it.
	for i := len(items) - 1; i > 0; i-- {
		fmt.Printf("%s\n\t%s\n\t%s\n", items[i].title,
			items[i].author, items[i].link)
	}
}
