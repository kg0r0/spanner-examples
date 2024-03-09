package cmd

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func Query(ctx context.Context, w io.Writer, client *spanner.Client) error {
	var firstName, lastName string

	stmt := spanner.Statement{SQL: "SELECT FirstName, LastName FROM Singers WHERE SingerId = 1"}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			panic(err)
		}
		if err := row.Columns(&firstName, &lastName); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s %s\n", firstName, lastName)
	}
}

// Ref: https://pkg.go.dev/cloud.google.com/go/spanner@v1.55.0#RowIterator.Do
func QueryRowIteratorDo(ctx context.Context, w io.Writer, client *spanner.Client) error {
	var firstName, lastName string

	stmt := spanner.Statement{SQL: "SELECT FirstName, LastName FROM Singers WHERE SingerId = 1"}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	err := iter.Do(func(row *spanner.Row) error {
		if err := row.Columns(&firstName, &lastName); err != nil {
			return err
		}
		fmt.Fprintf(w, "%s %s\n", firstName, lastName)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
