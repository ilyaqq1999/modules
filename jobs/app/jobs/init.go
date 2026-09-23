package jobs

import (
	"github.com/ilyaqq1999/revel"
)

var jobLog = revel.AppLog

func init() {
	revel.RegisterModuleInit(func(m *revel.Module) {
		jobLog = m.Log
	})
}
