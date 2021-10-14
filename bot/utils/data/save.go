package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type database struct {
	RuledUsers     []int64       `json:"ruled_users"`
	RulebreakCount map[int64]int `json:"rulebreak_count"`
	NotruledCount  map[int64]int `json:"notruled_count"`
}

var D *database

func SaveInFile() {
	b, err := json.Marshal(D)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = ioutil.WriteFile("base.json", b, 0)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
