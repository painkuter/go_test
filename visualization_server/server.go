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

// This is application for visualization of experiments results
// PLAN:
// Read all files from group to memory
// Calculate size
// Select needed values
// Serve it

const (
	resultsDir = "results"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/results", getResults)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("visualization_server/static"))))

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", r)
}

func getResults(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//response := fmt.Sprintf("Product %s", id)

	files, err := ioutil.ReadDir("./" + resultsDir)
	if err != nil {
		log.Fatal(err)
	}
	var data [][]string
	for _, f := range files {
		data = dataFromFile("./" + resultsDir + "/" + f.Name())
	}

	buf, err := json.Marshal(convertToSeries(data))
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

func convertToSeries(data [][]string) [][]float64 {
	count := len(data[0]) // number of series
	result := make([][]float64, count)
	for i := 0; i < count; i++ {
		arr := make([]float64, len(data))
		result[i] = arr
	}

	for k1, v1 := range data {
		for k2, v2 := range v1 {
			value, err := strconv.ParseFloat(v2, 64)
			if err != nil {
				panic(err)
			}
			result[k2][k1] = value
		}
	}
	return result
}

func convertToMap(data [][]float64) map[string][]float64 {
	result := make(map[string][]float64)
	for k, v := range data {
		result[strconv.Itoa(k)] = v
	}
	return result
}
