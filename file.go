package robi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type File struct {
}

func (f *File) LoadJson(file string) interface{} {

	jsonFile, err := os.Open(file)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	return result

}

func (f *File) WriteJson(file string, data interface{}) interface{} {
	jsonString, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	err = ioutil.WriteFile(file, jsonString, os.ModePerm)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return nil
}

func (f *File) Exist(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false

	} else {
		panic(err)
	}
}
