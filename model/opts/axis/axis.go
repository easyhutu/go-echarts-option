package axis

import (
	"github.com/easyhutu/go-echarts-option/model/types"
)

const (
	Value    Type = "value"
	Category Type = "category"
	Time     Type = "time"
	Log      Type = "log"
)

// Axis 直角坐标系 grid 中的 x,y 轴，一般情况下单个 grid 组件最多只能放上下两个 x 轴 https://echarts.apache.org/zh/option.html#xAxis
type Axis struct {
	GridIndex     int            `json:"gridIndex,omitempty"`
	Position      types.Position `json:"position,omitempty"`
	Offset        int            `json:"offset,omitempty"`
	Type          Type           `json:"type,omitempty"`
	Name          string         `json:"name,omitempty"`
	NameLocation  types.Location `json:"nameLocation,omitempty"`
	NameTextStyle string         `json:"nameTextStyle,omitempty"`
}

type Type string
