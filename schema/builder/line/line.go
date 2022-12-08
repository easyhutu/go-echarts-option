package line

import (
	"github.com/easyhutu/go-echarts-option/model/charts"
	"github.com/easyhutu/go-echarts-option/schema"
	"github.com/easyhutu/go-echarts-option/schema/builder"
)

type Builder interface {
	SetXData(ctx schema.Context) Builder
	SetYData(ctx schema.Context) Builder
	SetBase(base builder.Base) Builder
	Build() (charts.Line, error)
}

type line struct {
	ctx   schema.Context
	chart *charts.Chart
}

func (l line) SetBase(base builder.Base) Builder {
	panic("implement me")
}

func (l line) SetXData(ctx schema.Context) Builder {
	panic("implement me")
}

func (l line) SetYData(ctx schema.Context) Builder {
	panic("implement me")
}

func (l line) Build() (charts.Line, error) {
	panic("implement me")
}

func NewLineBuilder() Builder {
	return &line{}
}
