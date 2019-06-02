package resk

import (
	"github.com/resk/infra"
	"github.com/resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	//infra.DemoRegister(&base.PropsStarter{})
}
