package custApp

import "github.com/xh-dev-go/xhUtils/pathConfig"
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
	return pathConfig.ConfigDir+string(pathConfig.SysSeparator)+ customApp.AppName
}

func (customApp CustomApp) getHomeDir() string {
	return pathConfig.HomeDir+string(pathConfig.SysSeparator)+"."+ customApp.AppName
}

func (customApp CustomApp) getCacheDir() string {
	return pathConfig.CacheDir+string(pathConfig.SysSeparator)+ customApp.AppName
}

func (customApp CustomApp) Info(){
	println(customApp.HomeDir)
	println(customApp.CacheDir)
	println(customApp.ConfigDir)
	println(customApp.AppName)
}
