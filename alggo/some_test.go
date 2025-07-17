package alggo

import (
	"testing"
	"time"
)

func TestFilterByOldestDate(t *testing.T) {
	// Helper function to parse time
	parseTime := func(date string) time.Time {
		t, _ := time.Parse("2006-01-02", date)
		return t
	}

	tests := []struct {
		name     string
		input    []Book
		expected []Book
	}{
		{
			name: "Single book",
			input: []Book{
				{Title: "Book A", PublishedAt: parseTime("2020-01-01")},
			},
			expected: []Book{
				{Title: "Book A", PublishedAt: parseTime("2020-01-01")},
			},
		},
		{
			name: "two books with the same oldest date",
			input: []Book{
				{Title: "Book A", PublishedAt: parseTime("2020-04-04")},
				{Title: "Book B", PublishedAt: parseTime("2020-04-04")},
				{Title: "Book C", PublishedAt: parseTime("2021-01-01")},
			},
			expected: []Book{
				{Title: "Book A", PublishedAt: parseTime("2020-04-04")},
				{Title: "Book B", PublishedAt: parseTime("2020-04-04")},
			},
		},
		{
			name: "Multiple books with different dates",
			input: []Book{
				{Title: "Book A", PublishedAt: parseTime("2021-01-01")},
				{Title: "Book B", PublishedAt: parseTime("2020-01-01")},
				{Title: "Book C", PublishedAt: parseTime("2022-01-01")},
			},
			expected: []Book{
				{Title: "Book B", PublishedAt: parseTime("2020-01-01")},
			},
		},
		{
			name:     "Empty list",
			input:    []Book{},
			expected: []Book{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterByOldestDate(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d books, got %d", len(tt.expected), len(result))
			}
			for i, book := range result {
				if book.Title != tt.expected[i].Title || !book.PublishedAt.Equal(tt.expected[i].PublishedAt) {
					t.Errorf("expected book %v, got %v", tt.expected[i], book)
				}
			}
		})
	}
}
