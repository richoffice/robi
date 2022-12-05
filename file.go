package robi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/richoffice/richframe"
	"gopkg.in/yaml.v3"
)

type File struct {
}

func (f *File) Base(file string) string {
	return filepath.Base(file)
}

func (f *File) RelativePath(base string, originBase string, originPath string) string {
	baseDir := filepath.Dir(base)
	relBase, err := filepath.Rel(baseDir, originBase)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	newBase := filepath.Dir(relBase)
	newRel := filepath.Join(newBase, originPath)

	return strings.ReplaceAll(newRel, "\\", "/")
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

func (f *File) WriteExcelByTemp(data map[string]richframe.RichFrame, targetFile string, tempFile string, defPath string) interface{} {

	err := richframe.ExportRichFramesByTemp(data, targetFile, tempFile, defPath, nil)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return nil
}

func (f *File) LoadCSV(srcFile string, keys []string) interface{} {
	rf, err := richframe.LoadCSV(srcFile, keys)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return rf
}

func (f *File) WriteCSV(csvpath string, rf richframe.RichFrame, keys []string, isAppend bool) interface{} {

	err := richframe.SaveCSV(csvpath, rf, keys, isAppend)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return nil
}

func (f *File) LoadFile(file string) interface{} {
	b, err := os.ReadFile(file) // just pass the file name
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	str := string(b)
	return str
}

func (f *File) WriteFile(file string, content string, isAppend bool) interface{} {
	var fi *os.File
	var err error
	if isAppend {
		fi, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return map[string]interface{}{
				"errmsg": err.Error(),
			}
		}
	} else {
		fi, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return map[string]interface{}{
				"errmsg": err.Error(),
			}
		}
	}
	defer fi.Close()
	if _, err := fi.WriteString(content); err != nil {
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

func (f *File) Mkdir(path string) interface{} {
	err := os.MkdirAll(path, os.ModePerm)
	return err
}

func (f *File) List(path string) interface{} {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	return files
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
