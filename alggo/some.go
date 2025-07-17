package alggo

// TODO thing a better name for this file

import (
	"time"
)

// TODO do it generic
type Book struct {
	Title       string
	PublishedAt time.Time
}

// filterByOldestDate returns a list of books that share the oldest published date.
func filterByOldestDate(list []Book) []Book {
	var oldestList []Book
	for _, item := range list {
		// the first item is always added
		if len(oldestList) == 0 {
			oldestList = append(oldestList, item)
			continue
		}

		compared := item.PublishedAt.Compare(oldestList[0].PublishedAt)
		if compared > 0 {
			continue
		}
		if compared < 0 {
			// if the current item is older than the first one, we clear the list
			// and add the current item
			oldestList = oldestList[:0]
			oldestList = append(oldestList, item)
		} else {
			oldestList = append(oldestList, item)
		}
	}
	return oldestList
}
