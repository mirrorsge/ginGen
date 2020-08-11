package config

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func genConf() {
	const Tmp = `
	package config

	import (
		"log"
		"time"
	
		"github.com/BurntSushi/toml"
	)
	
	type Config struct {
		App      App
		Http     Http
		Database Database
		Redis    Redis
	}
	
	
	type App struct {
		RuntimeRootPath string
	
		ImagePrefixUrl string
		ImageSavePath  string
		ImageMaxSize   int
		ImageAllowExts []string
	
		LogSavePath string
		LogSaveName string
		LogFileExt  string
		TimeFormat  string
	}
	
	var AppSetting = &App{}
	
	type Http struct {
		Mode         string
		Port         int
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		PageSize     int
		CookieDomain string
		TokenMaxAge  int64
	}
	
	var HttpSetting = &Http{}
	
	type Database struct {
		Type        string
		DSN         string
		TablePrefix string
	}
	
	var DatabaseSetting = &Database{}
	
	type Redis struct {
		Host        string
		Password    string
		MaxIdle     int
		MaxActive   int
		IdleTimeout time.Duration
	}
	
	var RedisSetting = &Redis{}
	
	
	func Setup() {
		var config Config
		filePath := "config/config.toml"
		if _, err := toml.DecodeFile(filePath, &config); err != nil {
			log.Fatalf("Fail to parse 'config/config.toml': %v", err)
		}
	
		//Init app settings
		AppSetting = &config.App
	
		AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	
		//Init server settings
		HttpSetting = &config.Http
	
		HttpSetting.ReadTimeout = HttpSetting.ReadTimeout * time.Second
		HttpSetting.WriteTimeout = HttpSetting.ReadTimeout * time.Second
	
		//Init database settings
		DatabaseSetting = &config.Database
	
		//Init redis settings
		RedisSetting = &config.Redis
	}
	`
	tomlTmpl, err := template.New("").Parse(Tmp)
	if err != nil {
		panic("gen config_tmpl failed," + err.Error())
	}
	buff := bytes.NewBufferString("")
	err = tomlTmpl.Execute(buff, "")
	//进行格式化
	err = ioutil.WriteFile("./config/config.go", buff.Bytes(), 0666)
	if err != nil {
		panic("gen config/config.go failed," + err.Error())
	}
}
