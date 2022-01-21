package app

import (
	"fmt"
	"github.com/xh-dev-go/xhUtils/pathConfig"
	"io/ioutil"
	"os"
	"strconv"
)

type Info struct {
	Pid int
	Hostname string
}

var appInfo Info

func init() {
	if hostname, err := os.Hostname(); err != nil {
		panic(err)
	} else {
		appInfo = Info{
			Pid:      os.Getpid(),
			Hostname: hostname,
		}
	}
}

func GetApInfo() Info {
	return appInfo
}

type App struct{
	AppName string;
	CacheDir string;
	HomeDir string;
	ConfigDir string;
}

func NewApp(name string) App {
	app := App{AppName: name}
	app.CacheDir = app.getCacheDir()
	app.HomeDir = app.getHomeDir()
	app.ConfigDir = app.getConfigDir()
	if _, err := os.Stat(app.HomeDir); os.IsNotExist(err){
		os.Mkdir(app.HomeDir, os.ModeDir)
	}
	if _, err := os.Stat(app.CacheDir); os.IsNotExist(err){
		os.Mkdir(app.CacheDir, os.ModeDir)
	}
	if _, err := os.Stat(app.ConfigDir); os.IsNotExist(err){
		os.Mkdir(app.ConfigDir, os.ModeDir)
	}

	return app
}

func (customApp App) getConfigDir() string {
	return pathConfig.ConfigDir+pathConfig.SysSeparator+ customApp.AppName
}

func (customApp App) getHomeDir() string {
	return pathConfig.HomeDir+pathConfig.SysSeparator+".custApp"+pathConfig.SysSeparator+ customApp.AppName
}

func (customApp App) getCacheDir() string {
	return pathConfig.CacheDir+ pathConfig.SysSeparator + customApp.AppName
}

func (customApp App) GetKVPath(name string) string{
	return fmt.Sprintf("%s%s%s.kv",customApp.HomeDir,pathConfig.SysSeparator,name)
}

func (customApp App) GetKV(name string) (string,error) {
	var result string
	dir := customApp.GetKVPath(name)
	if _,err := os.Stat(dir); err != nil {
		return result, err
	} else if bs, err := ioutil.ReadFile(dir); err != nil {
		return result, err
	} else {
		result = string(bs)
		return result, nil
	}
}
func (customApp App) GetIntKV(name string) (int,error) {
	var result int
	if v, err := customApp.GetKV(name); err != nil {
		return result, err
	} else{
		return strconv.Atoi(v)
	}
}

func (customApp App) SetKV(name string, v string) error {
	file := customApp.GetKVPath(name)
	return ioutil.WriteFile(file, []byte(v), 0777)
}
func (customApp App) SetIntKV(name string, v int) error {
	return customApp.SetKV(name, strconv.Itoa(v))
}

func (customApp App) Info(){
	println(customApp.HomeDir)
	println(customApp.CacheDir)
	println(customApp.ConfigDir)
	println(customApp.AppName)
}
