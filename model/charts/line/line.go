package line

import "github.com/easyhutu/go-echarts-option/model/opts"

type Line struct {
	Title opts.Title `json:"title,omitempty"`
	XAxis opts.Axis  `json:"xAxis,omitempty"`
	YAxis opts.Axis  `json:"yAxis,omitempty"`
}
