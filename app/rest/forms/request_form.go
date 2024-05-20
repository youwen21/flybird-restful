package forms

type UpdateCondition struct {
	HasPRI         bool                   `json:"hasPRI" form:"hasPRI"`
	ConditionItems map[string]interface{} `json:"conditionItems" form:"conditionItems"`
}

type UpdateForm struct {
	ConnAbsDTO
	UpdateCondition UpdateCondition        `json:"updateCondition" form:"updateCondition"`
	Params          map[string]interface{} `json:"params" form:"params"`
}

type PutForm struct {
	ConnAbsDTO
	Params map[string]interface{} `json:"params" form:"params"`
}

type GetForm struct {
	ConnAbsDTO
	Id int `json:"id"`
}

type QueryForm struct {
	PageForm

	ConnAbsDTO

	Params    map[string]interface{} `json:"params" form:"params"`
	Condition `json:"condition" form:"condition"`

	ExportType string `json:"exportType" form:"exportType"`
}

type SqlQueryForm struct {
	PageForm

	ConnAbsDTO

	Sql        string `json:"sql" form:"sql"`
	ExportType string `json:"exportType" form:"exportType"`
}

type SqlExecuteForm struct {
	ConnAbsDTO

	Sql string `json:"sql" form:"sql"`
}

type AnyForm struct {
	PageForm

	ConnAbsDTO

	RestMethod string `json:"x_rest_method" form:"x_rest_method"`

	Id int `json:"id" form:"id"`

	Params map[string]interface{} `json:"params" form:"params"`
}

type RestUpdateForm struct {
	ConnAbsDTO
	Id     int                    `json:"id"`
	Params map[string]interface{} `json:"params" form:"params"`
}
