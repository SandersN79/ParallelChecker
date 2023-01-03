# Parallel Uniqueness Check
CLI Tool that checks for duplicate csv file values. This tool uses no third party dependencies.

## Install CLI
* go install
```bash
$ go install -v github.com/SandersN79/parallelChecker/cmd/checker@latest
$ checker -h
```
* go build
```bash
$ git clone https://github.com/SandersN79/parallelChecker.git
$ cd parallelChecker
$ go build ./cmd/checker
$ ./checker -h
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
