package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golangR99/models"
	"os"
)

func WriteToFile[T models.ObjectDB](item T) string {
	itemFolder := fmt.Sprintf("./_%s", item.GetType())
	itemName := fmt.Sprintf("%s.json", item.GetName())
	if _, err := os.Stat(itemFolder); os.IsNotExist(err) {
		// create character/ psychube folder if it doesn't already exist
		if err := os.MkdirAll(itemFolder, os.ModeAppend); err != nil {
			return err.Error()
		}
	}
	// if possible to convert item to bytes, write it to file
	if itemBytes, err := convertToBytes(item); err != nil {
		return fmt.Sprintf("%s could not be written.", itemName)
	} else {
		fullPath := fmt.Sprintf("%s/%s", itemFolder, itemName)
		os.WriteFile(fullPath, itemBytes, 0644)
		return fmt.Sprintf("%s created.", itemName)
	}
}

func convertToBytes[T models.ObjectDB](item T) ([]byte, error) {
	var buf bytes.Buffer
	// set encoder
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)  // retain symbols
	encoder.SetIndent("", "    ") // indent = 4 spaces
	// encode data
	if err := encoder.Encode(item); err != nil {
		return nil, err
	}
	// remove newline from encoder and return
	result := bytes.TrimSpace(buf.Bytes())
	return result, nil
}
