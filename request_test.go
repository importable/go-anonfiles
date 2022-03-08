package anonfiles

import (
	"strings"
	"testing"
)

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
