package dataframe

import "fmt"

func New(matrix [][]string) DataFrame {

	var headers []string
	if len(matrix) > 0 {
		headers = matrix[0]
	}

	var data [][]string
	if len(matrix) > 1 {
		data = matrix[1:]
	}

	return DataFrame{
		headers: headers,
		data:    data,
	}
}

type DataFrame struct {
	headers []string
	data    [][]string
}

func (df DataFrame) String() string {
	strTable := fmt.Sprintf("%v\n", df.headers)

	for _, row := range df.data {
		// Replace with buffers
		strTable = fmt.Sprintf("%v%v\n", strTable, row)
	}

	return strTable
}

func (df DataFrame) Select(cols ...string) DataFrame {
	if len(cols) == 0 {
		return DataFrame{}
	}

	if len(cols) == 1 && cols[0] == "*" {
		return df
	}

	// Deberia fallar si no encuentra una cabezera
	positions := make([]int, 0, len(cols))
	for _, col := range cols {
		for j, header := range df.headers {
			if col == header {
				positions = append(positions, j)
			}
		}
	}

	newData := make([][]string, len(df.data))

	for i, row := range df.data {
		newRow := make([]string, len(positions))
		for j, pos := range positions {
			newRow[j] = row[pos]
		}
		newData[i] = newRow
	}

	return DataFrame{
		data:    newData,
		headers: cols,
	}
}

func (df DataFrame) Limit(lm int) DataFrame {

	ln := len(df.data)

	if ln < lm {
		return df
	}

	return DataFrame{
		headers: df.headers,
		data:    df.data[:lm],
	}
}

func (df DataFrame) Offset(offset int) DataFrame {

	ln := len(df.data)

	if ln < offset {
		return DataFrame{}
	}

	return DataFrame{
		headers: df.headers,
		data:    df.data[offset:],
	}
}
