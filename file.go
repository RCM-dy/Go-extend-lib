package Go_extend_lib

import (
	"bytes"
	"encoding/json"
	"os"
)

func WriteBytes(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
func FmtJsonBytes(data []byte) ([]byte, error) {
	var s bytes.Buffer
	err := json.Indent(&s, data, "", "    ")
	if err != nil {
		return []byte(""), err
	}
	return s.Bytes(), nil
}
func WriteFmtJsonBytes(filename string, jsondata []byte) error {
	jsonb, err := FmtJsonBytes(jsondata)
	if err != nil {
		return err
	}
	return WriteBytes(filename, jsonb)
}
