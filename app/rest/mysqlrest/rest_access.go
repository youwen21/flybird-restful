package mysqlrest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"gofly/app/rest/forms"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type dao struct{}

var (
	DAO = &dao{}

	// http://books.studygolang.com/gorm/advanced.html#sb
	//queryErr = errors.New("查询失败")
)

type QueryResult struct {
	Columns []string                 `json:"columns" from:"columns"`
	Count   int64                    `sql:"count" form:"count" json:"count"`
	Rows    []map[string]interface{} `sql:"rows" form:"rows" json:"rows"`
}

type RawQueryResult struct {
	Finished string                   `sql:"finished" form:"finished" json:"finished"`
	Columns  []string                 `json:"columns" from:"columns"`
	Count    int64                    `sql:"count" form:"count" json:"count"`
	Rows     []map[string]interface{} `sql:"rows" form:"rows" json:"rows"`
}

func (d *dao) Execute(gorm *gorm.DB, sql string) error {
	gorm.Exec(sql)
	return gorm.Error
}

func (d *dao) SqlRawQuery(ctx context.Context, gorm *gorm.DB, form *forms.SqlQueryForm) (RawQueryResult, error) {
	var count int64
	//db, _ := gorm.DB()
	db := gorm.Raw(form.Sql)

	//session.Limit(form.Limit()).Offset(form.Offset())
	rows, err := db.Rows()
	if nil != err {
		log.Println("查询失败:", err.Error())
		return RawQueryResult{}, err
	}
	defer func() {
		_ = rows.Close()
	}()

	columns, _ := rows.Columns()
	columnTypes, _ := rows.ColumnTypes()

	finished := "yes"
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		if int(count) > form.Limit() {
			finished = "no"
			count--
			break
		}

		cache := PrepareValues(columnTypes) // 放到循环外面会导致数据覆盖
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			columnScanType := columnTypes[i].ScanType().Name()
			item[columns[i]] = GetValue(columnScanType, data)
		}
		list = append(list, item)
		count++
	}
	return RawQueryResult{finished, columns, count, list}, nil
}

func (d *dao) Count(ctx context.Context, gorm *gorm.DB, form *forms.QueryForm) (int64, error) {
	session := gorm.Table(form.Table).Select("*")
	for k, v := range form.Params {
		strV := v.(string)
		if len(strV) > 20 {
			session.Where(k+"=?", v)
			continue
		}
		if intV, err := strconv.Atoi(strV); err == nil {
			session.Where(k+"=?", intV)
		} else {
			session.Where(k+"=?", v)
		}
	}
	var count int64
	session.Count(&count)
	return count, nil
}

func (d *dao) Query(ctx context.Context, gorm *gorm.DB, form *forms.QueryForm) (QueryResult, error) {
	if form.Condition.Raw {
		return d.QueryCondition(ctx, gorm, form)
	} else {
		return d.QueryForm(ctx, gorm, form)
	}
}

func (d *dao) QueryCondition(ctx context.Context, gorm *gorm.DB, form *forms.QueryForm) (QueryResult, error) {
	session := gorm.Table(form.Table).Select("*")

	session.Where(form.Condition.RawCondition)

	var count int64
	session.Count(&count)

	session.Limit(form.Limit()).Offset(form.Offset())
	rows, err := session.Rows()
	if nil != err {
		log.Println("查询失败:", err.Error())
	}
	defer func() {
		_ = rows.Close()
	}()

	columns, _ := rows.Columns()
	columnTypes, _ := rows.ColumnTypes()

	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		cache := PrepareValues(columnTypes) // 放到循环外面会导致数据覆盖
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			columnScanType := columnTypes[i].ScanType().Name()
			item[columns[i]] = GetValue(columnScanType, data)
		}
		list = append(list, item)
	}
	return QueryResult{columns, count, list}, nil
}

