package base

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisrecover "github.com/kataras/iris/middleware/recover"
	"github.com/resk/infra"
	"github.com/sirupsen/logrus"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisServerStarter struct {
	infra.BaseStarter
}

func (i *IrisServerStarter) Init(ctx infra.StarterContext) {
	//创建iris application实例
	irisApplication = initIris()
	//日志组件的配置和扩展
	logger := irisApplication.Logger()
	logger.Install(logrus.StandardLogger())

}

func (i *IrisServerStarter) Start(ctx infra.StarterContext) {
	//把路由信息打印到控制台
	routes := Iris().GetRoutes()
	for _, r := range routes {
		logrus.Info(r.Trace())
	}
	//启动Iris
	port := ctx.Props().GetDefault("app.server.port", "18080")
	Iris().Run(iris.Addr(":" + port))
}

func (i *IrisServerStarter) StartBlocking() bool {
	return true
}

func initIris() *iris.Application {
	app := iris.New()
	app.Use(irisrecover.New())
	//主要中间件的配置： recover，日志输出中间件的自定义
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s | %s |",
				now.Format("2006-01-02.15:04:05.000000"),
				latency.String(), status, ip, method, path, message, headerMessage,
			)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
