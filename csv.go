package osutils

import (
	"encoding/csv"
	"io"
)

func (o *osutils) CSVNewReader(r io.Reader) *csv.Reader {
	return csv.NewReader(r)
}

func (o *osutils) CSVRead(csvReader *csv.Reader) (record []string, err error) {
	return csvReader.Read()
}
