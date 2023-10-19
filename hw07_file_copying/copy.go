package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

var (
	ErrFileDoesNotExist      = errors.New("file does not exist")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	err := CheckFile(fromPath)
	if err != nil {
		return err
	}

	fFrom, err := os.Open(fromPath)
	defer fFrom.Close()
	if err != nil {
		return ErrUnsupportedFile
	}

	fTo, err := os.Create(toPath)
	defer fTo.Close()
	if err != nil {
		return ErrUnsupportedFile
	}

	fContent, _ := fFrom.Stat()
	if fContent.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit+offset > fContent.Size() {
		limit = fContent.Size()
	}

	bOffset, err := fFrom.Seek(offset, 0)
	if err != nil || bOffset != offset {
		return ErrUnsupportedFile
	}

	bSize := int64(5)
	pBar := pb.StartNew(int(limit))

	for i := int64(0); i < limit; i += bSize {
		w, err := io.CopyN(fTo, fFrom, bSize)
		if (err != nil && err != io.EOF) || w > limit {
			return err
		}
		pBar.Add(int(i))
	}
	pBar.Finish()
	return err
}

func CheckFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return ErrFileDoesNotExist
		}
	}

	return nil
}