func (d *dao) QueryForm(ctx context.Context, gorm *gorm.DB, form *forms.QueryForm) (QueryResult, error) {
	session := gorm.Table(form.Table).Select("*")
	//for k, v := range form.Params {
	//	strV := v.(string)
	//	if len(strV) > 20 {
	//		session.Where(k+"=?", v)
	//		continue
	//	}
	//	if intV, err := strconv.Atoi(strV); err == nil {
	//		session.Where(k+"=?", intV)
	//	} else {
	//		session.Where(k+"=?", v)
	//	}
	//}
	for _, v := range form.Condition.Items {
		if v.Status != 1 {
			continue
		}
		if v.Checked != true {
			continue
		}

		if v.FiledName == "" {
			continue
		}

		switch v.Compare {
		case "is equal to":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s = ?", v.FiledName), v.Value)
			} else {
				session.Or(fmt.Sprintf("%s = ?", v.FiledName), v.Value)
			}
		case "is nor equal to":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s != ?", v.FiledName), v.Value)
			} else {
				session.Or(fmt.Sprintf("%s != ?", v.FiledName), v.Value)
			}
		case "is less then":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s < ?", v.FiledName), v.Value)
			} else {
				session.Or(fmt.Sprintf("%s < ?", v.FiledName), v.Value)
			}
		case "is less then or equal to":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s <= ?", v.FiledName), v.Value)
			} else {
				session.Or(fmt.Sprintf("%s <= ?", v.FiledName), v.Value)
			}
		case "is greater then":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s > ?", v.FiledName), v.Value)
			} else {
				session.Or(fmt.Sprintf("%s > ?", v.FiledName), v.Value)
			}
		case "is greater then or equal to":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s >= ?", v.FiledName), v.Value)
			} else {
				session.Or(fmt.Sprintf("%s >= ?", v.FiledName), v.Value)
			}
		case "contains":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s LIKE ?", v.FiledName), fmt.Sprintf("%%%v%%", v.Value))
			} else {
				session.Or(fmt.Sprintf("%s LIKE ?", v.FiledName), fmt.Sprintf("%%%v%%", v.Value))
			}
		case "does not contain":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s NOT LIKE ?", v.FiledName), fmt.Sprintf("%%%v%%", v.Value))
			} else {
				session.Or(fmt.Sprintf("%s NOT LIKE ?", v.FiledName), fmt.Sprintf("%%%v%%", v.Value))
			}
		case "begin with":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s LIKE ?", v.FiledName), fmt.Sprintf("%v%%", v.Value))
			} else {
				session.Or(fmt.Sprintf("%s LIKE ?", v.FiledName), fmt.Sprintf("%v%%", v.Value))
			}
		case "does not begin with":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s NOT LIKE ?", v.FiledName), fmt.Sprintf("%v%%", v.Value))
			} else {
				session.Or(fmt.Sprintf("%s NOT LIKE ?", v.FiledName), fmt.Sprintf("%v%%", v.Value))
			}
		case "end with":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s LIKE ?", v.FiledName), fmt.Sprintf("%%%v", v.Value))
			} else {
				session.Or(fmt.Sprintf("%s LIKE ?", v.FiledName), fmt.Sprintf("%%%v", v.Value))
			}
		case "does not end with":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s NOT LIKE ?", v.FiledName), fmt.Sprintf("%%%v", v.Value))
			} else {
				session.Or(fmt.Sprintf("%s NOT LIKE ?", v.FiledName), fmt.Sprintf("%%%v", v.Value))
			}
		case "is null":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s IS NULL", v.FiledName))
			} else {
				session.Or(fmt.Sprintf("%s IS NULL", v.FiledName))
			}
		case "is not null":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s IS NOT NULL", v.FiledName))
			} else {
				session.Or(fmt.Sprintf("%s IS NOT NULL", v.FiledName))
			}
		case "is empty":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s = ''", v.FiledName))
			} else {
				session.Or(fmt.Sprintf("%s = ''", v.FiledName))
			}
		case "is not empty":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s != ''", v.FiledName))
			} else {
				session.Or(fmt.Sprintf("%s != ''", v.FiledName))
			}
		case "is between":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s between ? and ?", v.FiledName), v.Value, v.Value2)
			} else {
				session.Where(fmt.Sprintf("%s between ? and ?", v.FiledName), v.Value, v.Value2)
			}
		case "is not between":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s not between ? and ?", v.FiledName), v.Value, v.Value2)
			} else {
				session.Where(fmt.Sprintf("%s not between ? and ?", v.FiledName), v.Value, v.Value2)
			}
		case "is in list":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s in (?)", v.FiledName), strings.Split(v.Value, ","))
			} else {
				session.Where(fmt.Sprintf("%s in (?)", v.FiledName), strings.Split(v.Value, ","))
			}
		case "is not in list":
			if v.Bond == "and" {
				session.Where(fmt.Sprintf("%s not in (?)", v.FiledName), strings.Split(v.Value, ","))
			} else {
				session.Where(fmt.Sprintf("%s not in (?)", v.FiledName), strings.Split(v.Value, ","))
			}
		}
	}

	var count int64
	session.Count(&count)

	session.Limit(form.Limit()).Offset(form.Offset())
	rows, err := session.Rows()
	if nil != err {
		log.Println("查询失败:", err.Error())
	}
	defer func() {
		_ = rows.Close()
	}()
	columns, _ := rows.Columns()
	columnTypes, _ := rows.ColumnTypes()

	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		cache := PrepareValues(columnTypes) // 放到循环外面会导致数据覆盖
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			columnScanType := columnTypes[i].ScanType().Name()
			item[columns[i]] = GetValue(columnScanType, data)
		}
		list = append(list, item)
	}
	return QueryResult{columns, count, list}, nil
}

