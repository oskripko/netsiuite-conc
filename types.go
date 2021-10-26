package netsuitePres

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerItem struct {
	ID  string
	Err error
}

const (
	richTransactions = "rich_transactions"
	WorkerPoolSize = 30
)


type Datastore struct {
	db *mongo.Database
}

type CustomerLinesItem struct {
	Lines []ConsumeReportLine
	Err   error
}


type ConsumeReportLine struct {
	TransactionID   string
	NSClientID      string
	CreationDate    time.Time
	ExpirationDate  time.Time
	Product         string
	Used            int
	Location        string
	Rate            float64
	IsFree          bool
	TransactionType string
	ContractID      string
	ContractName    string
	Reason          string
}
