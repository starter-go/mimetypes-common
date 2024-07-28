package main4mimetypescommon
import (
    p0ef6f2938 "github.com/starter-go/application"
    pf34fc1962 "github.com/starter-go/mimetypes-common/src/main/golang/lib"
     "github.com/starter-go/application"
)

// type pf34fc1962.CommonMediaTypesLoader in package:github.com/starter-go/mimetypes-common/src/main/golang/lib
//
// id:com-f34fc19624b2155d-lib-CommonMediaTypesLoader
// class:class-85a4d026daf77828ef49edb2adfd695e-Registry
// alias:
// scope:singleton
//
type pf34fc19624_lib_CommonMediaTypesLoader struct {
}

func (inst* pf34fc19624_lib_CommonMediaTypesLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f34fc19624b2155d-lib-CommonMediaTypesLoader"
	r.Classes = "class-85a4d026daf77828ef49edb2adfd695e-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf34fc19624_lib_CommonMediaTypesLoader) new() any {
    return &pf34fc1962.CommonMediaTypesLoader{}
}

func (inst* pf34fc19624_lib_CommonMediaTypesLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf34fc1962.CommonMediaTypesLoader)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)


    return nil
}


func (inst*pf34fc19624_lib_CommonMediaTypesLoader) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


