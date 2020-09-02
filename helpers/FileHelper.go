package helpers

import "io/ioutil"


func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "tpl not load!", err
	}
	return string(bytes), nil
}