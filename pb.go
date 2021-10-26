package netsuitePres

import (
	"log"
	"sync/atomic"
	"time"
)

type ProgressBar struct {
	currentValue *int32
	finish       chan struct{} // HLfinish
	timeStart    time.Time
	logger 		 log.Logger
}

type Counter interface {
	AddValues(value int)
	Finish()
}

func NewProgressBar() ProgressBar {
	var zero int32

	p := ProgressBar{
		currentValue: &zero,
		finish:       make(chan struct{}), // HLfinish
		logger:       log.Logger{},
		timeStart:    time.Now(),
	}

	go p.run()

	return p
}
// START OMIT
func (p ProgressBar) AddValues(value int) { // HLname
	atomic.AddInt32(p.currentValue, int32(value))
}

func (p ProgressBar) Finish() { // HLname
	p.finish <- struct{}{}
}

func (p ProgressBar) run() { // HLname
	for {
		select {
		case <-p.finish:
			p.logger.Printf("It took %v to make reports", time.Since(p.timeStart))
			p.logger.Printf("Finished. Lines processed: %d", atomic.LoadInt32(p.currentValue))
			return // HLname
		case <-time.After(10 * time.Second):
			p.logger.Printf("Lines processed: %d", atomic.LoadInt32(p.currentValue))
		}
	}
}
// END OMIT
