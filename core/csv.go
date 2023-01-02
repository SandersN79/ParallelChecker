package core

import (
	"bytes"
	"errors"
	"os"
)

type csvData struct {
	barcode  string
	code     string
	yearWeek string
}

func (d *csvData) parse(in []byte, pMap map[string]int) {
	tmp := bytes.Split(in, []byte(","))
	if len(tmp) > 2 {
		d.barcode = string(tmp[pMap["barcode"]])
		d.code = string(tmp[pMap["code"]])
		d.yearWeek = string(tmp[pMap["YearWeek"]])
	}
	return
}

type csvFile struct {
	fileName string
	csvData  []*csvData
}

// getParseMap builds a parseMap for the headers of the CSV file
func getParseMap(sDat [][]byte) (map[string]int, error) {
	pMap := make(map[string]int, 3)
	if len(sDat) == 0 {
		return pMap, errors.New("no data in csv file")
	}
	for idx, header := range bytes.Split(sDat[0], []byte(",")) {
		switch string(header) {
		case "barcode":
			pMap["barcode"] = idx
		case "code":
			pMap["code"] = idx
		case "YearWeek":
			pMap["YearWeek"] = idx
		}
	}
	return pMap, nil
}

func newCsvFile(filePath string) (*csvFile, error) {
	f := &csvFile{fileName: filePath}
	var data []*csvData
	dat, err := os.ReadFile(f.fileName)
	if err != nil {
		return f, err
	}
	sDat := bytes.Split(dat, []byte("\n"))
	pMap, err := getParseMap(sDat)
	if err != nil {
		return f, err
	}
	for i, _ := range sDat {
		if i > 0 && i < len(sDat) {
			var d csvData
			d.parse(sDat[i], pMap)
			data = append(data, &d)
		}
	}
	f.csvData = data
	return f, nil
}
