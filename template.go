package robi

import (
	"bytes"
	"html/template"
)

type RobiTemplate struct {
}

func (rt *RobiTemplate) Execute(tmpContent string, data interface{}) interface{} {
	t, err := template.New("main").Parse(tmpContent)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}

	var w bytes.Buffer
	err = t.Execute(&w, data)
	if err != nil {
		return map[string]interface{}{
			"errmsg": err.Error(),
		}
	}
	return w.String()
}
