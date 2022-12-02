package axis

const (
	Value    Type        = "value"
	Category Type        = "category"
	Time     Type        = "time"
	Log      Type        = "log"
	Line     PointerType = "line"
	Shadow   PointerType = "shadow"
	None     PointerType = "none"
)

type Type string
type PointerType string

// Axis 直角坐标系 grid 中的 x,y 轴，一般情况下单个 grid 组件最多只能放上下两个 x 轴 https://echarts.apache.org/zh/option.html#xAxis
type Axis interface {
	Update(singles ...Single)
}
