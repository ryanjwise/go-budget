package io

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveJSONToFile(t *testing.T) {
	testData := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	tempFile, e := ioutil.TempFile("", "test_data.json")
	if e != nil {
		t.Fatalf("Error saving JSON to file: %v", e)
	}
	defer os.Remove((tempFile.Name()))

	e = SaveJSONToFile(testData, tempFile.Name())
	if e != nil {
		t.Fatalf("Error saving JSON to file: %v", e)
	}

	jsonData, e := ioutil.ReadFile(tempFile.Name())
	if e != nil {
		t.Fatalf("Error reading JSON rom file: %v", e)
	}

	expectedJSON, e := json.Marshal(testData)
	if e != nil {
		t.Fatalf("Error marshaling expected JSON: %v", e)
	}

	var actualJSON bytes.Buffer
	json.Compact(&actualJSON, jsonData)

	if !bytes.Equal(actualJSON.Bytes(), expectedJSON) {
		t.Errorf("Expected JSON: \n%s\n Actual JSON: \n%s", string(expectedJSON), actualJSON.String())
	}
}
