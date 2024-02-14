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
)

type command func(ctx context.Context, w io.Writer, client *spanner.Client) error

var (
	commands = map[string]command{
		"readrow":     cmd.ReadRow,
		"query":       cmd.Query,
		"batchupdate": cmd.BatchUpdate,
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
	flag.Parse()
	if len(flag.Args()) < 2 || len(flag.Args()) > 3 {
		flag.Usage()
		os.Exit(2)
	}

	cmd, db, arg := flag.Arg(0), flag.Arg(1), flag.Arg(2)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if err := run(ctx, os.Stdout, cmd, db, arg); err != nil {
		os.Exit(1)
	}
}
