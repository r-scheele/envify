package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Json2Env(file_path, output_file string) error {

	jsonStr, err := ReadFile(file_path)
	if err != nil {
		return err
	}

	var jsonObj interface{}
	json.Unmarshal([]byte(jsonStr), &jsonObj)

	result := make(map[string]interface{})
	FlattenJSON("", jsonObj, result)

	var str string
	for key, val := range result {

		str += fmt.Sprintf("%s=%v\n", strings.ToUpper(key), val)

	}
	err = WriteFile(output_file, str)
	if err != nil {
		return err
	}
	return nil
}

func FlattenJSON(prefix string, jsonObj interface{}, result map[string]interface{}) {

	switch jsonObj := jsonObj.(type) {
	case map[string]interface{}:
		for key, val := range jsonObj {
			if prefix == "" {
				FlattenJSON(key, val, result)
			} else {
				FlattenJSON(fmt.Sprintf("%s_%s", prefix, key), val, result)
			}
		}
	default:
		result[strings.ReplaceAll(prefix, ".", "_")] = jsonObj
	}

}
