package netsuitePres

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
)

func (s *Datastore) IterateThroughCustomers(ctx context.Context, startDate, endDate time.Time) <-chan CustomerItem {
	// START OMIT
	ch := make(chan CustomerItem)

	go func() {
		defer close(ch)

		cur, err := s.db.Collection(richTransactions).Aggregate(ctx, []bson.M{
			{"$match": bson.M{}},
		})
		if err != nil {
			ch <- CustomerItem{Err: err}
			return
		}

		for cur.Next(ctx) {
			val := struct {
				CustomerID string `bson:"_id"`
			}{}
			if err := cur.Decode(&val); err != nil {}

			ch <- CustomerItem{ID: val.CustomerID}
		}

		if err := cur.Err(); err != nil {}
	}()
	// END OMIT

	return ch
}
