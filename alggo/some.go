package alggo

// TODO thing a better name for this file

import (
	"fmt"
	"time"
)

// TODO do it generic
type Book struct {
	Title       string
	PublishedAt time.Time
}

// filterByOldestDate returns a list of books that share the oldest published date.
func filterByOldestDate(list []Book) []Book {
	fmt.Println("omar alberto barra")
	var oldestList []Book
	for _, file := range list {
		if len(oldestList) == 0 {
			oldestList = append(oldestList, file)
			continue
		}
		compared := file.PublishedAt.Compare(oldestList[0].PublishedAt)
		if compared > 0 {
			continue
		}
		if compared < 0 {
			oldestList = oldestList[:0]
			oldestList = append(oldestList, file)
		} else {
			oldestList = append(oldestList, file)
		}
	}
	return oldestList
}
