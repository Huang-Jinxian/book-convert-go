package txt

import (
	"fmt"
	"image"
	"os"
	"path"
	"strings"

	"github.com/leotaku/mobi"
)

const (
	azwSectionTitleTemplate = `<h1>%s</h1>`
	azwSectionBodyTemplate  = `<p>%s</p>`
	azwFileTemplate         = `%s.azw3`
)

func (book *txtBook) Convert2Azw() error {

	azwChapters := make([]mobi.Chapter, len(book.Chapters))

	for _, ch := range book.Chapters {
		sectionBuilder := strings.Builder{}
		sectionBuilder.WriteString(fmt.Sprintf(azwSectionTitleTemplate, ch.Title))

		for _, section := range ch.Sections {
			sectionBuilder.WriteString("\n")
			sectionBuilder.WriteString(fmt.Sprintf(azwSectionBodyTemplate, section))
		}

		azwChapter := mobi.Chapter{Title: ch.Title, Chunks: mobi.Chunks(sectionBuilder.String())}
		azwChapters = append(azwChapters, azwChapter)
	}

	coverFile, err := os.Open(book.CoverPath)
	if err != nil {
		return err
	}
	defer func(coverFile *os.File) {
		_ = coverFile.Close()
	}(coverFile)

	cover, _, err := image.Decode(coverFile)

	azwBook := mobi.Book{
		Title:      book.Title,
		Authors:    []string{book.Author},
		Chapters:   azwChapters,
		CoverImage: cover,
	}

	// converts a Book to a PalmDB Database.
	db := azwBook.Realize()

	// create destination file path
	dstPath := "./"
	if book.DstFilePath != "" {
		dstPath = book.DstFilePath
	}
	err = os.MkdirAll(dstPath, os.ModePerm)
	if err != nil {
		return err
	}
	// create azw3 file
	file, err := os.Create(path.Join(dstPath, fmt.Sprintf(azwFileTemplate, book.Title)))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// write file
	err = db.Write(file)
	if err != nil {
		return err
	}
	return nil
}
