package unit

import (
	"github.com/starter-go/mimetypes"
	mimetypescommon "github.com/starter-go/mimetypes-common"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// RequiredTypesUnit ...
type RequiredTypesUnit struct {

	//starter:component

	_as func(units.Units) //starter:as(".")

	TM mimetypes.Manager //starter:inject("#")

}

func (inst *RequiredTypesUnit) _impl() units.Units {
	return inst
}

// Units ...
func (inst *RequiredTypesUnit) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     NameRequiredSuffixList,
		Enabled:  true,
		Priority: 0,
		Test:     inst.testSuffixList,
	})

	list = append(list, &units.Registration{
		Name:     NameRequiredTypeList,
		Enabled:  true,
		Priority: 0,
		Test:     inst.testTypeNameList,
	})

	return list
}

func (inst *RequiredTypesUnit) prepareOptions() *mimetypes.Options {
	return &mimetypes.Options{
		Language: "zh_cn",
	}
}

func (inst *RequiredTypesUnit) testSuffixList() error {
	opt := inst.prepareOptions()
	all := mimetypescommon.GetRequiredFileNameSuffixList()
	var lastErr error
	countErr := 0
	for _, suffix := range all {
		info, err := inst.TM.FindBySuffix(suffix, opt)
		if err != nil {
			lastErr = err
			vlog.Error(err.Error())
			continue
		}
		inst.logTypeInfo("testSuffixList: ", info)
	}
	if countErr > 0 {
		vlog.Warn("err.count=%d", countErr)
	}
	return lastErr
}

func (inst *RequiredTypesUnit) testTypeNameList() error {
	opt := inst.prepareOptions()
	var lastErr error
	countErr := 0
	all := mimetypescommon.GetRequiredTypeNameList()
	for _, tn := range all {
		info, err := inst.TM.Find(tn, opt)
		if err != nil {
			lastErr = err
			vlog.Error(err.Error())
			continue
		}
		inst.logTypeInfo("testTypeNameList: ", info)
	}
	if countErr > 0 {
		vlog.Warn("err.count=%d", countErr)
	}
	return lastErr
}

func (inst *RequiredTypesUnit) logTypeInfo(tag string, info *mimetypes.Info) {
	name := info.Type.Pure()
	label := info.Label
	desc := info.Description
	vlog.Info("[tag:'%s'  type:'%s'  label:'%s'  desc:'%s']", tag, name, label, desc)
}
