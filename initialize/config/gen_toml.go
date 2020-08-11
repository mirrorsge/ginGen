package config

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func genToml() {
	const tomlTmp = `
	[app]
	  RuntimeRootPath = "runtime/"
	  #MB
	  ImageMaxSize = 5
	  ImageAllowExts = [".jpg",".jpeg",".png"]
	
	  LogSavePath = "logs/"
	  LogSaveName = "log"
	  LogFileExt = "log"
	  TimeFormat = "20060102"
	
	  [http]
	  Port = 8000
	  Mode = "debug"
	  ReadTimeout = 60
	  WriteTimeout = 60
	  PageSize = 10
	  CookieDomain = "****.cn"
	  #过期时间 单位:秒 ,86400 = 60 * 60 * 24 一天
	  TokenMaxAge = 86400
	
	
	  [database]
	  Type = "postgres"
	  DSN = "host=localhost port=5432 user=gin dbname=gin sslmode=disable password=12345678"
	  TablePrefix = "db_"
	
	  [redis]
	  Host = "127.0.0.1:6379"
	  Password = ""
	  MaxIdle = 30
	  MaxActive = 30
	  IdleTimeout = 200
	`
	tomlTmpl, err := template.New("").Parse(tomlTmp)
	if err != nil {
		panic("gen tmpl failed," + err.Error())
	}
	buff := bytes.NewBufferString("")
	err = tomlTmpl.Execute(buff, "")
	//进行格式化
	err = ioutil.WriteFile("./config/config.toml", buff.Bytes(), 0666)
	if err != nil {
		panic("gen config/config.toml failed," + err.Error())
	}
}
