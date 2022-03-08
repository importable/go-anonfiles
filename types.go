package anonfiles

import "regexp"

var reId = regexp.MustCompile(`(https?:\/\/)?(anonfiles.com\/)(([a-zA-Z0-9][0-f]){5})(\/.*|\W|$)`)

type Info struct {
	Status bool `json:"status"`
	Data   Data `json:"data"`
}
type URL struct {
	Full  string `json:"full"`
	Short string `json:"short"`
}
type Size struct {
	Bytes    int    `json:"bytes"`
	Readable string `json:"readable"`
}
type Metadata struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size Size   `json:"size"`
}
type File struct {
	URL      URL      `json:"url"`
	Metadata Metadata `json:"metadata"`
}
type Data struct {
	File File `json:"file"`
}
