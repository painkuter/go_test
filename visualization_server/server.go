package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	resultsDir = "results"
)

// This is application for visualization of experiments results
// PLAN:
// Read all files from group to memory
// Calculate size
// Select needed values
// Serve it

type resultResponse []series

type series struct {
	Name string    `json:"name"`
	Data []float64 `json:"data"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/results", getResults)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("visualization_server/static"))))

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8181", r)
	if err != nil {
		panic(err)
	}
}

func getResults(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//response := fmt.Sprintf("Product %s", id)

	files, err := ioutil.ReadDir("./" + resultsDir)
	if err != nil {
		log.Fatal(err)
	}
	//var data [][]string
	data := make(map[string][][]string)
	for _, f := range files {
		data[f.Name()] = dataFromFile("./" + resultsDir + "/" + f.Name())
	}

	buf, err := json.Marshal(convertWithFileName(data))
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(buf))

}

func dataFromFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return data
}

func convertWithFileName(data map[string][][]string) resultResponse {
	var result resultResponse
	for k, v := range data {
		result = append(result, convertToSeries(v, k)...)
	}
	return result
}

func convertToSeries(data [][]string, name string) resultResponse {
	count := len(data[0]) // number of series
	if count == 0 || len(data) == 0 {
		panic("Empty data or data[0]")
	}

	var result2 resultResponse

	result := make([][]float64, count)
	for i := 0; i < count; i++ {
		arr := make([]float64, len(data))
		result[i] = arr
	}

	// check keys
	var keys []string
	_, err := strconv.ParseFloat(data[0][0], 64)
	if err != nil {
		keys = data[0]
		data = data[1:]
	} else {
		for i := 0; i < 0; i++ {
			keys = append(keys, strconv.Itoa(i))
		}
	}

	// add file name
	for i := range keys {
		keys[i] = name + ":" + keys[i]
	}

	// create items

	for i := 0; i < count; i++ {
		arr := make([]float64, len(data))
		result2 = append(result2, series{Name: keys[i], Data: arr})
	}

	for k1, v1 := range data {
		for k2, v2 := range v1 {
			value, err := strconv.ParseFloat(v2, 64)
			if err != nil {
				panic(err)
			}
			result2[k2].Data[k1] = value
		}
	}
	return result2
}

func convertToMap(data [][]float64) map[string][]float64 {
	result := make(map[string][]float64)
	for k, v := range data {
		result[strconv.Itoa(k)] = v
	}
	return result
}

//func checkKeys(data [][]string) []string {
//
//}
