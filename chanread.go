package netsuitePres

import (
	"fmt"
	"sync"
)

func read () {
	someChan := make(chan struct{})
	someChan <- struct{}{}
	for rt := range someChan {
		fmt.Println(rt)
	}
}

func closeChan(){
	linesChan := make(chan ConsumeReportLine)
	customersChan := make(chan CustomerLinesItem)
	customersChan <- CustomerLinesItem{}
	defer close(linesChan)

	var wg sync.WaitGroup // HLwg
	go func() {
		AppendConsumeLineToCSVWG(linesChan, &wg) // HLwg
	}()

	for customer := range customersChan {
		wg.Add(1) // HLwg
		fmt.Println(customer)
		linesChan <- ConsumeReportLine{}
	}

	wg.Wait() // HLwg
}

func AppendConsumeLineToCSVWG(linesChan chan ConsumeReportLine, wg *sync.WaitGroup) {
	for reportLine := range linesChan {
		fmt.Println(reportLine.ContractName)
		wg.Done() // HLwg
	}
}

func AppendConsumeLineToCSV(linesChan chan ConsumeReportLine) {
	for reportLine := range linesChan {
		fmt.Println(reportLine.ContractName)
	}
}

func readAndSync() {
	customersChan := make(chan CustomerLinesItem) // HLcustomersChan
	linesChan := make(chan ConsumeReportLine)
	defer close(linesChan)

	go func() {
		AppendConsumeLineToCSV(linesChan)
	}()

	numberOfCustomers, _ := GetCustomersAmount()

	for i := 0; i < int(numberOfCustomers); i++ {
		customerLines := <-customersChan  // HLcustomersChan
		if customerLines.Err != nil {}

		for _, line := range customerLines.Lines {
			linesChan <- line
		}
	}
}

func GetCustomersAmount() (int, error) {
	return 0, nil
}
