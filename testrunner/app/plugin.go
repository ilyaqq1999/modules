package app

import (
	"github.com/ilyaqq1999/revel"
)

func init() {
	revel.OnAppStart(func() {
		revel.AppLog.Info("Go to /@tests to run the tests.")
	})
}
