package main

import (
	"errors"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrFileDoesNotExist      = errors.New("file does not exist")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var err error
	err = CheckFile(fromPath)

	if err != nil {
		return err
	}

	bar := pb.StartNew(100)

	finalBoard := int(limit)
	fContent, readErr := os.ReadFile(fromPath)

	if readErr != nil {
		return ErrUnsupportedFile
	}

	if len(fContent) < int(offset) {
		return ErrOffsetExceedsFileSize
	}

	if offset > 0 {
		finalBoard = int(offset + limit)
	}

	if len(fContent) < int(offset+limit) || limit == 0 {
		finalBoard = len(fContent)
	}

	printPB(bar, 50)

	toFile, _ := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE, 0o666)

	defer func() {
		printPB(bar, 100)
		toFile.Close()
	}()

	_, err = toFile.Write(fContent[offset:finalBoard])
	printPB(bar, 96)

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

func printPB(bar *pb.ProgressBar, n int) {
	bar.Add(n)
	time.Sleep(time.Millisecond * 10)
}
