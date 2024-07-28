package mimetypescommon

import (
	"github.com/starter-go/application"
	mimetypescommon "github.com/starter-go/mimetypes-common"
	"github.com/starter-go/mimetypes-common/gen/main4mimetypescommon"
	"github.com/starter-go/mimetypes-common/gen/test4mimetypescommon"
	"github.com/starter-go/mimetypes/modules/mimetypes"
	"github.com/starter-go/units/modules/units"
)

// Module  ...
func Module() application.Module {
	mb := mimetypescommon.NewMainModule()
	mb.Components(main4mimetypescommon.ExportComponents)

	mb.Depend(mimetypes.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := mimetypescommon.NewTestModule()
	mb.Components(test4mimetypescommon.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}
