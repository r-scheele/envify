package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Converts an environment variables file to JSON format.

// Parameters:
// - file_path (string): Path of the environment variables file.
// - output_file (string): Path of the output JSON file.

// Returns:
// - error: Returns an error if reading or writing the files fails.

// The function reads the environment variables from the file, parses them, and
// writes them to the output file in JSON format.
func Env2JSON(file_path, output_file string) error {

	env, err := ReadFile(file_path)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	lines := strings.Split(env, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}
		key := strings.ToLower(parts[0])
		value := strings.Replace(parts[1], "\t", "", -1)
		updateMap(data, key, value)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	result := string(jsonData)

	err = WriteFile(output_file, result)
	if err != nil {
		return err
	}

	return nil

}

// Updates a nested map with a new key-value pair.

// Parameters:
// - data (map[string]interface{}): The map to update.
// - key (string): The key for the new value.
// - value (string): The value to add to the map.

// Returns:
// - None

// The function creates a new key-value pair in the nested map using the specified key and value.
// If the key has a '_' delimiter, it is split into a group and subkey identifier. If the group
// identifier does not exist, a new map is created. If the subkey identifier has further '_' delimiters,
// the function is called recursively to create a nested map. Otherwise, the value is added to the map
// using the subkey as the key.
func updateMap(data map[string]interface{}, key string, value string) {
	keys := strings.SplitN(key, "_", 2)
	if len(keys) < 2 {
		return
	}
	group := keys[0]
	subkey := keys[1]
	if _, ok := data[group]; !ok {
		data[group] = make(map[string]interface{})
	}
	if strings.Contains(subkey, "_") {
		updateMap(data[group].(map[string]interface{}), subkey, value)
	} else {
		data[group].(map[string]interface{})[subkey] = value
	}
}
