package base

import (
	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	//定义日志格式
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	formatter.ForceFormatting = true
	formatter.SetColorScheme(&prefixed.ColorScheme{
		DebugLevelStyle: "yellow",
		WarnLevelStyle:  "orange",
		TimestampStyle:  "10",
	})
	log.SetFormatter(formatter)

	//日志级别
	//level := os.Getenv("log.debug")
	//if level == "false" {
	log.SetLevel(log.DebugLevel)
	//}
	//控制台高亮显示
	formatter.ForceColors = true
	//formatter.DisableColors = false
	//日志文件和滚动配置
	//log.Info("测试")
	log.Debug("测试")
	log.Info("-----------测试------------")
}
