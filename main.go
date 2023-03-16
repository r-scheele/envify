/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// func main() {
// 	// Example JSON input
// input := `
// {
// 	"database": {
// 		"host": "localhost",
// 		"port": 5432,
// 		"username": "test_user",
// 		"password": "test_password",
// 		"name": "test_db"
// 	},
// 	"server": {
// 		"host": "localhost",
// 		"port": 8080,
// 		"debug": true
// 	},
// 	"logging": {
// 		"level": "info",
// 		"filename": "/var/log/myapp.log"
// 	}
// }
// `

// 	// Flatten the JSON data using gjson
// 	flatData := utils.FlattenJSON(input)
// 	fmt.Println(flatData)

// 	// Write flattened data to .env file
// 	// for key, value := range flatData {
// 	// 	fmt.Printf("%s=%s\n", strings.ToUpper(key), value)
// 	// }

// 	// cmd.Execute()
// }

package main

import (
	"github.com/r-scheele/cmd"
)

func main() {
	cmd.Execute()
}
