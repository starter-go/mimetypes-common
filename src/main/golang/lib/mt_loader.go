package lib

import (
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/vlog"
)

// CommonMediaTypesLoader ...
type CommonMediaTypesLoader struct {

	//starter:component

	_as func(mimetypes.Registry) //starter:as(".")

	AppContext application.Context //starter:inject("context")
}

func (inst *CommonMediaTypesLoader) _impl() mimetypes.Registry { return inst }

func (inst *CommonMediaTypesLoader) listCSVFilePaths() []string {
	const (
		base = "/media-types/common/"
	)
	list := make([]string, 0)
	list = append(list, base+"application.csv")
	list = append(list, base+"audio.csv")
	list = append(list, base+"font.csv")
	list = append(list, base+"haptics.csv")
	list = append(list, base+"image.csv")
	list = append(list, base+"message.csv")
	list = append(list, base+"model.csv")
	list = append(list, base+"multipart.csv")
	list = append(list, base+"text.csv")
	list = append(list, base+"video.csv")
	list = append(list, base+"other.csv")
	return list
}

// ListRegistrations ...
func (inst *CommonMediaTypesLoader) ListRegistrations() []*mimetypes.Registration {

	builder := &innerTypeListBuilder{loader: inst}
	builder.init()
	src := inst.listCSVFilePaths()

	for _, path := range src {
		err := inst.loadFromResFile(path, builder)
		if err != nil {
			vlog.Warn(err.Error())
		}
	}

	return builder.list
}

func (inst *CommonMediaTypesLoader) loadFromResFile(path string, builder *innerTypeListBuilder) error {
	ac := inst.AppContext
	text, err := ac.GetResources().ReadText(path)
	if err != nil {
		return err
	}
	rows := inst.parseRows(text)
	for i, row := range rows {
		err := builder.parseRow(i, row)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *CommonMediaTypesLoader) parseRows(text string) []string {
	const (
		ch1 = "\r"
		ch2 = "\n"
	)
	txt2 := strings.ReplaceAll(text, ch1, ch2)
	return strings.Split(txt2, ch2)
}

////////////////////////////////////////////////////////////////////////////////

type innerTypeListBuilder struct {
	loader *CommonMediaTypesLoader
	list   []*mimetypes.Registration
}

func (inst *innerTypeListBuilder) init() {
	inst.list = make([]*mimetypes.Registration, 0)
}

func (inst *innerTypeListBuilder) parseRow(iRow int, text string) error {

	text = strings.TrimSpace(text)

	if text == "" {
		return nil // 跳过空行
	}
	if strings.HasPrefix(text, "#") {
		return nil // 跳过注释
	}

	const minCountCells = 3
	cells := strings.Split(text, ",")
	if len(cells) < minCountCells {
		return nil // 跳过列数不够的行
	}

	// 对每个数据单元整形
	for i, cell := range cells {
		cell = strings.TrimSpace(cell)
		cell = strings.ToLower(cell)
		cells[i] = cell
	}

	// format: Name,Template,Reference
	a := cells[0]
	b := cells[1]
	c := cells[2]
	s := cells[len(cells)-1] // 最后为后缀
	if a == "name" && b == "template" && c == "reference" {
		return nil // 跳过 header
	}

	return inst.handleRow(a, b, s)
}

func (inst *innerTypeListBuilder) handleRow(name, template, suffixes string) error {

	item := &mimetypes.Registration{
		Name:     name,
		Suffixes: inst.parseSuffixes(suffixes),
		Priority: 99,
	}

	info := &item.Info
	info.Icon = "/image/media/types/" + name + ".png"
	info.Label = "{{strings.media.types." + name + ".label}}"
	info.Description = "{{strings.media.types." + name + ".description}}"
	info.Type = mimetypes.Type(template)

	inst.list = append(inst.list, item)
	return nil
}

func (inst *innerTypeListBuilder) parseSuffixes(text string) []string {
	result := []string{}
	if strings.HasPrefix(text, "(") && strings.HasSuffix(text, ")") {
		i1 := 1
		i2 := len(text) - 1
		text = text[i1:i2]
	} else {
		return result
	}
	list := strings.Split(text, ".")
	for _, suffix := range list {
		suffix = strings.TrimSpace(suffix)
		suffix = strings.ToLower(suffix)
		if suffix == "" {
			continue
		}
		result = append(result, "."+suffix)
	}
	return result
}
