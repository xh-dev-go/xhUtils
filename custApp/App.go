package custApp

import (
	"fmt"
	"github.com/xh-dev-go/xhUtils/pathConfig"
	"io/ioutil"
	"strconv"
)
import "os"

type CustomApp struct{
	AppName string;
	CacheDir string;
	HomeDir string;
	ConfigDir string;
}

func NewCustomApp(name string) CustomApp{
	app := CustomApp{AppName: name}
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

func (customApp CustomApp) getConfigDir() string {
	return pathConfig.ConfigDir+pathConfig.SysSeparator+ customApp.AppName
}

func (customApp CustomApp) getHomeDir() string {
	return pathConfig.HomeDir+pathConfig.SysSeparator+".custApp"+pathConfig.SysSeparator+ customApp.AppName
}

func (customApp CustomApp) getCacheDir() string {
	return pathConfig.CacheDir+string(pathConfig.SysSeparator)+ customApp.AppName
}

func (customApp CustomApp) GetKVPath(name string) string{
	return fmt.Sprintf("%s%s%s.kv",customApp.HomeDir,pathConfig.SysSeparator,name)
}

func (customApp CustomApp) GetKV(name string) (string,error) {
	var result string
	dir := customApp.GetKVPath(name)
	if bs, err := ioutil.ReadFile(dir); err != nil {
		return result, err
	} else {
		result = string(bs)
		return result, nil
	}
}
func (customApp CustomApp) GetIntKV(name string) (int,error) {
	var result int
	if v, err := customApp.GetKV(name); err != nil {
		return result, err
	} else{
		return strconv.Atoi(v)
	}
}

func (customApp CustomApp) SetKV(name string, v string) error {
	file := customApp.GetKVPath(name)
	return ioutil.WriteFile(file, []byte(v), 0777)
}
func (customApp CustomApp) SetIntKV(name string, v int) error {
	return customApp.SetKV(name, strconv.Itoa(v))
}

func (customApp CustomApp) Info(){
	println(customApp.HomeDir)
	println(customApp.CacheDir)
	println(customApp.ConfigDir)
	println(customApp.AppName)
}