func (d *dao) Get(ctx context.Context, gorm *gorm.DB, form *forms.GetForm) (map[string]interface{}, error) {
	session := gorm.Table(form.Table).Select("*")
	session.Where("id = ?", form.Id).Limit(1)

	rows, err := session.Rows()
	if nil != err {
		log.Println("查询失败:", err.Error())
	}
	defer func() {
		_ = rows.Close()
	}()
	ok := rows.Next()
	if !ok {
		return nil, errors.New("recorde not exist")
	}

	columns, _ := rows.Columns()
	columnTypes, _ := rows.ColumnTypes()
	cache := PrepareValues(columnTypes)
	err = rows.Scan(cache...)
	if err != nil {
		return nil, fmt.Errorf("recorde scan error:%v", err)
	}

	item := make(map[string]interface{})
	for i, data := range cache {
		columnScanType := columnTypes[i].ScanType().Name()
		item[columns[i]] = GetValue(columnScanType, data)
	}
	return item, nil
}

func (d *dao) GetByPk(ctx context.Context, gorm *gorm.DB, form *forms.GetForm, pkName string) (map[string]interface{}, error) {
	session := gorm.Table(form.Table).Select("*")
	session.Where(fmt.Sprintf("`%s` = ?", pkName), form.Id).Limit(1)

	rows, err := session.Rows()
	if nil != err {
		log.Println("查询失败:", err.Error())
	}
	defer func() {
		_ = rows.Close()
	}()
	ok := rows.Next()
	if !ok {
		return nil, errors.New("recorde not exist")
	}

	columns, _ := rows.Columns()
	columnTypes, _ := rows.ColumnTypes()
	cache := PrepareValues(columnTypes)
	err = rows.Scan(cache...)
	if err != nil {
		return nil, fmt.Errorf("recorde scan error:%v", err)
	}

	item := make(map[string]interface{})
	for i, data := range cache {
		columnScanType := columnTypes[i].ScanType().Name()
		item[columns[i]] = GetValue(columnScanType, data)
	}
	return item, nil
}

