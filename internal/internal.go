package internal

import (
	"bytes"
	"os"
)

func ReadFileLines(filePath string) ([]string, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var bufStrLines = []string{}
	for _, v := range bytes.Split(fileBytes, []byte("\n")) {
		if len(v) > 0 {
			bufStrLines = append(bufStrLines, string(v))
		}
	}
	return bufStrLines, nil
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
