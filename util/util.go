package util

import (
	"bytes"
	"encoding/binary"
)

// NumToBytes converts an integer to a byte slice
func NumToBytes(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return []byte{}, err
	}
	return buff.Bytes(), nil
}
