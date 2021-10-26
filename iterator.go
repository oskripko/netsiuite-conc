package netsuitePres

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
)
// START OMIT
func (s *Datastore) IterateThroughCustomers(
	ctx context.Context, startDate, endDate time.Time,
) <-chan CustomerItem { // HLiter
	ch := make(chan CustomerItem) // HLiter

	go func() {  // HLiter
		defer close(ch) // HLiter

		cur, err := s.db.Collection(richTransactions).Aggregate(ctx, []bson.M{
			{"$match": bson.M{}},
		})
		if err != nil {
			ch <- CustomerItem{Err: err} // HLiter
			return
		}

		for cur.Next(ctx) {
			...
			ch <- CustomerItem{ID: val.CustomerID} // HLiter
		}
	}()

	return ch
}
// END OMIT
