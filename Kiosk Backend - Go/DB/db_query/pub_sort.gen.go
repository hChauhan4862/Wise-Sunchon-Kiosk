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

func newPUBSORT(db *gorm.DB, opts ...gen.DOOption) pUBSORT {
	_pUBSORT := pUBSORT{}

	_pUBSORT.pUBSORTDo.UseDB(db, opts...)
	_pUBSORT.pUBSORTDo.UseModel(&model.PUBSORT{})

	tableName := _pUBSORT.pUBSORTDo.TableName()
	_pUBSORT.ALL = field.NewAsterisk(tableName)
	_pUBSORT.SORTCD = field.NewString(tableName, "SORT_CD")
	_pUBSORT.SORTNM = field.NewString(tableName, "SORT_NM")
	_pUBSORT.SORTSEQ = field.NewInt64(tableName, "SORT_SEQ")
	_pUBSORT.REGDT = field.NewString(tableName, "REG_DT")
	_pUBSORT.UPDDT = field.NewString(tableName, "UPD_DT")

	_pUBSORT.fillFieldMap()

	return _pUBSORT
}

type pUBSORT struct {
	pUBSORTDo

	ALL     field.Asterisk
	SORTCD  field.String
	SORTNM  field.String
	SORTSEQ field.Int64
	REGDT   field.String
	UPDDT   field.String

	fieldMap map[string]field.Expr
}

func (p pUBSORT) Table(newTableName string) *pUBSORT {
	p.pUBSORTDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pUBSORT) As(alias string) *pUBSORT {
	p.pUBSORTDo.DO = *(p.pUBSORTDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pUBSORT) updateTableName(table string) *pUBSORT {
	p.ALL = field.NewAsterisk(table)
	p.SORTCD = field.NewString(table, "SORT_CD")
	p.SORTNM = field.NewString(table, "SORT_NM")
	p.SORTSEQ = field.NewInt64(table, "SORT_SEQ")
	p.REGDT = field.NewString(table, "REG_DT")
	p.UPDDT = field.NewString(table, "UPD_DT")

	p.fillFieldMap()

	return p
}

func (p *pUBSORT) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pUBSORT) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["SORT_CD"] = p.SORTCD
	p.fieldMap["SORT_NM"] = p.SORTNM
	p.fieldMap["SORT_SEQ"] = p.SORTSEQ
	p.fieldMap["REG_DT"] = p.REGDT
	p.fieldMap["UPD_DT"] = p.UPDDT
}

func (p pUBSORT) clone(db *gorm.DB) pUBSORT {
	p.pUBSORTDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pUBSORT) replaceDB(db *gorm.DB) pUBSORT {
	p.pUBSORTDo.ReplaceDB(db)
	return p
}

type pUBSORTDo struct{ gen.DO }

type IPUBSORTDo interface {
	gen.SubQuery
	Debug() IPUBSORTDo
	WithContext(ctx context.Context) IPUBSORTDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPUBSORTDo
	WriteDB() IPUBSORTDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPUBSORTDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPUBSORTDo
	Not(conds ...gen.Condition) IPUBSORTDo
	Or(conds ...gen.Condition) IPUBSORTDo
	Select(conds ...field.Expr) IPUBSORTDo
	Where(conds ...gen.Condition) IPUBSORTDo
	Order(conds ...field.Expr) IPUBSORTDo
	Distinct(cols ...field.Expr) IPUBSORTDo
	Omit(cols ...field.Expr) IPUBSORTDo
	Join(table schema.Tabler, on ...field.Expr) IPUBSORTDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPUBSORTDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPUBSORTDo
	Group(cols ...field.Expr) IPUBSORTDo
	Having(conds ...gen.Condition) IPUBSORTDo
	Limit(limit int) IPUBSORTDo
	Offset(offset int) IPUBSORTDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPUBSORTDo
	Unscoped() IPUBSORTDo
	Create(values ...*model.PUBSORT) error
	CreateInBatches(values []*model.PUBSORT, batchSize int) error
	Save(values ...*model.PUBSORT) error
	First() (*model.PUBSORT, error)
	Take() (*model.PUBSORT, error)
	Last() (*model.PUBSORT, error)
	Find() ([]*model.PUBSORT, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PUBSORT, err error)
	FindInBatches(result *[]*model.PUBSORT, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PUBSORT) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPUBSORTDo
	Assign(attrs ...field.AssignExpr) IPUBSORTDo
	Joins(fields ...field.RelationField) IPUBSORTDo
	Preload(fields ...field.RelationField) IPUBSORTDo
	FirstOrInit() (*model.PUBSORT, error)
	FirstOrCreate() (*model.PUBSORT, error)
	FindByPage(offset int, limit int) (result []*model.PUBSORT, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPUBSORTDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pUBSORTDo) Debug() IPUBSORTDo {
	return p.withDO(p.DO.Debug())
}

func (p pUBSORTDo) WithContext(ctx context.Context) IPUBSORTDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pUBSORTDo) ReadDB() IPUBSORTDo {
	return p.Clauses(dbresolver.Read)
}

