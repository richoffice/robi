package robi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/richoffice/richframe"
	"gopkg.in/yaml.v3"
)

type File struct {
}

func (f *File) LoadExcel(srcFile string, defPath string) interface{} {
	rf, err := richframe.LoadRichFrames(srcFile, defPath, nil)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return rf
}

func (f *File) WriteExcel(data map[string]richframe.RichFrame, targetFile string, defPath string) interface{} {

	err := richframe.ExportRichFrames(data, targetFile, defPath, nil)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return nil
}

func (f *File) LoadYaml(file string) interface{} {
	yamlFile, err := os.Open(file)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	defer yamlFile.Close()

	byteValue, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	var result map[string]interface{}
	err = yaml.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return result
}

func (f *File) WriteYaml(file string, data interface{}) interface{} {
	yamlString, err := yaml.Marshal(data)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	err = ioutil.WriteFile(file, yamlString, os.ModePerm)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return nil
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

func (f *File) Glob(pattern string) interface{} {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	return files
}
