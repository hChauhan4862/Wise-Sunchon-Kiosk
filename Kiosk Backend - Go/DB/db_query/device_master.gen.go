// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db_query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"WISE_SOFTWARE/DB/model"
)

func newDEVICEMASTER(db *gorm.DB, opts ...gen.DOOption) dEVICEMASTER {
	_dEVICEMASTER := dEVICEMASTER{}

	_dEVICEMASTER.dEVICEMASTERDo.UseDB(db, opts...)
	_dEVICEMASTER.dEVICEMASTERDo.UseModel(&model.DEVICEMASTER{})

	tableName := _dEVICEMASTER.dEVICEMASTERDo.TableName()
	_dEVICEMASTER.ALL = field.NewAsterisk(tableName)
	_dEVICEMASTER.SERIALNO = field.NewString(tableName, "SERIAL_NO")

	_dEVICEMASTER.fillFieldMap()

	return _dEVICEMASTER
}

type dEVICEMASTER struct {
	dEVICEMASTERDo

	ALL      field.Asterisk
	SERIALNO field.String

	fieldMap map[string]field.Expr
}

func (d dEVICEMASTER) Table(newTableName string) *dEVICEMASTER {
	d.dEVICEMASTERDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dEVICEMASTER) As(alias string) *dEVICEMASTER {
	d.dEVICEMASTERDo.DO = *(d.dEVICEMASTERDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dEVICEMASTER) updateTableName(table string) *dEVICEMASTER {
	d.ALL = field.NewAsterisk(table)
	d.SERIALNO = field.NewString(table, "SERIAL_NO")

	d.fillFieldMap()

	return d
}

func (d *dEVICEMASTER) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dEVICEMASTER) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 1)
	d.fieldMap["SERIAL_NO"] = d.SERIALNO
}

func (d dEVICEMASTER) clone(db *gorm.DB) dEVICEMASTER {
	d.dEVICEMASTERDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dEVICEMASTER) replaceDB(db *gorm.DB) dEVICEMASTER {
	d.dEVICEMASTERDo.ReplaceDB(db)
	return d
}

type dEVICEMASTERDo struct{ gen.DO }

type IDEVICEMASTERDo interface {
	gen.SubQuery
	Debug() IDEVICEMASTERDo
	WithContext(ctx context.Context) IDEVICEMASTERDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDEVICEMASTERDo
	WriteDB() IDEVICEMASTERDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDEVICEMASTERDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDEVICEMASTERDo
	Not(conds ...gen.Condition) IDEVICEMASTERDo
	Or(conds ...gen.Condition) IDEVICEMASTERDo
	Select(conds ...field.Expr) IDEVICEMASTERDo
	Where(conds ...gen.Condition) IDEVICEMASTERDo
	Order(conds ...field.Expr) IDEVICEMASTERDo
	Distinct(cols ...field.Expr) IDEVICEMASTERDo
	Omit(cols ...field.Expr) IDEVICEMASTERDo
	Join(table schema.Tabler, on ...field.Expr) IDEVICEMASTERDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDEVICEMASTERDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDEVICEMASTERDo
	Group(cols ...field.Expr) IDEVICEMASTERDo
	Having(conds ...gen.Condition) IDEVICEMASTERDo
	Limit(limit int) IDEVICEMASTERDo
	Offset(offset int) IDEVICEMASTERDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDEVICEMASTERDo
	Unscoped() IDEVICEMASTERDo
	Create(values ...*model.DEVICEMASTER) error
	CreateInBatches(values []*model.DEVICEMASTER, batchSize int) error
	Save(values ...*model.DEVICEMASTER) error
	First() (*model.DEVICEMASTER, error)
	Take() (*model.DEVICEMASTER, error)
	Last() (*model.DEVICEMASTER, error)
	Find() ([]*model.DEVICEMASTER, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DEVICEMASTER, err error)
	FindInBatches(result *[]*model.DEVICEMASTER, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DEVICEMASTER) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDEVICEMASTERDo
	Assign(attrs ...field.AssignExpr) IDEVICEMASTERDo
	Joins(fields ...field.RelationField) IDEVICEMASTERDo
	Preload(fields ...field.RelationField) IDEVICEMASTERDo
	FirstOrInit() (*model.DEVICEMASTER, error)
	FirstOrCreate() (*model.DEVICEMASTER, error)
	FindByPage(offset int, limit int) (result []*model.DEVICEMASTER, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDEVICEMASTERDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dEVICEMASTERDo) Debug() IDEVICEMASTERDo {
	return d.withDO(d.DO.Debug())
}

func (d dEVICEMASTERDo) WithContext(ctx context.Context) IDEVICEMASTERDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dEVICEMASTERDo) ReadDB() IDEVICEMASTERDo {
	return d.Clauses(dbresolver.Read)
}

