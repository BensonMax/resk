package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"time"
)

func main() {
	file := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileConfigSource(file)
	port := conf.GetIntDefault("app.server.port", 18080)
	fmt.Println(port)
	fmt.Println(conf.GetDefault("app.name", "unknow"))
	fmt.Println(conf.GetBoolDefault("app.enabled", false))
	fmt.Println(conf.GetDurationDefault("app.time", time.Second))

}
