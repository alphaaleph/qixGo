/*
Package data manages the memory file.
*/
package data

import (
	"fmt"
	"os"
)

//qix file name

//Read used to read game data
func Read() {

	const cqixdat = "./cqixdat"

	//check if the file doesn't exist, load the default data into memory
	fileInfo, err := os.Stat(cqixdat)
	if os.IsNotExist(err) || fileInfo.Size() <= 0 {
		decode(reset())
		return
	}

	//open the data file
	memFile, err := os.Open(cqixdat)
	defer memFile.Close()
	if err != nil {
		fmt.Println("error opening meme file:", err)
		return
	}

	//read the file
	var buffer = make([]byte, fileInfo.Size())
	if _, err := memFile.Read(buffer); err != nil {
		fmt.Println("error reading byte data from mem file:", err)
		return
	}

	//load into memory
	if err := decode(buffer); err != nil {
		decode(reset())
	}
	return
}

//Save used to save game data
func Save() {

	const cqixdat = "./cqixdat"

	var memFile *os.File
	defer memFile.Close()

	//check if the file doesn't exists, create a new one
	if _, err := os.Stat(cqixdat); os.IsNotExist(err) {
		if memFile, err = os.Create(cqixdat); err != nil {
			fmt.Println("error creating meme file:", err)
			return
		}
	} else {
		if memFile, err = os.Open(cqixdat); err != nil {
			fmt.Println("error opening meme file:", err)
			return
		}
	}

	//save the memory data into the file
	byteData, _ := encode()
	memFile.Write(byteData)

	return
}
