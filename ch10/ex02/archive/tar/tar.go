package tar

import (
	"archive/tar"
	"io"
	"os"

	"github.com/shrhdk/gopl/ch10/ex02/archive"
)

// Reader provies implementation of archive.Reader
type Reader struct {
	*tar.Reader
}

// Next advances to the next entry in the tar archive.
func (r *Reader) Next() (fileInfo os.FileInfo, err error) {
	header, err := r.Reader.Next()
	if err != nil {
		return nil, err
	}

	return header.FileInfo(), nil
}

// NewReader creates a new Reader reading from r.
func NewReader(r *io.SectionReader) archive.Reader {
	return &Reader{tar.NewReader(r)}
}

func init() {
	archive.RegisterFormat("tar", []byte{'u', 's', 't', 'a', 'r', 0}, 257, NewReader)
	archive.RegisterFormat("tar", []byte{'u', 's', 't', 'a', 'r', 0, '4', '0', 0, '4', '0'}, 257, NewReader)
}
