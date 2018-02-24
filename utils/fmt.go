package utils

import (
	"encoding/json"
)

func JsonStr2Map(data interface{}) (map[string]interface{}) {
	var d map[string]interface{}
	if str, ok := data.(string); ok {
		json.Unmarshal([]byte(str), &d)
	}
	return d
}

func Map2JsonStr(data interface{}) string {
	if _, ok := data.(map[string]interface{}); ok {
		s, _ := json.Marshal(data)
		return string(s)
	}
	return ""
}

func Struct2Map(obj interface{}) map[string]interface{} {
	s, _ := json.Marshal(obj);
	var data = make(map[string]interface{})
	json.Unmarshal(s, &data)
	return data
}
