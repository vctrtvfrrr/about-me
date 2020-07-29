package latestbooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Books struct which contains an array of books
type Books struct {
	Books []Book `json:"books"`
}

// Book struct which contains a book details
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "…"
	}
	return bnoden
}

// LatestBooks returns a list of the last books readed
func LatestBooks() (result [10]string, err error) {
	jsonFile, err := os.Open("data/latest_books.json")
	defer jsonFile.Close()
	if err != nil {
		return
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	var b Books
	json.Unmarshal(byteValue, &b)

	maxResults := 10
	if len(b.Books) < maxResults {
		maxResults = len(b.Books)
	}

	for i := 0; i < maxResults; i++ {
		book := b.Books[i]
		title := truncateString(book.Title, 50)
		result[i] = fmt.Sprintf("%s\t–\t_%s_", title, book.Author)
	}

	return
}
