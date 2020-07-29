package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
	"github.com/vctrtvfrrr/vctrtvfrrr/pkg/latestbooks"
	"github.com/vctrtvfrrr/vctrtvfrrr/pkg/profile"
	"github.com/vctrtvfrrr/vctrtvfrrr/pkg/topartists"
)

func generateArtistsAndBooksTable() string {
	// Search for books and artists
	books, err := latestbooks.LatestBooks()
	if err != nil {
		log.Fatal(err)
	}
	artists, err := topartists.TopArtists()
	if err != nil {
		log.Fatal(err)
	}

	tableString := &strings.Builder{}

	table := tablewriter.NewWriter(tableString)

	table.SetHeader([]string{"🎧 Artistas da semana", "📚 Últimos livros lidos"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)

	for i := 0; i < 10; i++ {
		table.Append([]string{artists[i], books[i]})
	}

	table.Render()

	return tableString.String()
}

func getUpdateDatetime() string {
	dt := time.Now().UTC()
	return dt.Format("02-01-2006 15:04:05 MST")
}

func init() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	file, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	readme := fmt.Sprintf(
		"%s\n\n## O que eu tenho feito\n\n%s\n\n---\n\n🚀 **Última atualização:** %s",
		profile.Info(),
		generateArtistsAndBooksTable(),
		getUpdateDatetime(),
	)

	file.WriteString(readme)
}
