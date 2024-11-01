package main

func main() {
	server := NewServer(5000)

	server.Handle("/min", "POST", HandleMin)
	server.Handle("/max", "POST", HandleMax)
	server.Handle("/avg", "POST", HandleAvg)
	server.Handle("/median", "POST", HandleMedian)
	server.Handle("/percentile", "POST", HandlePercentile)

	server.Listen()
}
