package infra

import (
	"github.com/tietang/props/kvs"
)

//应用程序启动管理器
type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

//通过New方法获取 应用程序启动管理器
func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{conf: conf, starterContext: StarterContext{}}
	b.starterContext[KeyProps] = conf
	return b
}

//应用启动方法
func (b *BootApplication) Start() {
	//1. 初始化starter
	b.init()
	//2. 安装starter
	b.setup()
	//3. 启动starter
	b.start()
}

//初始化方法
func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}
func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {

		if starter.StartBlocking() {
			//如果是最后一个可阻塞的，直接启动并阻塞
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else {
				//如果不是，使用goroutine来异步启动，
				// 防止阻塞后面starter
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
