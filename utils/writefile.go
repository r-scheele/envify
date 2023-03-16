package utils

import "io/ioutil"

func WriteFile(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}
