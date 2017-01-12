package tabler

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestTableBuilder_AddColumn(t *testing.T) {
	tableBuilder := TableBuilder{
		headers: []string{"name"},
		data: []map[string]string{
			map[string]string{"name": "test"},
			map[string]string{"name": "foo"},
			map[string]string{"name": "bar"},
			map[string]string{"name": "abcdefghijklmnopqrstuvwxyz"},
		},
	}
	lettersInName := func(header string) string {
		return strconv.Itoa(len(header))
	}

	tableBuilder.AddColumn("nameLength", lettersInName)

	table := tableBuilder.Build()
	assert.Equal(t, table.Headers, []string{"name", "nameLength"})
	assert.Equal(t, [][]string{
		[]string{"test", "4"},
		[]string{"foo", "3"},
		[]string{"bar", "3"},
		[]string{"abcdefghijklmnopqrstuvwxyz", "26"}}, table.Rows)
}

func TestTableBuilder_AddRow(t *testing.T) {
	tableBuilder := TableBuilder{
		headers: []string{"name", "nameLength", "reverse"},
	}

	rowValues := func(header string, value string) string {
		var retValue string
		switch header {
		case "nameLength":
			retValue = strconv.Itoa(len(value))
		case "reverse":
			runes := []rune(value)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			retValue = string(runes)
		}

		return retValue
	}

	tableBuilder.AddRow("test", rowValues)
	tableBuilder.AddRow("foo", rowValues)
	tableBuilder.AddRow("bar", rowValues)
	tableBuilder.AddRow("abcdefghijklmnopqrstuvwxyz", rowValues)

	table := tableBuilder.Build()
	assert.Equal(t, table.Headers, []string{"name", "nameLength", "reverse"})
	assert.Equal(t, [][]string{
		[]string{"test", "4", "tset"},
		[]string{"foo", "3", "oof"},
		[]string{"bar", "3", "rab"},
		[]string{"abcdefghijklmnopqrstuvwxyz", "26", "zyxwvutsrqponmlkjihgfedcba"}}, table.Rows)
}
