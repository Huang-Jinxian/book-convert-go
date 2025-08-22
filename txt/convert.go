package txt

import (
	"fmt"
	"path"
	"strings"

	"github.com/go-shiori/go-epub"
)

const (
	sectionTitleTemplate    = `<h2>%s</h2>`
	sectionBodyTemplate     = `<p>%s</p>`
	sectionFileNameTemplate = `part_%d`
	epubFileTemplate        = `%s.epub`
)

func (book *txtBook) Convert2Epub(dstPath string) error {
	newBook, _ := epub.NewEpub(book.Title)
	newBook.SetAuthor(book.Author)

	if book.CoverPath != "" {
		image, err := newBook.AddImage(book.CoverPath, "cover.jpg")
		if err != nil {
			return err
		}
		err = newBook.SetCover(image, "")
		if err != nil {
			return err
		}
	}

	for i, ch := range book.Chapters {
		var body string
		var sectionBody = append([]string{fmt.Sprintf(sectionTitleTemplate, ch.Title)})
		for _, section := range ch.Sections {
			sectionBody = append(sectionBody, fmt.Sprintf(sectionBodyTemplate, section))
		}

		body = strings.Join(sectionBody, "\n")

		_, err := newBook.AddSection(body, ch.Title, fmt.Sprintf(sectionFileNameTemplate, i), "")
		if err != nil {
			return err
		}
	}

	destFile := path.Join(dstPath, fmt.Sprintf(epubFileTemplate, book.Title))
	err := newBook.Write(destFile)
	return err
}
