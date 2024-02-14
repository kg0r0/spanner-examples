package cmd

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
)

func BatchUpdate(ctx context.Context, w io.Writer, client *spanner.Client) error {
	insertSQL := `INSERT INTO Singers (SingerId, FirstName, LastName) 
	VALUES (@p1, @p2, @p3)`
	stmts := []spanner.Statement{
		{
			SQL: insertSQL,
			Params: map[string]interface{}{
				"p1": 1,
				"p2": "Alice",
				"p3": "Henderson",
			},
		},
		{
			SQL: insertSQL,
			Params: map[string]interface{}{
				"p1": 2,
				"p2": "Bruce",
				"p3": "Allison",
			},
		},
	}
	var updateCounts []int64
	var err error
	_, err = client.ReadWriteTransaction(context.Background(), func(ctx context.Context, transaction *spanner.ReadWriteTransaction) error {
		updateCounts, err = transaction.BatchUpdate(ctx, stmts)
		return err
	})
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Inserted %v singers\n", updateCounts)
	return nil
}
