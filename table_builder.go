package tabler

type TableBuilder struct {
	// Headers are a ordered list of column header names
	headers []string
	// Data holds the value of each row in unordered key/value pairs
	data    []map[string]string
}

func (tableBuilder TableBuilder) headerIndex(targetHeader string) int {
	for i, header := range tableBuilder.headers {
		if header == targetHeader {
			return i
		}
	}
	return -1
}

func (tableBuilder *TableBuilder) AddColumn(header string, callback ColumnCallback) {
	//TODO: Append unless exists? Overwrite or duplicate?
	tableBuilder.headers = append(tableBuilder.headers, header)
	for i, row := range tableBuilder.data {
		firstColumnName := tableBuilder.headers[0]
		firstColumnValue := row[firstColumnName]
		row[header] = callback(firstColumnValue)
		value := callback(firstColumnValue)
		row[header] = value
		tableBuilder.data[i] = row
	}
}

func (tableBuilder *TableBuilder) AddRow(firstValue string, callback RowCallback) {
	rowData := make(map[string]string)
	firstColumnName := tableBuilder.headers[0]

	for _, header := range tableBuilder.headers {
		if header == firstColumnName {
			//	First column
			rowData[firstColumnName] = firstValue
		} else {
			rowData[header] = callback(header, firstValue)
		}
	}
	tableBuilder.data = append(tableBuilder.data, rowData)
}

func (tableBuilder TableBuilder) Build() Table {
	var allRowsData [][]string
	for _, row := range tableBuilder.data {
		var rowData = make([]string, len(tableBuilder.headers))
		for j, headerName := range tableBuilder.headers {
			rowData[j] = row[headerName]
		}
		allRowsData = append(allRowsData, rowData)
	}
	return Table{
		Headers: tableBuilder.headers,
		Rows: allRowsData,
	}
}

type ColumnCallback func(string) string
type RowCallback func(string, string) string
