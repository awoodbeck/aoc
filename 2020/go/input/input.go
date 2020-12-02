package input

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

const dataPath = "../../data"

func Read(day int) ([]string, error) {
	bs, err := ioutil.ReadFile(fmt.Sprintf("%s/%d.txt", dataPath, day))
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var data []string

	for _, b := range bytes.Split(bytes.TrimSpace(bs), []byte("\n")) {
		data = append(data, string(bytes.TrimSpace(b)))
	}

	return data, nil
}
