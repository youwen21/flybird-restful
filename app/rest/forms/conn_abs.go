package forms

import (
	"errors"
)

type ConnAbsDTO struct {
	Ident    string `json:"ident" form:"ident" zh:"连接名hash ID"`
	Name     string `json:"name" form:"name" zh:"连接名"`
	Database string `json:"database" form:"database" zh:"打开的数据库"`
	Table    string `json:"table" form:"table" zh:"操作的数据表"`
}

func (c *ConnAbsDTO) CheckValid() error {
	if c.Name == "" {
		return errors.New("connection empty")
	}

	return nil
}
