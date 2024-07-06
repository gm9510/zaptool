package dataframe

import (
	"fmt"
	"testing"
)

var dataarray = [][]string{
	{"A", "B", "C"},
	{"1", "2", "3"},
	{"4", "", "nan"},
}

func TestNewDataFrame(t *testing.T) {

	df := New(dataarray)
	fmt.Println(df)

	dfA := df.Select("C", "A")
	fmt.Println(dfA)

}

func TestLimit(t *testing.T) {
	df := New(dataarray)
	fmt.Println(df)

	dfA := df.Select("C", "A").Limit(1)
	fmt.Println(dfA)

}

func TestOffset(t *testing.T) {
	df := New(dataarray)
	fmt.Println(df)

	dfA := df.Select("C", "A").Offset(1)
	fmt.Println(dfA)

}
