# spanner-examples

```
e.g.)
$  go run main.go readrow "projects/[PROJECT_ID]/instances/[INSTANCE_NAME]/databases/[TABLE_NAME]"
```

```
$ docker-compose up -d 
$ docker-compose exec spanner-cli spanner-cli -p test -i test -d test
Connected.
spanner>
```

```
$ export SPANNER_EMULATOR_HOST=localhost:9011
```

## References
- https://pkg.go.dev/cloud.google.com/go/spanner#section-readme
- https://pkg.go.dev/cloud.google.com/go/spanner#hdr-Single_Reads
- https://github.com/GoogleCloudPlatform/golang-samples/blob/main/functions/spanner/spanner.go
- https://github.com/GoogleCloudPlatform/golang-samples/tree/main/spanner/spanner_snippets
- https://cloud.google.com/spanner/docs/getting-started/go?hl=ja
- https://cloud.google.com/spanner/docs/emulator#install-docker