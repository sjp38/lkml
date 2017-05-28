package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	keyword = flag.String("keyword", "", "Keyword to be exist in title")
	delay   = flag.Int("delay", 0, "Delay between updates in seconds")
	count   = flag.Int("count", 1, "Updates count")
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
		} else if isElemOf(txt, "title") {
			item.title = contentOf(txt, "title")
		} else if isElemOf(txt, "author") {
			item.author = contentOf(txt, "author")
		} else if isElemOf(txt, "link") {
			item.link = contentOf(txt, "link")
		} else if txt == "</item>" {
			items = append(items, item)
		}
	}

	return items
}

func fetchRSS() string {
	resp, err := http.Get("https://lkml.org/rss.php")
	if err != nil {
		panic(fmt.Sprintf("failed to get rss: %s", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("failed to read body: %s", err))
	}

	return string(body)
}

func itemsAfter(items []rssItem, last rssItem) []rssItem {
	for i := len(items) - 1; i > 0; i-- {
		it := items[i]
		if it.link == last.link {
			return items[i+1:]
		}
	}
	return items
}

var lastItem rssItem

func printLKML() {
	items := parseRSS(fetchRSS())
	items = itemsAfter(items, lastItem)

	if len(items) > 0 {
		lastItem = items[len(items)-1]
	} else {
		fmt.Printf(" ...")
		return
	}

	for i := len(items) - 1; i >= 0; i-- {
		it := items[i]
		if !strings.Contains(it.title, *keyword) {
			continue
		}
		fmt.Printf("%s\n\t%s\n\t%s\n\n", it.title, it.author, it.link)
	}
}

func main() {
	flag.Parse()

	i := 0
	for {
		printLKML()

		i++
		if *count == -1 {
			goto cont
		}
		if i >= *count {
			break
		}
	cont:
		time.Sleep(time.Duration(*delay) * time.Second)
	}
}
