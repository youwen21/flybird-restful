package forms

type Condition struct {
	Raw          bool            `json:"raw" form:"raw"`
	RawCondition string          `json:"rawCondition" form:"rawCondition"`
	Items        []ConditionItem `json:"items" form:"items"`
}

type ConditionItem struct {
	Checked   bool   `json:"checked" form:"checked"`
	FiledName string `json:"filed_name" form:"filed_name"`
	Compare   string `json:"compare" form:"compare"`
	Bond      string `json:"bond" form:"bond"`
	Value     string `json:"value" form:"value"`
	Value2    string `json:"value2" form:"value2"`
	Index     int    `json:"index" form:"index"`
	Status    int    `json:"status" form:"status"`
}
