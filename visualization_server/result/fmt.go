package result

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	resultsDir = "results"
)

// get caller for directory name
// get time for name
// TODO: use write sequence

//type resultMap []map[string]float64

var (
	//resultSet       resultMap
	fileName        string
	file            *os.File
	isFieldsWritten bool
	//keys            []string
)

func init() {
	_, err := ioutil.ReadDir("../" + resultsDir)
	if err != nil {
		err := os.Mkdir("../"+resultsDir, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func createFile() {
	t := time.Now()
	ts := strconv.FormatInt(t.UnixNano(), 10)
	pc, _, line, ok := runtime.Caller(4)
	if !ok {
		//pc, _, line, ok = runtime.Caller(4)
		//if !ok {
		//	panic("fmt init runtime.caller")
		//}
	}

	function := runtime.FuncForPC(pc)
	functionName := function.Name()

	s := strings.Split(functionName, "go_test/")
	fmt.Println(s)
	//fileName = filepath.Base(caller) + "_" + strconv.Itoa(line) + "_" + ts + ".csv"

	var name string
	if len(s) > 1 {
		name = s[1]
	} else {
		name = s[0]
	}
	fileName = name + "_" + strconv.Itoa(line) + "_" + ts + ".csv"

	var err error
	file, err = os.Create("../" + resultsDir + "/" + fileName)
	if err != nil {
		panic(err) // "Error creating file for results",
	}
}

func Results(a ...interface{}) {
	defer fmt.Print(a)
	results(a...)
}

func results(a ...interface{}) {
	if file == nil {
		createFile()
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var s []string
	for _, v := range a {
		s = append(s, fmt.Sprintf("%v", v))
	}
	writer.Write(s)
}

func Resultsf(format string, a ...interface{}) {
	defer fmt.Printf(format, a...) // Print result to console anyway
	results(a...)
}

func ResultsKV(a ...interface{}) {
	// check keys
	size := len(a)
	if size%2 != 0 {
		fmt.Println("Wrong params len. Use [Key, Value] or Resultsf() function.")
		return
	}

	if !isFieldsWritten {
		// get keys
		var keys []interface{}
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				keys = append(keys, fmt.Sprint(a[i]))
			}
		}
		results(keys...)
		isFieldsWritten = true
	}

	var values []interface{}
	for i := 0; i < size; i++ {
		if i%2 != 0 {
			values = append(values, a[i])
		}
	}

	results(values...)
}
