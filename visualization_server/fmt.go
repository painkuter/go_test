package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

// get caller for directory name
// get time for name
// TODO: use write sequence

type resultMap []map[string]float64

var (
	resultSet       resultMap
	fileName        string
	file            *os.File
	isFieldsWritten bool
)

func init() {
	t := time.Now()
	ts := strconv.FormatInt(t.UnixNano(), 10)
	_, caller, line, ok := runtime.Caller(0)
	if !ok {
		panic("fmt init runtime.caller")
	}

	fileName = filepath.Base(caller) + "_" + strconv.Itoa(line) + "_" + ts + ".csv"

	var err error
	file, err = os.Create("../" + resultsDir + "/" + fileName)
	if err != nil {
		panic(err) // "Error creating file for results",
	}
}

//
//func results() *resultMap {
//
//	if resultSet
//	return &resultMap{}
//}

// TODO ResultsKV()
func Results(format string, a ...interface{}) {
	defer fmt.Printf(format, a...) // Print result to console anyway

	//size := len(a)
	//if size%2 != 0 {
	//	fmt.Println("Wrong params len. Use [Key, Value]")
	//	return
	//}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var s []string
	for _, v := range a {
		s = append(s, fmt.Sprintf("%v", v))
	}
	fmt.Println(s)
	writer.Write(s)
}

//func ResultsKV(format string, a ...interface{}){
//	Results(a)
//}