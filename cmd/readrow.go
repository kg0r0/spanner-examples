package cmd

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
)

func ReadRow(ctx context.Context, w io.Writer, client *spanner.Client, _ []string) error {
	var firstName, lastName string
	row, err := client.Single().ReadRow(ctx, "Singers", spanner.Key{12}, []string{"FirstName", "LastName"})
	if err != nil {
		return err
	}
	if err := row.Columns(&firstName, &lastName); err != nil {
		return err
	}
	fmt.Printf("%s %s\n", firstName, lastName)
	return nil
}
