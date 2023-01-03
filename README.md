# Parallel Uniqueness Check
CLI Tool that checks for duplicate csv file values. This tool uses no third party dependencies.

## Install CLI
```bash
$ go install github.com/SandersN79/parallelChecker@latest
```

## Usage
#### A) Check included csv files in testData
```bash
$ checker
```
#### B) Specific folder directory and/or csv files to check
```bash
$ checker -path=/dir/files/
```
```bash
$ checker -files=file.csv,file2.csv
```
```bash
$ checker -path=/dir/files/ -files=file.csv,file2.csv
```

#### C) CLI Help
```bash
$ checker -h 
```

## Test CLI
```bash
$ go test ./cmd/checker -v
```
