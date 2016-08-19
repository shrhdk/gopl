package archive

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// Reader is reader interface of archive files.
type Reader interface {
	Next() (fileInfo os.FileInfo, err error)
	Read(b []byte) (n int, err error)
}

func match(r *io.SectionReader, magic []byte, offset int64) bool {
	p := make([]byte, len(magic))
	_, err := r.ReadAt(p, offset)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	return reflect.DeepEqual(p, magic)
}

// NewReader creates new ArchiveReader from io.Reader
func NewReader(r *io.SectionReader) (Reader, error) {
	for _, f := range formats {
		if match(r, f.magic, f.offset) {
			fmt.Println(f.name, f.magic)
			return f.newReader(r), nil
		}
	}
	return nil, fmt.Errorf("Unknown Format")
}

type format struct {
	name      string
	magic     []byte
	offset    int64
	newReader func(*io.SectionReader) Reader
}

var formats []format

// RegisterFormat registers new archive format to support.
func RegisterFormat(name string, magic []byte, offset int64, newReader func(r *io.SectionReader) Reader) {
	formats = append(formats, format{name, magic, offset, newReader})
}
