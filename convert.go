package convert

import (
	"log"
	"regexp"

	"github.com/Huang-Jinxian/book-convert-go/txt"
	"golang.org/x/text/encoding"
)

func Convert(title string, author string, filePath string, coverPath string, chapterTitleReg *regexp.Regexp, fileEncoding encoding.Encoding, dstFilePath string) error {
	txtBook := txt.NewTxtBook(title, author, filePath, coverPath, chapterTitleReg)

	if dstFilePath == "" {
		txtBook.SetDstFilePath(dstFilePath)
	}

	if fileEncoding != nil {
		txtBook.SetFileEncoding(fileEncoding)
	}

	err := txtBook.Parse()
	if err != nil {
		log.Printf("can't parse the book: %s", err)
		return err
	}

	err = txtBook.Convert2Epub()
	if err != nil {
		return err
	}

	return nil

}
