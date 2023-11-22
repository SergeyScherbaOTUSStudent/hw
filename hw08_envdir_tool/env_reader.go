package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrDoesNotExist    = errors.New("dir does not exist")
	ErrUnsupportedFile = errors.New("unsupported file")
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func NewEnvValue(val string, needRemove bool) EnvValue {
	return EnvValue{
		Value:      val,
		NeedRemove: needRemove,
	}
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, ErrDoesNotExist
	}

	env := make(Environment)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if false == info.IsDir() {
			cont, _ := ReadFileJustFirstLine(path)
			env[info.Name()] = NewEnvValue(cont, true)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return env, nil
}

func ReadFileJustFirstLine(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", ErrUnsupportedFile
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	var res []byte
	for fileScanner.Scan() {
		if i != 0 {
			continue
		}
		res = fileScanner.Bytes()
		i++
	}

	res = bytes.ReplaceAll(res, []byte{0x00}, []byte{0x0a})
	res = bytes.TrimRight(res, " ")

	return string(res), err
}
