/*
Package mem manages the memory file.
*/
package mem

/*Data defines the memory file layout.

Total bytes: 12

version					- 12 C		- memory file game matching version.
screenHeight			- 2 B		- screen vertical size in pixels
screenWidth				- 2 B		- screen horizontal size in pixels
*/
type Data struct {
	version      [12]common.Char
	screenHeight uint16
	screenWidth  uint16
}
