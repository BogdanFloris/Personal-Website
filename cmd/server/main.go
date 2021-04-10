package main

import (
	"bogdanfloris-com/pkg/accessor"
	"bogdanfloris-com/pkg/logging"
)

func main() {
	logging.InitLoggers()
	acc := accessor.GetAccessor()
	defer acc.Close()

	acc.Test()
}
