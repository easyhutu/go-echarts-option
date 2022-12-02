package grid

import (
	"github.com/easyhutu/go-echarts-option/model/opts/base"
	"github.com/easyhutu/go-echarts-option/model/types"
)

// Grid 直角坐标系内绘图网格，单个 grid 内最多可以放置上下两个 X 轴，左右两个 Y 轴 https://echarts.apache.org/zh/option.html#grid
type Grid struct {
	base.Sizes
	ContainLabel    bool        `json:"containLabel,omitempty"`
	BackgroundColor types.Color `json:"backgroundColor,omitempty"`
	BorderColor     types.Color `json:"borderColor,omitempty"`
}
