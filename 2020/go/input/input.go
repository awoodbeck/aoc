package input

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

const dataPath = "../../data"

func ReadAll(day int) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("%s/%d.txt", dataPath, day))
}

func ReadBytes(day int) ([][]byte, error) {
	bs, err := ReadAll(day)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var data [][]byte

	for _, b := range bytes.Split(bytes.TrimSpace(bs), []byte("\n")) {
		data = append(data, bytes.TrimSpace(b))
	}

	return data, nil
}

func ReadStrings(day int) ([]string, error) {
	bs, err := ReadBytes(day)
	if err != nil {
		return nil, err
	}

	data := make([]string, 0, len(bs))
	for _, b := range bs {
		data = append(data, string(b))
	}

	return data, nil
}
