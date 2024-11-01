package main

import (
	"encoding/json"
	"net/http"
)

// HandleMin [POST] given an array of floats return the min
func HandleMin(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var numbers Numbers

	err := decoder.Decode(&numbers)

	if err != nil {
		panic(err)
	}

	minNumber := Min(numbers.Numbers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonRes, _ := json.Marshal(&Response{Result: minNumber})

	w.Write(jsonRes)
}

// HandleMax [POST] given an array of floats return the max
func HandleMax(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var numbers Numbers

	err := decoder.Decode(&numbers)

	if err != nil {
		panic(err)
	}

	maxNumber := Max(numbers.Numbers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonRes, _ := json.Marshal(&Response{Result: maxNumber})

	w.Write(jsonRes)
}

// HandleMax [POST] given an array of floats return the average
func HandleAvg(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var numbers Numbers

	err := decoder.Decode(&numbers)

	if err != nil {
		panic(err)
	}

	// Prevent division by zero
	var average float64 = 0

	if len(numbers.Numbers) != 0 {
		average = Average(numbers.Numbers)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonRes, _ := json.Marshal(&Response{Result: average})

	w.Write(jsonRes)
}

// HandleMax [POST] given an array of floats return its median
func HandleMedian(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var numbers Numbers

	err := decoder.Decode(&numbers)

	if err != nil {
		panic(err)
	}

	// Prevent division by zero
	median := Median(numbers.Numbers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonRes, _ := json.Marshal(&Response{Result: median})

	w.Write(jsonRes)
}

// HandleMax [POST] given an array of floats return its median
func HandlePercentile(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var numbers Numbers

	err := decoder.Decode(&numbers)

	if err != nil {
		panic(err)
	}

	// Prevent division by zero
	perc := Percentile(numbers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonRes, _ := json.Marshal(&Response{Result: perc})

	w.Write(jsonRes)
}
