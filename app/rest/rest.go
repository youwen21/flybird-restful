package rest

import (
	"context"
	"gofly/app/rest/forms"
)

type DbRESTFul interface {
	Query(ctx context.Context, form *forms.QueryForm) ([]map[string]interface{}, error)
	Get(ctx context.Context, form *forms.GetForm) (map[string]interface{}, error)
	Insert(ctx context.Context, form *forms.PutForm) (int64, error)
	Update(ctx context.Context, form *forms.UpdateForm) (int64, error)
	Delete(ctx context.Context, form *forms.GetForm) (int64, error)
}

var (
	MysqlDbRest  DbRESTFul
	PgDbRest     DbRESTFul
	SqliteDbRest DbRESTFul
	MongoDbRest  DbRESTFul
)
