package test4mimetypescommon

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p749c9040ec_demotypes_DemoTypes{})
    inst.register(&pccc8cc1735_unit_DemoUnit{})
    inst.register(&pccc8cc1735_unit_FindTypeUnit{})
    inst.register(&pccc8cc1735_unit_ListTypesUnit{})
    inst.register(&pccc8cc1735_unit_RequiredTypesUnit{})


    return nil
}
