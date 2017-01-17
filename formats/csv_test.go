package formats

import (
	"testing"
	"bytes"
	"os"
	"github.com/maxlinc/tabler"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func TestReadCSV(t *testing.T) {
	table, err := readSampleCSV()
	if (err != nil) {
		panic(err)
	}
	assert.Equal(t, table.Headers, []string{"ID", "Size", "Status"})
	assert.Equal(t, [][]string{
		[]string{"1", "M", "Done"},
		[]string{"2", "S", "InProgress"},
		[]string{"3", "L", "Done"},
	}, table.Rows)
}

func readSampleCSV() (*tabler.Table, error) {
	sourceCSVFile, err := os.Open("testdata/sample.csv")
	inputTable, err := ReadCSV(sourceCSVFile)
	if (err != nil) {
		return nil, err
	}
	return inputTable, nil
}

func sampleCSVContent() (string, error) {
	bytes, err := ioutil.ReadFile("testdata/sample.csv")
	if (err != nil) {
		return "", err
	}
	return string(bytes), nil
}

func TestCsvRenderer_Render(t *testing.T) {
	var renderer Renderer
	var outputBuffer bytes.Buffer
	renderer = CSV{
		Writer: &outputBuffer,
	}
	inputTable, err := readSampleCSV(); if (err != nil) {
		panic(err)
	}
	err = renderer.Render(*inputTable); if (err != nil) {
		panic(err)
	}
	expectedCSVContent, err := sampleCSVContent(); if (err != nil) {
		panic(err)
	}
	assert.Equal(t, expectedCSVContent, outputBuffer.String())
}