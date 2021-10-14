package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadFile(filepath string) []byte {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Failed to read data file")
	}
	return b
}

func ReadInFile() *database {
	var d *database
	err := json.Unmarshal(ReadFile("base.json"), &d)
	if err != nil {
		log.Fatal("Failed to restore data file")
	}
	return d
}

func StoreOutFile() {
	D = ReadInFile()
	if D.RulebreakCount == nil {
		D.RulebreakCount = make(map[int64]int)
	}
	if D.NotruledCount == nil {
		D.NotruledCount = make(map[int64]int)
	}
}
