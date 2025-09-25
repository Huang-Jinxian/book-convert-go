package util

import (
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

// detectEnc used to detect src's encoding, default is simplifiedchinese.GB18030
func detectEnc(src []byte) *encoding.Encoding {
	var enc encoding.Encoding
	det := chardet.NewTextDetector()
	best, err := det.DetectBest(src)
	if err != nil {
		return &simplifiedchinese.GB18030
	}
	switch best.Charset {
	case "UTF-8":
		enc = unicode.UTF8
	case "EUC-JP":
		enc = japanese.EUCJP
	case "EUC-KR":
		enc = korean.EUCKR
	case "Big5":
		enc = traditionalchinese.Big5
	case "GBK":
	case "GB18030":
	default:
		enc = simplifiedchinese.GB18030
	}
	return &enc
}
