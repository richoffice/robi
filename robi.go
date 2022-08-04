package robi

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/richoffice/richframe"
	"github.com/spf13/viper"
)

type Robi struct {
	Base   string
	Engine RuleEngine
	Mailer *Mailer
}

func NewRobi(base string) (*Robi, error) {

	store, err := NewStoreFromDirectory(base)
	if err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(base)
	if err != nil {
		return nil, err
	}

	robi := &Robi{
		Base: absPath,
	}

	var vars map[string]interface{} = map[string]interface{}{
		"robi": robi,
		"log":  fmt.Println,
	}

	engine := NewMemoryRuleEngine(store, vars)
	robi.Engine = engine

	err = GetConfig(absPath)
	if err != nil {
		fmt.Println("no config file")
		return nil, err
	}

	host := viper.GetString("mail.smtp")
	port := viper.GetInt("mail.port")
	user := viper.GetString("mail.user")
	password := viper.GetString("mail.password")
	mailer := NewMailer(host, port, user, password)
	robi.Mailer = mailer

	return robi, nil
}

func GetConfig(absPath string) error {
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(absPath, "conf"))

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func (robi *Robi) NewRichFrames() map[string]*richframe.RichFrame {
	return make(map[string]*richframe.RichFrame, 0)
}

func (robi *Robi) LoadDict(name string) interface{} {
	excelPath := filepath.Join(robi.Base, "dicts", name+".xlsx")
	defPath := filepath.Join(robi.Base, "dicts", name+".json")
	rf, err := richframe.LoadRichFrames(excelPath, defPath, nil)
	if err != nil {
		panic(err)
	}
	return rf
}

func (robi *Robi) Import(defPath string, srcFile string) interface{} {
	fullPath := defPath

	if !strings.HasPrefix(defPath, "/") {
		fullPath = filepath.Join(robi.Base, "defs", defPath)
	}
	rf, err := richframe.LoadRichFrames(srcFile, fullPath, nil)
	if err != nil {
		panic(err)
	}
	return rf
}

func (robi *Robi) Export(data map[string]*richframe.RichFrame, defPath string, targetFile string) interface{} {
	fullPath := defPath
	if !strings.HasPrefix(defPath, "/") {
		fullPath = filepath.Join(robi.Base, "defs", defPath)
	}
	err := richframe.ExportRichFrames(data, targetFile, fullPath, nil)
	if err != nil {
		panic(err)
	}

	return nil
}

func (robi *Robi) Execute(task string, args interface{}) (interface{}, error) {
	return robi.Engine.ApplyRule(task, args)
}

func (robi *Robi) Now() time.Time {
	return time.Now()
}

func (robi *Robi) FromUnix(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func (robi *Robi) MonthStart(in time.Time) time.Time {
	return time.Date(in.Year(), in.Month(), 0, 0, 0, 0, 0, in.Location())
}

func (robi *Robi) MonthEnd(in time.Time) time.Time {
	return time.Date(in.Year(), in.Month(), 0, 0, 0, 0, 0, in.Location()).AddDate(0, 1, 0)
}

func (robi *Robi) ExeTemplate(tpl string, data interface{}) string {

	tmp := template.Must(template.ParseFiles(filepath.Join(robi.Base, "tpl", tpl)))

	var w bytes.Buffer
	tmp.Execute(&w, data)

	return w.String()

}

func (robi *Robi) WriteFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	return err
}
