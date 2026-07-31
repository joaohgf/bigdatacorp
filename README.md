# BigDataCorp Batch Challenge

Streams football-club records from JSONL and generates `clubs.csv` and `players.csv` without loading the full input into memory.

## Requirements

- Go 1.26.5 or compatible

## Run

```bash
go run ./cmd/cli example/sample_clubes.jsonl
```

Choose output paths with flags:

```bash
go run ./cmd/cli --clubs-output output/clubs.csv --players-output output/players.csv example/sample_clubes.jsonl
```

The input path is required. Missing or invalid inputs return a nonzero exit status.

## Data handling

- Only Série A and Série B clubs are emitted.
- Malformed JSONL lines, null clubs, clubs without `club_id`, null players, and players without `player_id` are skipped.
- Invalid dates and missing/null fields become empty CSV fields.
- Colors are joined with `|`.
- CSV output is UTF-8 and escaped with Go's RFC 4180-compatible `encoding/csv` package.

## Large sample

Generate an ignored 250,000-club fixture with four players per club:

```bash
go run ./cmd/cli generate-sample
```

Customize it with `--clubs`, `--players`, and `--output`. The default 250 MB fixture contains one million nested player records plus malformed and incomplete examples. A measured run completed in approximately 3.3 seconds with about 12 MiB peak resident memory.

## Bonus API

Start the API:

```bash
go run ./cmd/api
```

Upload JSONL and receive a ZIP containing both CSV files:

```bash
curl -F file=@example/sample_clubes.jsonl http://localhost:8080/api/v1/upload --output result.zip
```

Customize response names:

```bash
curl -F file=@example/sample_clubes.jsonl \
  'http://localhost:8080/api/v1/upload?clubs-output=clubes&players-output=jogadores&archive-output=entrega' \
  --output entrega.zip
```

Uploads are limited to 1 GiB and processed in request-scoped temporary directories.

## Verification

```bash
go test ./...
go vet ./...
```
