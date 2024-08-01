package main

import (
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/mimetypes-common/modules/mimetypescommon"
	"github.com/starter-go/starter"
)

func main() {
	m := mimetypescommon.ModuleForTest()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}

func getModuleT() application.Module {
	return mimetypescommon.ModuleForTest()
}
