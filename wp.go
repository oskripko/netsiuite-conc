package netsuitePres

// START OMIT
func main() {
	customersChan := make(chan CustomerLinesItem)
	jobs := make(chan string, WorkerPoolSize)
	for i := 1; i <= WorkerPoolSize; i++ {
		go worker(customersChan, jobs)
	}
}

func worker(linesChan chan CustomerLinesItem, jobsChan chan string) {
	for customerID := range jobsChan {
		getLinesForCustomer(linesChan, customerID)
	}
}

func getLinesForCustomer(linesChan chan CustomerLinesItem, id string) {
	//Process something
}
// END OMIT
