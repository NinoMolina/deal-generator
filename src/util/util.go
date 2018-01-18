package util

import (
	"log"
	"encoding/json"
	"io/ioutil"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	CheckErr(err)
	return string(bytes)
}

func ReadJsonFile(jsonFile string) []byte {
	raw, err := ioutil.ReadFile(jsonFile)
	CheckErr(err)
	return raw
}