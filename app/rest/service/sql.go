package service

import (
	"context"
	"gofly/app/rest/forms"
	"gofly/app/rest/mysqlrest"
	"gofly/conf"
)

type sqlSrv struct{}

var SqlSrv = &sqlSrv{}

func (s *sqlSrv) Query(form *forms.QueryForm) (*mysqlrest.QueryResult, error) {
	gormDb := conf.Config.Mysql.GetSession()
	res, err := mysqlrest.DAO.Query(context.Background(), gormDb, form)

	return &res, err
}

func (s *sqlSrv) SqlRawQuery(form *forms.SqlQueryForm) (*mysqlrest.RawQueryResult, error) {
	//defer func() {
	//	_, _ = tools.HistorySrv.Record(form.Sql)
	//}()

	gormDb := conf.Config.Mysql.GetSession()

	res, err := mysqlrest.DAO.SqlRawQuery(context.Background(), gormDb, form)
	return &res, err
}

func (s *sqlSrv) Get(form *forms.GetForm) (map[string]interface{}, error) {
	gormDb := conf.Config.Mysql.GetSession()
	return mysqlrest.DAO.Get(context.Background(), gormDb, form)
}

func (s *sqlSrv) Insert(form *forms.PutForm) (int64, error) {
	gormDb := conf.Config.Mysql.GetSession()
	return mysqlrest.DAO.Insert(context.Background(), gormDb, form)
}

func (s *sqlSrv) Update(form *forms.RestUpdateForm) (int64, error) {
	gormDb := conf.Config.Mysql.GetSession()
	return mysqlrest.DAO.Update(context.Background(), gormDb, form)
}

func (s *sqlSrv) Delete(form *forms.GetForm) (int64, error) {
	gormDb := conf.Config.Mysql.GetSession()
	return mysqlrest.DAO.Delete(context.Background(), gormDb, form)
}

func (s *sqlSrv) Execute(form *forms.SqlExecuteForm) error {
	//defer func() {
	//	_, _ = tools.HistorySrv.Record(form.Sql)
	//}()

	gormDb := conf.Config.Mysql.GetSession()

	err := mysqlrest.DAO.Execute(gormDb, form.Sql)
	return err
}
