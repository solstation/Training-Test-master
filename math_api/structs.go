package main

// Numbers is the struct that receive the numbers
type Numbers struct {
	Amount  float64   `json: amount`
	Numbers []float64 `json: numbers`
}

// Response is the struct to send as a response
type Response struct {
	Result float64 `json: result`
}
