package csvparser

import (
	"encoding/csv"
	"os"
	"errors"
	"log"
	"io"
	"fmt"
	"strings"
	)

type Record struct {
	Question, Answer string
}

func (r *Record) CheckAnswer(ans string) bool {
	result := false
	if strings.Compare(r.Answer, strings.TrimSpace(ans)) == 0 {
		result = true
	}
	return result
}

func Parse(fileName string) ([]*Record, error) {
	records := make([]*Record, 0)
	file, err := os.Open(fileName)
	if err != nil {
		errors.New("File not found")
	}
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		newRecord := &Record {record[0], record[1]}
		records = append(records, newRecord)
	}
	return records, nil
}

