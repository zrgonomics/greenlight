package main

import (
	"os"
)

func CreateFile(filename string) (*os.File, error) {
	return os.Create(filename)
}

func WriteToFile(file *os.File, content string) (int, error) {
	return file.WriteString(content)
}
