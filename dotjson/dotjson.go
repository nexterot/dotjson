// Package dotjson provides DotFormat function for formatting json objects
// with "dot syntax".
package dotjson

import (
	"encoding/json"
	"reflect"
)

// DotFormat replaces json objects with composite keys delimited by dot in
// hierarchical order. It is guaranteed keys in input json don't contain
// dots. Example: {"a":{"b":1}} -> {"a.b":1}.
func DotFormat(jsonStr string) (string, error) {
	jsonMap := make(map[string]interface{})
	// Decode json
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		return "", err
	}
	// Process json
	jsonMap = dotFormat(jsonMap)
	// Encode json
	jsonBytes, err := json.MarshalIndent(jsonMap, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// dotFormat wraps dotFormatInner and returns new map with processed json.
func dotFormat(jsonMap map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	// dotFormatInner writes keys to jsonMap
	var dotFormatInner func(jsonMap map[string]interface{}, keySuffix string)
	dotFormatInner = func(jsonMap map[string]interface{}, keySuffix string) {
		for k, v := range jsonMap {
			if v != nil && reflect.TypeOf(v).Kind() == reflect.Map {
				// Recur if type of v is map
				dotFormatInner(v.(map[string]interface{}), keySuffix+k+".")
			} else {
				// Otherwise write key to resulting map
				res[keySuffix+k] = v
			}
		}
	}
	dotFormatInner(jsonMap, "")
	return res
}
