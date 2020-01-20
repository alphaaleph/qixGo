/*
Package data manages the memory file.
*/
package data

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	errDataMissing  = errors.New("no data available")
	errWrongVersion = errors.New("wrong version")
)

//decode will retrive the data from the memory file
func decode(byteData []byte) error {

	//didn't received any data to decode
	if len(byteData) <= 0 {
		return errDataMissing
	}

	//TODO: check for version number, it should match the application version, so it can parse correctly
	//return error if versions do not match ... errWrongVersion

	//read the data
	reader := bytes.NewReader(byteData)

	if err := binary.Read(reader, binary.BigEndian, instance()); err != nil {
		fmt.Println("binary.Read failed:", err)
		return err
	}
	return nil
}

//encode will update the memory file with the Record structure
func encode() ([]byte, error) {

	//load data into slice
	var byteData []byte
	writer := bytes.NewBuffer(byteData)

	//write the data
	if err := binary.Write(writer, binary.BigEndian, instance()); err != nil {
		fmt.Println("binary.Write failed:", err)
		return nil, err
	}
	return writer.Bytes(), nil
}
