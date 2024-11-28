package cmd

import (
	"context"
	"io"

	"cloud.google.com/go/spanner"
	"github.com/kg0r0/spanner-examples/models"
)

// Ref: https://github.com/googleapis/google-cloud-go/blob/main/spanner/examples_test.go#L164
func UpdateStruct(ctx context.Context, w io.Writer, client *spanner.Client) error {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		m, err := spanner.UpdateStruct("Singers", &models.Singers{
			SingerId:  12,
			LastName:  "Test",
			FirstName: "Test2",
		})
		if err != nil {
			return err
		}
		return txn.BufferWrite([]*spanner.Mutation{m})
	})
	if err != nil {
		return err
	}
	return nil
}
