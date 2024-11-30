package cmd

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
)

// Ref: https://cloud.google.com/spanner/docs/introspection/troubleshooting-with-tags#transaction_tags
func ReadWriteTransactionWithTag(ctx context.Context, w io.Writer, client *spanner.Client, _ []string) error {
	_, err := client.ReadWriteTransactionWithOptions(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `UPDATE Venues SET Capacity = CAST(Capacity/4 AS INT64) WHERE OutdoorVenue = false`,
		}
		_, err := txn.UpdateWithOptions(ctx, stmt, spanner.QueryOptions{RequestTag: "app=concert,env=dev,action=update"})
		if err != nil {
			return err
		}
		fmt.Fprint(w, "Venue capacities updated.\n")
		stmt = spanner.Statement{
			SQL: `INSERT INTO Venues (VenueId, VenueName, Capacity, OutdoorVenue, LastUpdateTime) 
			VALUES (@venueId, @venueName, @capacity, @outdoorVenue, PENDING_COMMIT_TIMESTAMP())`,
			Params: map[string]interface{}{
				"venueId":      81,
				"venueName":    "Venue 81",
				"capacity":     1440,
				"outdoorVenue": true,
			},
		}
		_, err = txn.UpdateWithOptions(ctx, stmt, spanner.QueryOptions{RequestTag: "app=concert,env=dev,action=insert"})
		if err != nil {
			return err
		}
		fmt.Fprint(w, "New venue inserted.\n")
		return nil
	}, spanner.TransactionOptions{TransactionTag: "app=concert,env=dev"})
	return err
}

/*
[How to view transaction tags in Transaction Statistics table]
SELECT t.fprint,
       t.transaction_tag,
       t.read_columns,
       t.commit_attempt_count,
       t.avg_total_latency_seconds
FROM SPANNER_SYS.TXN_STATS_TOP_10MINUTE AS t
LIMIT 3;

[Sample Output]
fprint              transaction_tag     read_columns                                       commit_attempt_count avg_total_latency_seconds
768230775380599424                      Venues._exists                                     1                    1.990432922
848985723542418304  app=concert,env=dev Venues.Capacity,Venues.OutdoorVenue,Venues._exists 2                    0.5142774985
1112437037500834816                     Venues._exists                                     1                    1.372883303
*/
