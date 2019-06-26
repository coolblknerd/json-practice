package encodeprac

import (
	"encoding/json"
	"fmt"
)

// Album represents a piece of art
type Album struct {
	Title  string
	Artist string
}

// CreateAlbum ruturns a album object
func CreateAlbum(title string, artist string) *Album {
	return &Album{
		Title:  title,
		Artist: artist,
	}
}

// CreateJSON returns a json data printed out
func (a *Album) CreateJSON() {
	jsonData, err := json.Marshal(&a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))
}
