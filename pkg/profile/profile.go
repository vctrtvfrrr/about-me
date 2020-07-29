package profile

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vctrtvfrrr/vctrtvfrrr/pkg/info"
)

func getFileContent(filename string) string {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	return text
}

// Info returns the formated profile description.
func Info() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		getFileContent("data/about.md"),
		info.GetInfo(),
		getFileContent("data/curiosities.md"),
	)
}
