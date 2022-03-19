package anonfiles

// Anonfiles.com API wrapper
// written in Go programming language

import (
	"encoding/json"
	"fmt"
)

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

func GetInfo(url string) (Info, error) {
	// File info structure
	var info Info

	// Get information about the file
	fileId, err := GetId(url)
	if err != nil {
		return info, err
	}

	// Make request with file id
	resp, err := client.Get(fmt.Sprintf(apiURL, fileId))
	if err != nil {
		return info, err
	}

	defer resp.Body.Close()

	rawJSON := resp.Body

	err = json.NewDecoder(rawJSON).Decode(&info)
	if err != nil {
		return info, err
	}

	return info, nil
}

func GetStatus(url string) (bool, error) {
	var info Info

	info, err := GetInfo(url)
	if err != nil {
		return false, err
	}

	return info.Status, nil
}

func GetFullURL(url string) (string, error) {
	var info Info

	info, err := GetInfo(url)
	if err != nil {
		return "", err
	}

	return info.Data.File.URL.Full, nil
}

func GetShortURL(url string) (string, error) {
	var info Info

	info, err := GetInfo(url)
	if err != nil {
		return "", err
	}

	return info.Data.File.URL.Short, nil
}

func GetName(url string) (string, error) {
	var info Info

	info, err := GetInfo(url)
	if err != nil {
		return "", err
	}

	return info.Data.File.Metadata.Name, nil
}

func GetSizeBytes(url string) (int, error) {
	var info Info

	info, err := GetInfo(url)
	if err != nil {
		return 0, err
	}

	return info.Data.File.Metadata.Size.Bytes, nil
}

func GetSizeReadable(url string) (string, error) {
	var info Info

	info, err := GetInfo(url)
	if err != nil {
		return "0 B", err
	}

	return info.Data.File.Metadata.Size.Readable, nil
}

// func Upload(path string) (string, error) {

// }
