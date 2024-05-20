package forms

type PageForm struct {
	OrderBy     string `form:"orderBy" json:"orderBy" `        // 字段
	OrderDirect string `form:"orderDirect" json:"orderDirect"` // 方式
	Page        int    `form:"page,default=1" json:"page,default=1"`
	Psize       int    `form:"psize,default=100" json:"psize,default=100"`
}

//Limit 获取每页记录数
func (p *PageForm) Limit() int {
	return p.Psize
}

//Offset 获取当前页开始ID
func (p *PageForm) Offset() int {
	if p.Page == 0 {
		return 0
	}

	return (p.Page - 1) * p.Psize
}
