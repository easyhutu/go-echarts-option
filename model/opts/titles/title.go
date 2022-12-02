package titles

const (
	TargetSelf  = TargetType("self")
	TargetBlank = TargetType("blank")
)

type TargetType string

// Title 标题组件，包含主标题和副标题。https://echarts.apache.org/zh/option.html#title
type Title struct {
	Show      bool       `json:"show,omitempty"`
	Text      string     `json:"text,omitempty"`
	Link      string     `json:"link,omitempty"`
	Target    TargetType `json:"target,omitempty"`
	Subtext   string     `json:"subtext,omitempty"`
	Sublink   string     `json:"sublink,omitempty"`
	Subtarget TargetType `json:"subtarget,omitempty"`
}
