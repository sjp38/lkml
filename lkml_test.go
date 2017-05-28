package main

import (
	"reflect"
	"testing"
)

func TestItemsAfter(t *testing.T) {
	items := []rssItem{
		rssItem{
			title:  "jkl",
			author: "foo4",
			link:   "4",
		},
		rssItem{
			title:  "ghi",
			author: "foo3",
			link:   "3",
		},
		rssItem{
			title:  "def",
			author: "foo2",
			link:   "2",
		},
		rssItem{
			title:  "abc",
			author: "foo",
			link:   "1",
		},
	}

	after := itemsAfter(items, items[1])
	if !reflect.DeepEqual(after, items[:1]) {
		t.Errorf("expected %v but %v\n", items[2:], after)
	}

	after = itemsAfter(items, items[0])
	if !reflect.DeepEqual(after, []rssItem{}) {
		t.Errorf("expected %v but %v\n", []rssItem{}, after)
	}
}
