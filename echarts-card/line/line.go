package line

import (
	"github.com/easyhutu/go-echarts-option/model/charts"
	"github.com/easyhutu/go-echarts-option/schema"
	"github.com/easyhutu/go-echarts-option/schema/builder"
	line2 "github.com/easyhutu/go-echarts-option/schema/builder/line"
)

func NewLine(ctx schema.Context) charts.Line {
	base := builder.NewBaseBuilder().SetTooltip().SetTitle(ctx)
	line, _ := line2.NewLineBuilder().
		SetBase(base).
		SetXData(ctx).
		SetYData(ctx).Build()
	return line
}