func (p pUBSORTDo) WriteDB() IPUBSORTDo {
	return p.Clauses(dbresolver.Write)
}

func (p pUBSORTDo) Session(config *gorm.Session) IPUBSORTDo {
	return p.withDO(p.DO.Session(config))
}

func (p pUBSORTDo) Clauses(conds ...clause.Expression) IPUBSORTDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pUBSORTDo) Returning(value interface{}, columns ...string) IPUBSORTDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pUBSORTDo) Not(conds ...gen.Condition) IPUBSORTDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pUBSORTDo) Or(conds ...gen.Condition) IPUBSORTDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pUBSORTDo) Select(conds ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pUBSORTDo) Where(conds ...gen.Condition) IPUBSORTDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pUBSORTDo) Order(conds ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pUBSORTDo) Distinct(cols ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pUBSORTDo) Omit(cols ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pUBSORTDo) Join(table schema.Tabler, on ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pUBSORTDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pUBSORTDo) RightJoin(table schema.Tabler, on ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pUBSORTDo) Group(cols ...field.Expr) IPUBSORTDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pUBSORTDo) Having(conds ...gen.Condition) IPUBSORTDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pUBSORTDo) Limit(limit int) IPUBSORTDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pUBSORTDo) Offset(offset int) IPUBSORTDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pUBSORTDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPUBSORTDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pUBSORTDo) Unscoped() IPUBSORTDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pUBSORTDo) Create(values ...*model.PUBSORT) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pUBSORTDo) CreateInBatches(values []*model.PUBSORT, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pUBSORTDo) Save(values ...*model.PUBSORT) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pUBSORTDo) First() (*model.PUBSORT, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PUBSORT), nil
	}
}

func (p pUBSORTDo) Take() (*model.PUBSORT, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PUBSORT), nil
	}
}

func (p pUBSORTDo) Last() (*model.PUBSORT, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PUBSORT), nil
	}
}

func (p pUBSORTDo) Find() ([]*model.PUBSORT, error) {
	result, err := p.DO.Find()
	return result.([]*model.PUBSORT), err
}

func (p pUBSORTDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PUBSORT, err error) {
	buf := make([]*model.PUBSORT, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pUBSORTDo) FindInBatches(result *[]*model.PUBSORT, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pUBSORTDo) Attrs(attrs ...field.AssignExpr) IPUBSORTDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pUBSORTDo) Assign(attrs ...field.AssignExpr) IPUBSORTDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pUBSORTDo) Joins(fields ...field.RelationField) IPUBSORTDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pUBSORTDo) Preload(fields ...field.RelationField) IPUBSORTDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pUBSORTDo) FirstOrInit() (*model.PUBSORT, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PUBSORT), nil
	}
}

func (p pUBSORTDo) FirstOrCreate() (*model.PUBSORT, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PUBSORT), nil
	}
}

func (p pUBSORTDo) FindByPage(offset int, limit int) (result []*model.PUBSORT, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pUBSORTDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pUBSORTDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pUBSORTDo) Delete(models ...*model.PUBSORT) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pUBSORTDo) withDO(do gen.Dao) *pUBSORTDo {
	p.DO = *do.(*gen.DO)
	return p
}