package formats

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
)

//func readSampleCSV() (*tabler.Table, error) {
//	sourceCSVFile, err := os.Open("testdata/sample.csv")
//	inputTable, err := ReadCSV(sourceCSVFile)
//	if (err != nil) {
//		return nil, err
//	}
//	return inputTable, nil
//}


func TestTerminal_Render(t *testing.T) {
	var renderer Renderer
	var outputBuffer bytes.Buffer
	renderer = Terminal{
		Writer: &outputBuffer,
	}
	inputTable, err := readSampleCSV(); if (err != nil) {
		panic(err)
	}
	err = renderer.Render(*inputTable); if (err != nil) {
		panic(err)
	}
	expectedContect := strings.Join([]string{
		"+----+------+------------+",
		"| ID | SIZE |   STATUS   |",
		"+----+------+------------+",
		"|  1 | M    | Done       |",
		"|  2 | S    | InProgress |",
		"|  3 | L    | Done       |",
		"+----+------+------------+",
		"",
	}, "\n")
	assert.Equal(t, expectedContect, outputBuffer.String())
}
