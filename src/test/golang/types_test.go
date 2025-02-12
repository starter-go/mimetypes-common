package main

import (
	"testing"

	"github.com/starter-go/mimetypes-common/src/test/golang/unit"
	"github.com/starter-go/units"
)

func TestFindType(t *testing.T) {
	args := []string{}
	mod := getModuleT()
	cfg := &units.Config{
		Args:       args,
		Cases:      unit.NameFindType,
		Module:     mod,
		Properties: nil,
		T:          t,
		UsePanic:   false,
	}
	units.Run(cfg)
}

func TestListTypes(t *testing.T) {
	args := []string{}
	mod := getModuleT()
	cfg := &units.Config{
		Args:       args,
		Cases:      unit.NameListTypes,
		Module:     mod,
		Properties: nil,
		T:          t,
		UsePanic:   false,
	}
	units.Run(cfg)
}
