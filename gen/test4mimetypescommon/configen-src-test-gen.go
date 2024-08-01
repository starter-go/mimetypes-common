package test4mimetypescommon
import (
    p85a4d026d "github.com/starter-go/mimetypes"
    pccc8cc173 "github.com/starter-go/mimetypes-common/src/test/golang/unit"
    p749c9040e "github.com/starter-go/mimetypes-common/src/test/golang/unit/demotypes"
     "github.com/starter-go/application"
)

// type pccc8cc173.DemoUnit in package:github.com/starter-go/mimetypes-common/src/test/golang/unit
//
// id:com-ccc8cc1735d8ba89-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pccc8cc1735_unit_DemoUnit struct {
}

func (inst* pccc8cc1735_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccc8cc1735d8ba89-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccc8cc1735_unit_DemoUnit) new() any {
    return &pccc8cc173.DemoUnit{}
}

func (inst* pccc8cc1735_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccc8cc173.DemoUnit)
	nop(ie, com)

	


    return nil
}



// type pccc8cc173.FindTypeUnit in package:github.com/starter-go/mimetypes-common/src/test/golang/unit
//
// id:com-ccc8cc1735d8ba89-unit-FindTypeUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pccc8cc1735_unit_FindTypeUnit struct {
}

func (inst* pccc8cc1735_unit_FindTypeUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccc8cc1735d8ba89-unit-FindTypeUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccc8cc1735_unit_FindTypeUnit) new() any {
    return &pccc8cc173.FindTypeUnit{}
}

func (inst* pccc8cc1735_unit_FindTypeUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccc8cc173.FindTypeUnit)
	nop(ie, com)

	
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*pccc8cc1735_unit_FindTypeUnit) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}



// type pccc8cc173.ListTypesUnit in package:github.com/starter-go/mimetypes-common/src/test/golang/unit
//
// id:com-ccc8cc1735d8ba89-unit-ListTypesUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pccc8cc1735_unit_ListTypesUnit struct {
}

func (inst* pccc8cc1735_unit_ListTypesUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccc8cc1735d8ba89-unit-ListTypesUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccc8cc1735_unit_ListTypesUnit) new() any {
    return &pccc8cc173.ListTypesUnit{}
}

func (inst* pccc8cc1735_unit_ListTypesUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccc8cc173.ListTypesUnit)
	nop(ie, com)

	
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*pccc8cc1735_unit_ListTypesUnit) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}



// type pccc8cc173.RequiredTypesUnit in package:github.com/starter-go/mimetypes-common/src/test/golang/unit
//
// id:com-ccc8cc1735d8ba89-unit-RequiredTypesUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pccc8cc1735_unit_RequiredTypesUnit struct {
}

func (inst* pccc8cc1735_unit_RequiredTypesUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccc8cc1735d8ba89-unit-RequiredTypesUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccc8cc1735_unit_RequiredTypesUnit) new() any {
    return &pccc8cc173.RequiredTypesUnit{}
}

func (inst* pccc8cc1735_unit_RequiredTypesUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccc8cc173.RequiredTypesUnit)
	nop(ie, com)

	
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*pccc8cc1735_unit_RequiredTypesUnit) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}



// type p749c9040e.DemoTypes in package:github.com/starter-go/mimetypes-common/src/test/golang/unit/demotypes
//
// id:com-749c9040ec595c54-demotypes-DemoTypes
// class:class-85a4d026daf77828ef49edb2adfd695e-Registry
// alias:
// scope:singleton
//
type p749c9040ec_demotypes_DemoTypes struct {
}

func (inst* p749c9040ec_demotypes_DemoTypes) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-749c9040ec595c54-demotypes-DemoTypes"
	r.Classes = "class-85a4d026daf77828ef49edb2adfd695e-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p749c9040ec_demotypes_DemoTypes) new() any {
    return &p749c9040e.DemoTypes{}
}

func (inst* p749c9040ec_demotypes_DemoTypes) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p749c9040e.DemoTypes)
	nop(ie, com)

	


    return nil
}


