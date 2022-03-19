package anonfiles

import (
	"strings"
	"testing"
)

var availableFileLink = "https://anonfiles.com/bbraP6M4x6/Screenshot_from_2022-03-09_02-01-02_png"
var unavailableFileLink = "https://anonfiles.com/u1C0ebc4b0/file.txt"

func TestGetId(t *testing.T) {
	// Possible links
	lines := `	https://anonfiles.com/u1C0ebc4b0  
	https://anonfiles.com/u1C0ebc4b0/hello-+world
http://anonfiles.com/u1C0ebc4b0
	http://anonfiles.com/u1C0ebc4b0/#(HFOEN effinf
	anonfiles.com/u1C0ebc4b0/eoriorine`

	links := strings.Split(lines, "\n")

	for _, link := range links {

		id, err := GetId(link)
		if err != nil {
			t.Logf("Error: '%v'", err)
			t.Fail()
		}

		if id != "u1C0ebc4b0" {
			t.Logf("Error: parsed %s instead of 'u1C0ebc4b0'", id)
			t.Fail()
		}
	}

}

func TestGetInfo(t *testing.T) {
	information, err := GetInfo(availableFileLink)
	if err != nil {
		t.Logf("Error: '%v'", err)
		t.Fail()
	}

	if information.Status == false {
		t.Logf("Error: status should be true. Link is %s", availableFileLink)
		t.Fail()
	}

	information, err = GetInfo(unavailableFileLink)
	if err != nil {
		t.Logf("Error: '%v'", err)
		t.Fail()
	}

	if information.Status == true {
		t.Logf("Error: status should be false. Link is %s", unavailableFileLink)
		t.Fail()
	}
}

func TestGetStatus(t *testing.T) {
	status, err := GetStatus(availableFileLink)
	if err != nil {
		t.Logf("Error: '%v'", err)
		t.Fail()
	}

	if status == false {
		t.Logf("Error: file status should be true. Link is %s", availableFileLink)
		t.Fail()
	}

	status, err = GetStatus(unavailableFileLink)
	if err != nil {
		t.Logf("Error: '%v'", err)
		t.Fail()
	}

	if status == true {
		t.Logf("Error: file status should be false. Link is %s", unavailableFileLink)
		t.Fail()
	}
}

func TestGetSizeBytes(t *testing.T) {
	size, err := GetSizeBytes(availableFileLink)
	if err != nil {
		t.Logf("Error: '%v'", err)
		t.Fail()
	}

	if size == 0 {
		t.Logf("Error: file size shouldn't be 0. Link is %s", availableFileLink)
		t.Fail()
	}

	size, err = GetSizeBytes(unavailableFileLink)
	if err != nil {
		t.Logf("Error: '%v'", err)
		t.Fail()
	}

	if size != 0 {
		t.Logf("Error: file status should be 0. Link is %s", unavailableFileLink)
		t.Fail()
	}
}
