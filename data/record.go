/*
Package data manages the memory file.
*/
package data

import (
	"fmt"
	"sync"
)

var (
	Memory *record
	once   sync.Once
)

/*
Record defines the memory file layout. !!!! The order of fields is important, when the data is decoded. !!!!

	Total bytes: 12
	-------------------------------------------------------------------------------------
	version					- 12 B		- memory file game matching version.
	screenHeight			-  2 B		- screen vertical size in pixels
	screenWidth				-  2 B		- screen horizontal size in pixels
*/
type record struct {
	ScreenWidth  uint16
	ScreenHeight uint16
}

//create a singleton of the game memory data
func instance() *record {
	once.Do(func() {
		Memory = new(record)
		fmt.Println("created memory data instance")
	})
	return Memory
}