func (d *dao) Insert(ctx context.Context, gorm *gorm.DB, form *forms.PutForm) (int64, error) {
	session := gorm.Table(form.Table)
	session.Create(form.Params)
	return session.RowsAffected, session.Error
}

func (d *dao) Update(ctx context.Context, gorm *gorm.DB, form *forms.RestUpdateForm) (int64, error) {
	session := gorm.Table(form.Table).Where("id = ?", form.Id)
	session.Updates(form.Params)
	return session.RowsAffected, session.Error
}

func (d *dao) Delete(ctx context.Context, gorm *gorm.DB, form *forms.GetForm) (int64, error) {
	session := gorm
	//session.Table(form.SourceName).Drop(nil, "10271") // 按主键对待
	session.Table(form.Table).Delete(nil, "id = ?", form.Id) // 按查询条件对待
	return session.RowsAffected, session.Error
}

func GetValue(filedType string, value interface{}) interface{} {
	// /usr/local/go/src/database/sql/export.go
	switch filedType {
	case "uint32":
		return cast.ToUint32(value)
	case "int32":
		return cast.ToInt32(value)
	case "int64":
		return cast.ToInt64(value)
	case "uint64":
		return cast.ToUint64(value)
	case "NullInt16":
		nullAble := **value.(**sql.NullInt16)
		if nullAble.Valid {
			return nullAble.Int16
		}
		return nil
	case "NullInt32":
		nullAble := **value.(**sql.NullInt32)
		if nullAble.Valid {
			return nullAble.Int32
		}
		return nil
	case "NullInt64":
		if value == nil {
			return nil
		}

		if *value.(**sql.NullInt64) == nil {
			return nil
		}
		nullAble := **value.(**sql.NullInt64)
		if nullAble.Valid {
			return nullAble.Int64
		}
		return nil
	case "RawBytes":
		if ret, ok := value.(**sql.RawBytes); ok {
			//return fmt.Sprintf("%v", *ret)
			if ret == nil || *ret == nil || **ret == nil {
				return nil
			}
			return string(**ret)
		}
		return "" // never go here
	case "NullTime":
		if *value.(**sql.NullTime) == nil {
			return nil
		}

		nullAble := **value.(**sql.NullTime)
		if nullAble.Valid {
			return nullAble.Time.Format("2006-01-02 15:04:05")
		} else {
			return ""
		}
	case "NullString":
		if *value.(**sql.NullString) == nil {
			return nil
		}

		nullAble := **value.(**sql.NullString)
		if nullAble.Valid {
			return nullAble.String
		}
		return nil
	case "NullByte":
		if *value.(**sql.NullByte) == nil {
			return nil
		}

		nullAble := **value.(**sql.NullByte)
		if nullAble.Valid {
			return string(nullAble.Byte)
		}
		return nil
	case "NullFloat64":
		if *value.(**sql.NullFloat64) == nil {
			return nil
		}

		nullAble := **value.(**sql.NullFloat64)
		if nullAble.Valid {
			return nullAble.Float64
		}
		return nil
	case "NullBool":
		if *value.(**sql.NullBool) == nil {
			return value // return nil 相同
		}

		nullAble := **value.(**sql.NullBool)
		if nullAble.Valid {
			return nullAble.Bool
		}
		return nil
	default:
		return value
	}
}

/**
来自gorm
/Users/owen/go/pkg/mod/gorm.io/gorm@v1.23.3/scan.go +14 PrepareValues
*/

func PrepareValues(columnTypes []*sql.ColumnType) []interface{} {
	var values = make([]interface{}, len(columnTypes))
	for idx, columnType := range columnTypes {
		// 针对 int not null 类型的优化， 导出excel 拿到地址，未拿到数据

		if columnType.ScanType() != nil {
			values[idx] = reflect.New(reflect.PtrTo(columnType.ScanType())).Interface()
		} else {
			values[idx] = new(interface{})
		}
	}

	return values
}
