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
	stmt := spanner.Statement{SQL: "SELECT * FROM Singers"}
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
		var singerID int64
		var firstName, lastName string
		var singerInfo []byte
		if err := row.Columns(&singerID, &firstName, &lastName, &singerInfo); err != nil {
			panic(err)
		}
		fmt.Printf("%d %s %s\n", singerID, firstName, lastName)
	}
}
