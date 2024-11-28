package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/kg0r0/spanner-examples/cmd"
	"github.com/kg0r0/spanner-examples/config"
)

type command func(ctx context.Context, w io.Writer, client *spanner.Client) error

var (
	commands = map[string]command{
		"readrow":            cmd.ReadRow,
		"query":              cmd.Query,
		"batchupdate":        cmd.BatchUpdate,
		"transactiontags":    cmd.ReadWriteTransactionWithTag,
		"queryrowiteratordo": cmd.QueryRowIteratorDo,
		"updatestruct":       cmd.UpdateStruct,
	}
)

func run(ctx context.Context, w io.Writer, cmd string, db string, arg string) error {
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer client.Close()

	cmdfunc := commands[cmd]
	if cmdfunc == nil {
		flag.Usage()
		os.Exit(2)
	}
	err = cmdfunc(ctx, w, client)
	if err != nil {
		fmt.Fprintf(w, "%s failed with %v", cmd, err)
	}
	return err
}

func main() {
	if v := os.Getenv("SPANNER_EMULATOR_HOST"); v == "" {
		fmt.Println("SPANNER_EMULATOR_HOST is not set")
		return
	}
	flag.Parse()
	if len(flag.Args()) < 1 || len(flag.Args()) > 2 {
		flag.Usage()
		os.Exit(2)
	}
	db := fmt.Sprintf("projects/%s/instances/%s/databases/%s", config.ProjectID, config.InstanceName, config.TableName)
	cmd, arg := flag.Arg(0), flag.Arg(1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if err := run(ctx, os.Stdout, cmd, db, arg); err != nil {
		os.Exit(1)
	}
}
