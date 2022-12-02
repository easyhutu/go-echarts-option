package axis

import (
	"github.com/easyhutu/go-echarts-option/model/opts/base"
	"github.com/easyhutu/go-echarts-option/model/types"
)

type Single struct {
	GridIndex     int            `json:"gridIndex,omitempty"`
	Position      types.Position `json:"position,omitempty"`
	Offset        int            `json:"offset,omitempty"`
	Type          Type           `json:"type,omitempty"`
	Name          string         `json:"name,omitempty"`
	NameLocation  types.Location `json:"nameLocation,omitempty"`
	NameTextStyle string         `json:"nameTextStyle,omitempty"`
	AxisLabel     Label          `json:"axisLabel,omitempty"`
	Data          []*Data        `json:"data,omitempty"`
	AxisPointer   Pointer        `json:"axisPointer,omitempty"`
}

func (s *Single) Update(singles ...Single) {
	if len(singles) == 0 {
		panic(singles)
	}
	single := singles[0]
	s.GridIndex = single.GridIndex
	s.Position = single.Position
	s.Offset = single.Offset
	s.Type = single.Type
	s.Name = single.Name
	s.NameLocation = single.NameLocation
	s.NameTextStyle = single.NameTextStyle
	s.AxisLabel = single.AxisLabel
	s.Data = single.Data
}

type Label struct {
	Interval  string `json:"interval,omitempty"`
	Inside    bool   `json:"inside,omitempty"`
	Rotate    int    `json:"rotate,omitempty"`
	Margin    int    `json:"margin,omitempty"`
	Formatter string `json:"formatter,omitempty"`
}

type Data struct {
	Value     string         `json:"value,omitempty"`
	TextStyle base.TextStyle `json:"textStyle,omitempty"`
}

type Pointer struct {
	Type           PointerType `json:"type,omitempty"`
	Snap           bool        `json:"snap,omitempty"`
	Z              int         `json:"z,omitempty"`
	Label          Label       `json:"label,omitempty"`
	Value          int         `json:"value,omitempty"`
	TriggerTooltip bool        `json:"triggerTooltip,omitempty"`
	LineStyle      LineStyle   `json:"lineStyle,omitempty"`
}

type LineStyle struct {
	Color types.Color `json:"color,omitempty"`
	Width int         `json:"width,omitempty"`
}
