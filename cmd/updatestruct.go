package cmd

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
	"github.com/kg0r0/spanner-examples/config"
)

func UpdateStruct(ctx context.Context, w io.Writer, client *spanner.Client) error {
	m, err := spanner.UpdateStruct(config.TableName, client)
	if err != nil {
		return err
	}
	fmt.Printf("Mutation: %v\n", m)
	return nil
}
