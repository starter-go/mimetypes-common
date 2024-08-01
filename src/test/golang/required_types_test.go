package main

import (
	"testing"

	"github.com/starter-go/mimetypes-common/src/test/golang/unit"
	"github.com/starter-go/units"
)

func TestRequiredSuffixList(t *testing.T) {
	args := []string{}
	mod := getModuleT()
	cfg := &units.Config{
		Args:       args,
		Cases:      unit.NameRequiredSuffixList,
		Module:     mod,
		Properties: nil,
		T:          t,
		UsePanic:   false,
	}
	units.Run(cfg)
}

func TestRequiredTypeList(t *testing.T) {
	args := []string{}
	mod := getModuleT()
	cfg := &units.Config{
		Args:       args,
		Cases:      unit.NameRequiredTypeList,
		Module:     mod,
		Properties: nil,
		T:          t,
		UsePanic:   false,
	}
	units.Run(cfg)
}
