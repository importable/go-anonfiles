package anonfiles

// Anonfiles.com API wrapper
// written in Go programming language

import "fmt"

func GetId(url string) (string, error) {
	var id string
	submatch := reId.FindStringSubmatch(url)
	if len(submatch) > 4 {
		id = submatch[3]
	}

	if id == "" {
		err := fmt.Errorf("not found file id in url %s", url)
		return id, err
	}

	return id, nil
}
