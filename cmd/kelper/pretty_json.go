package main

import "encoding/json"

func PrettyJson(value interface{}) {
	json.MarshalIndent(value, "", "\t")
}
