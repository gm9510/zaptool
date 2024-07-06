package main

import (
	"encoding/csv"
	"fmt"
	"gocsvop/dataframe"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	colsOps := args[1]
	colsOpsArray := strings.Split(colsOps, ",")

	_ = args[2]
	fileName := args[3]
	_ = args[4]
	sizeStr := args[5]

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		panic(err)
	}

	fileHandler, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	csvfile := csv.NewReader(fileHandler)
	alldata, err := csvfile.ReadAll()
	if err != nil {
		panic(err)
	}

	df := dataframe.New(alldata)

	dfend := df.Select(colsOpsArray...).Limit(size)

	fmt.Println(dfend)
}
