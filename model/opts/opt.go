package opts

import (
	"github.com/easyhutu/go-echarts-option/model/opts/axis"
	"github.com/easyhutu/go-echarts-option/model/opts/grid"
	"github.com/easyhutu/go-echarts-option/model/opts/legend"
	"github.com/easyhutu/go-echarts-option/model/opts/titles"
)

type (
	Title  = titles.Title
	Legend = legend.Legend
	Grid   = grid.Grid
	Axis   = axis.Axis
)

type Type string

const (
	Line          = Type("line")
	Bar           = Type("bar")
	Pie           = Type("pie")
	Scatter       = Type("scatter")
	EffectScatter = Type("effectScatter")
	Radar         = Type("radar")
	Tree          = Type("tree")
	TreeMap       = Type("treemap")
	SunBurst      = Type("sunburst")
	Boxplot       = Type("boxplot")
	Candlestick   = Type("candlestick")
	Heatmap       = Type("heatmap")
	Map           = Type("map")
	Parallel      = Type("parallel")
	Lines         = Type("lines")
	Graph         = Type("graph")
	Sankey        = Type("sankey")
	Funnel        = Type("funnel")
	Gauge         = Type("gauge")
	PictorialBar  = Type("pictorialBar")
	ThemeRiver    = Type("themeRiver")
	Custom        = Type("custom")
)
