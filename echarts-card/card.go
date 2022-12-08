package echarts_card

import (
	"github.com/easyhutu/go-echarts-option/model/opts"
	"github.com/easyhutu/go-echarts-option/schema"
)

type context struct {
	xData  []interface{}
	yData  []interface{}
	eType  opts.Type
	extend interface{}
}

func (c context) XData() []interface{} {
	return c.xData
}

func (c context) YData() []interface{} {
	return c.yData
}

func (c context) Extend() interface{} {
	return c.extend
}

func (c context) Type() opts.Type {
	return c.eType
}

func NewContext(tp opts.Type, x, y []interface{}, extend interface{}) schema.Context {
	return &context{x, y, tp, extend}
}
