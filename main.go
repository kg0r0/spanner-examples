package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func main() {
	projectid := os.Getenv("GCLOUD_PROJECT")
	db := fmt.Sprintf("projects/%s/instances/test/databases/test", projectid)
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	var firstName, lastName string

	// ReadRow
	row, err := client.Single().ReadRow(ctx, "Singers", spanner.Key{1}, []string{"FirstName", "LastName"})
	if err != nil {
		panic(err)
	}
	if err := row.Columns(&firstName, &lastName); err != nil {
		panic(err)
	}
	fmt.Printf("%s %s\n", firstName, lastName)

	// Statement
	stmt := spanner.Statement{SQL: "SELECT FirstName, LastName FROM Singers WHERE SingerId = 1"}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return
		}
		if err != nil {
			panic(err)
		}
		if err := row.Columns(&firstName, &lastName); err != nil {
			panic(err)
		}
		fmt.Printf("%s %s\n", firstName, lastName)
	}
}
