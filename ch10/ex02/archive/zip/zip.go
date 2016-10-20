package zip

import (
	"archive/zip"
	"io"
	"os"

	"github.com/shrhdk/gopl/ch10/ex02/archive"
)

// Reader provies implementation of archive.Reader
type Reader struct {
	*zip.Reader
	i int
}

// file returns current file
func (r *Reader) file() *zip.File {
	return r.Reader.File[r.i]
}

// Next advances to the next entry in the tar archive.
func (r *Reader) Next() (fileInfo os.FileInfo, err error) {
	r.i++
	if r.i == len(r.Reader.File) {
		return nil, io.EOF
	}

	return r.file().FileInfo(), nil
}

// Read reads from the current entry in the zip archive.
// It returns 0, io.EOF when it reaches the end of that entry,
// until Next is called to advance to the next entry.
func (r *Reader) Read(b []byte) (n int, err error) {
	rc, err := r.file().Open()
	if err != nil {
		return 0, err
	}

	defer rc.Close()
	return rc.Read(b)
}

// NewReader creates a new Reader reading from r.
func NewReader(r *io.SectionReader) archive.Reader {
	zipReader, _ := zip.NewReader(r, r.Size())
	return &Reader{zipReader, -1}
}

func init() {
	archive.RegisterFormat("zip", []byte{'P', 'K', 3, 4}, 0, NewReader)
	archive.RegisterFormat("zip", []byte{'P', 'K', 3, 6}, 0, NewReader)
}
