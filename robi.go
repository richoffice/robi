package robi

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/richoffice/richframe"
	"github.com/richoffice/weixinmp"
	"github.com/spf13/viper"
)

type Robi struct {
	Base     string
	Engine   RuleEngine
	Mailer   *Mailer
	Timer    *Timer
	Qrcode   *Qrcode
	Wxmp     *weixinmp.WeixinMp
	File     *File
	Template *RobiTemplate
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
	} else {
		host := viper.GetString("mail.smtp")
		port := viper.GetInt("mail.port")
		user := viper.GetString("mail.user")
		password := viper.GetString("mail.password")
		mailer := NewMailer(host, port, user, password)
		robi.Mailer = mailer

		appid := viper.GetString("wxmp.appid")
		appkey := viper.GetString("wxmp.appkey")
		encodingkey := viper.GetString("wxmp.encodingkey")
		wxmp := weixinmp.NewWeixinMp(appid, appkey, encodingkey)
		robi.Wxmp = wxmp
	}

	robi.Timer = &Timer{}
	robi.Qrcode = &Qrcode{}
	robi.File = &File{}
	robi.Template = &RobiTemplate{}

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

func (robi *Robi) ImportContainMergeCells(defPath string, srcFile string) interface{} {
	fullPath := defPath

	if !strings.HasPrefix(defPath, "/") {
		fullPath = filepath.Join(robi.Base, "defs", defPath)
	}
	rf, err := richframe.LoadRichFramesContainMergeCells(srcFile, fullPath, nil)
	if err != nil {
		panic(err)
	}
	return rf
}

func (robi *Robi) Export(data map[string]richframe.RichFrame, defPath string, targetFile string) interface{} {
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
func (robi *Robi) ExportByTemp(data map[string]richframe.RichFrame, defPath string, targetFile, name string) interface{} {
	fullPath := defPath
	if !strings.HasPrefix(defPath, "/") {
		fullPath = filepath.Join(robi.Base, "defs", defPath)
	}
	tmpExcelFile := filepath.Join(robi.Base, "dicts", name+".xlsx")
	err := richframe.ExportRichFramesByTemp(data, targetFile, tmpExcelFile, fullPath, nil)
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
	monthStartDay := in.AddDate(0, 0, -in.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, in.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, in.Location())
	return monthEndTime
}

func (robi *Robi) SubTimeByHour(end, start time.Time) float64 {
	return end.Sub(start).Hours()
}

func Percent(x int, y int) string {
	z := (float64(x) / float64(y)) * 100
	return fmt.Sprintf("%4.2f", z)
}

func (robi *Robi) GetTemplatePath(tpl string) string {
	return filepath.Join(robi.Base, "tpl", tpl)
}

func (robi *Robi) ExeTemplate(tpl string, data interface{}) string {

	ms := sprig.FuncMap()
	ms["Percent"] = Percent

	tmp := template.Must(template.New(tpl).Funcs(ms).ParseFiles(filepath.Join(robi.Base, "tpl", tpl)))

	var w bytes.Buffer
	tmp.Execute(&w, data)

	return w.String()

}

func (robi *Robi) WriteFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	return err
}

func (robi *Robi) ExecuteCommand(cmdstr string, args []string) error {
	cmd := exec.Command(cmdstr, args...)

	output, err := cmd.CombinedOutput()
	fmt.Println(output)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New(string(output)) // not include err: executable file not found in %PATH%
	}
	return nil
}

// 判断一个时间是否在本月及高于本月
func (robi *Robi) IfInMonthAndHigherMonth(t time.Time) bool {
	now := time.Now()
	currentMonth := int(now.Month())
	inputMonth := int(t.Month())

	if inputMonth == currentMonth || inputMonth > currentMonth {
		return true
	} else {
		return false
	}
}

//判断两个时间大小

func (robi *Robi) CompareTwoTime(t1, t2 time.Time) bool {
	if t1.After(t2) {
		return true
	} else {
		return false
	}
}

// 判断两个时间相差几天 只看天

func (robi *Robi) GetDayDiff(date1, date2 time.Time) int {
	// 获取日期部分
	day1 := time.Date(date1.Year(), date1.Month(), date1.Day(), 0, 0, 0, 0, time.UTC)
	day2 := time.Date(date2.Year(), date2.Month(), date2.Day(), 0, 0, 0, 0, time.UTC)

	// 直接比较日期部分
	days := math.Abs(float64(day2.Unix()-day1.Unix()) / 86400)

	if date1.After(date2) {
		return -int(days)
	} else {
		return int(days)
	}
}