func (d dEVICEMASTERDo) WriteDB() IDEVICEMASTERDo {
	return d.Clauses(dbresolver.Write)
}

func (d dEVICEMASTERDo) Session(config *gorm.Session) IDEVICEMASTERDo {
	return d.withDO(d.DO.Session(config))
}

func (d dEVICEMASTERDo) Clauses(conds ...clause.Expression) IDEVICEMASTERDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dEVICEMASTERDo) Returning(value interface{}, columns ...string) IDEVICEMASTERDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dEVICEMASTERDo) Not(conds ...gen.Condition) IDEVICEMASTERDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dEVICEMASTERDo) Or(conds ...gen.Condition) IDEVICEMASTERDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dEVICEMASTERDo) Select(conds ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dEVICEMASTERDo) Where(conds ...gen.Condition) IDEVICEMASTERDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dEVICEMASTERDo) Order(conds ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dEVICEMASTERDo) Distinct(cols ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dEVICEMASTERDo) Omit(cols ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dEVICEMASTERDo) Join(table schema.Tabler, on ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dEVICEMASTERDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dEVICEMASTERDo) RightJoin(table schema.Tabler, on ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dEVICEMASTERDo) Group(cols ...field.Expr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dEVICEMASTERDo) Having(conds ...gen.Condition) IDEVICEMASTERDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dEVICEMASTERDo) Limit(limit int) IDEVICEMASTERDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dEVICEMASTERDo) Offset(offset int) IDEVICEMASTERDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dEVICEMASTERDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDEVICEMASTERDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dEVICEMASTERDo) Unscoped() IDEVICEMASTERDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dEVICEMASTERDo) Create(values ...*model.DEVICEMASTER) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dEVICEMASTERDo) CreateInBatches(values []*model.DEVICEMASTER, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dEVICEMASTERDo) Save(values ...*model.DEVICEMASTER) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dEVICEMASTERDo) First() (*model.DEVICEMASTER, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DEVICEMASTER), nil
	}
}

func (d dEVICEMASTERDo) Take() (*model.DEVICEMASTER, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DEVICEMASTER), nil
	}
}

func (d dEVICEMASTERDo) Last() (*model.DEVICEMASTER, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DEVICEMASTER), nil
	}
}

func (d dEVICEMASTERDo) Find() ([]*model.DEVICEMASTER, error) {
	result, err := d.DO.Find()
	return result.([]*model.DEVICEMASTER), err
}

func (d dEVICEMASTERDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DEVICEMASTER, err error) {
	buf := make([]*model.DEVICEMASTER, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dEVICEMASTERDo) FindInBatches(result *[]*model.DEVICEMASTER, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dEVICEMASTERDo) Attrs(attrs ...field.AssignExpr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dEVICEMASTERDo) Assign(attrs ...field.AssignExpr) IDEVICEMASTERDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dEVICEMASTERDo) Joins(fields ...field.RelationField) IDEVICEMASTERDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dEVICEMASTERDo) Preload(fields ...field.RelationField) IDEVICEMASTERDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dEVICEMASTERDo) FirstOrInit() (*model.DEVICEMASTER, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DEVICEMASTER), nil
	}
}

func (d dEVICEMASTERDo) FirstOrCreate() (*model.DEVICEMASTER, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DEVICEMASTER), nil
	}
}

func (d dEVICEMASTERDo) FindByPage(offset int, limit int) (result []*model.DEVICEMASTER, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dEVICEMASTERDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dEVICEMASTERDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dEVICEMASTERDo) Delete(models ...*model.DEVICEMASTER) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dEVICEMASTERDo) withDO(do gen.Dao) *dEVICEMASTERDo {
	d.DO = *do.(*gen.DO)
	return d
}
