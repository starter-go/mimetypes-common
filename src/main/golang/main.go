package main

import (
	"os"

	"github.com/starter-go/mimetypes-common/modules/mimetypescommon"

	"github.com/starter-go/starter"
)

func main() {
	m := mimetypescommon.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
