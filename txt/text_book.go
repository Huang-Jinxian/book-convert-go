package txt

import (
	"bufio"
	"os"
	"regexp"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type txtBook struct {
	Title           string
	Author          string
	FilePath        string
	CoverPath       string
	ChapterTitleReg *regexp.Regexp

	Chapters []chapter

	FileEncoding encoding.Encoding
}

type chapter struct {
	Title    string
	Sections []string
}

func NewTxtBook(title string, author string, filePath string, coverPath string, chapterTitleReg *regexp.Regexp) *txtBook {
	return &txtBook{Title: title, Author: author, FilePath: filePath, CoverPath: coverPath, ChapterTitleReg: chapterTitleReg}
}

func (book *txtBook) SetFileEncoding(enc encoding.Encoding) {
	book.FileEncoding = enc
}

func (book *txtBook) Parse() error {
	bookFile, err := os.Open(book.FilePath)
	if err != nil {
		return err
	}
	defer func(bookFile *os.File) {
		_ = bookFile.Close()
	}(bookFile)

	var decoder *encoding.Decoder
	if book.FileEncoding == nil {
		decoder = simplifiedchinese.GB18030.NewDecoder()
	} else {
		decoder = book.FileEncoding.NewDecoder()
	}
	fileReader := transform.NewReader(bookFile, decoder)
	fileScanner := bufio.NewScanner(fileReader)

	ch := chapter{Title: book.Title}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if book.ChapterTitleReg != nil && book.ChapterTitleReg.MatchString(line) {
			book.Chapters = append(book.Chapters, ch)
			ch = chapter{Title: line}
		} else {
			ch.Sections = append(ch.Sections, line)
		}
	}
	book.Chapters = append(book.Chapters, ch)

	return nil
}
