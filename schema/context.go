package schema

import (
	"github.com/easyhutu/go-echarts-option/model/opts"
)

type Context interface {
	XData() []interface{}
	YData() []interface{}
	Extend() interface{}
	Type() opts.Type
}
