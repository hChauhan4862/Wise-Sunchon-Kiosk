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

func newSysdiagram(db *gorm.DB, opts ...gen.DOOption) sysdiagram {
	_sysdiagram := sysdiagram{}

	_sysdiagram.sysdiagramDo.UseDB(db, opts...)
	_sysdiagram.sysdiagramDo.UseModel(&model.Sysdiagram{})

	tableName := _sysdiagram.sysdiagramDo.TableName()
	_sysdiagram.ALL = field.NewAsterisk(tableName)
	_sysdiagram.Name = field.NewString(tableName, "name")
	_sysdiagram.PrincipalID = field.NewInt64(tableName, "principal_id")
	_sysdiagram.DiagramID = field.NewInt64(tableName, "diagram_id")
	_sysdiagram.Version = field.NewInt64(tableName, "version")
	_sysdiagram.Definition = field.NewField(tableName, "definition")

	_sysdiagram.fillFieldMap()

	return _sysdiagram
}

type sysdiagram struct {
	sysdiagramDo

	ALL         field.Asterisk
	Name        field.String
	PrincipalID field.Int64
	DiagramID   field.Int64
	Version     field.Int64
	Definition  field.Field

	fieldMap map[string]field.Expr
}

func (s sysdiagram) Table(newTableName string) *sysdiagram {
	s.sysdiagramDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysdiagram) As(alias string) *sysdiagram {
	s.sysdiagramDo.DO = *(s.sysdiagramDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysdiagram) updateTableName(table string) *sysdiagram {
	s.ALL = field.NewAsterisk(table)
	s.Name = field.NewString(table, "name")
	s.PrincipalID = field.NewInt64(table, "principal_id")
	s.DiagramID = field.NewInt64(table, "diagram_id")
	s.Version = field.NewInt64(table, "version")
	s.Definition = field.NewField(table, "definition")

	s.fillFieldMap()

	return s
}

func (s *sysdiagram) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysdiagram) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["name"] = s.Name
	s.fieldMap["principal_id"] = s.PrincipalID
	s.fieldMap["diagram_id"] = s.DiagramID
	s.fieldMap["version"] = s.Version
	s.fieldMap["definition"] = s.Definition
}

func (s sysdiagram) clone(db *gorm.DB) sysdiagram {
	s.sysdiagramDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysdiagram) replaceDB(db *gorm.DB) sysdiagram {
	s.sysdiagramDo.ReplaceDB(db)
	return s
}

type sysdiagramDo struct{ gen.DO }

type ISysdiagramDo interface {
	gen.SubQuery
	Debug() ISysdiagramDo
	WithContext(ctx context.Context) ISysdiagramDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysdiagramDo
	WriteDB() ISysdiagramDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysdiagramDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysdiagramDo
	Not(conds ...gen.Condition) ISysdiagramDo
	Or(conds ...gen.Condition) ISysdiagramDo
	Select(conds ...field.Expr) ISysdiagramDo
	Where(conds ...gen.Condition) ISysdiagramDo
	Order(conds ...field.Expr) ISysdiagramDo
	Distinct(cols ...field.Expr) ISysdiagramDo
	Omit(cols ...field.Expr) ISysdiagramDo
	Join(table schema.Tabler, on ...field.Expr) ISysdiagramDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysdiagramDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysdiagramDo
	Group(cols ...field.Expr) ISysdiagramDo
	Having(conds ...gen.Condition) ISysdiagramDo
	Limit(limit int) ISysdiagramDo
	Offset(offset int) ISysdiagramDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysdiagramDo
	Unscoped() ISysdiagramDo
	Create(values ...*model.Sysdiagram) error
	CreateInBatches(values []*model.Sysdiagram, batchSize int) error
	Save(values ...*model.Sysdiagram) error
	First() (*model.Sysdiagram, error)
	Take() (*model.Sysdiagram, error)
	Last() (*model.Sysdiagram, error)
	Find() ([]*model.Sysdiagram, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sysdiagram, err error)
	FindInBatches(result *[]*model.Sysdiagram, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Sysdiagram) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysdiagramDo
	Assign(attrs ...field.AssignExpr) ISysdiagramDo
	Joins(fields ...field.RelationField) ISysdiagramDo
	Preload(fields ...field.RelationField) ISysdiagramDo
	FirstOrInit() (*model.Sysdiagram, error)
	FirstOrCreate() (*model.Sysdiagram, error)
	FindByPage(offset int, limit int) (result []*model.Sysdiagram, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysdiagramDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysdiagramDo) Debug() ISysdiagramDo {
	return s.withDO(s.DO.Debug())
}

func (s sysdiagramDo) WithContext(ctx context.Context) ISysdiagramDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysdiagramDo) ReadDB() ISysdiagramDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysdiagramDo) WriteDB() ISysdiagramDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysdiagramDo) Session(config *gorm.Session) ISysdiagramDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysdiagramDo) Clauses(conds ...clause.Expression) ISysdiagramDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysdiagramDo) Returning(value interface{}, columns ...string) ISysdiagramDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysdiagramDo) Not(conds ...gen.Condition) ISysdiagramDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysdiagramDo) Or(conds ...gen.Condition) ISysdiagramDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysdiagramDo) Select(conds ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysdiagramDo) Where(conds ...gen.Condition) ISysdiagramDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysdiagramDo) Order(conds ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysdiagramDo) Distinct(cols ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysdiagramDo) Omit(cols ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysdiagramDo) Join(table schema.Tabler, on ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysdiagramDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysdiagramDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysdiagramDo) Group(cols ...field.Expr) ISysdiagramDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysdiagramDo) Having(conds ...gen.Condition) ISysdiagramDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysdiagramDo) Limit(limit int) ISysdiagramDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysdiagramDo) Offset(offset int) ISysdiagramDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysdiagramDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysdiagramDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysdiagramDo) Unscoped() ISysdiagramDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysdiagramDo) Create(values ...*model.Sysdiagram) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysdiagramDo) CreateInBatches(values []*model.Sysdiagram, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysdiagramDo) Save(values ...*model.Sysdiagram) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysdiagramDo) First() (*model.Sysdiagram, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdiagram), nil
	}
}

func (s sysdiagramDo) Take() (*model.Sysdiagram, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdiagram), nil
	}
}

func (s sysdiagramDo) Last() (*model.Sysdiagram, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdiagram), nil
	}
}

func (s sysdiagramDo) Find() ([]*model.Sysdiagram, error) {
	result, err := s.DO.Find()
	return result.([]*model.Sysdiagram), err
}

func (s sysdiagramDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sysdiagram, err error) {
	buf := make([]*model.Sysdiagram, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysdiagramDo) FindInBatches(result *[]*model.Sysdiagram, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysdiagramDo) Attrs(attrs ...field.AssignExpr) ISysdiagramDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysdiagramDo) Assign(attrs ...field.AssignExpr) ISysdiagramDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysdiagramDo) Joins(fields ...field.RelationField) ISysdiagramDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysdiagramDo) Preload(fields ...field.RelationField) ISysdiagramDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysdiagramDo) FirstOrInit() (*model.Sysdiagram, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdiagram), nil
	}
}

func (s sysdiagramDo) FirstOrCreate() (*model.Sysdiagram, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdiagram), nil
	}
}

func (s sysdiagramDo) FindByPage(offset int, limit int) (result []*model.Sysdiagram, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysdiagramDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysdiagramDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysdiagramDo) Delete(models ...*model.Sysdiagram) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysdiagramDo) withDO(do gen.Dao) *sysdiagramDo {
	s.DO = *do.(*gen.DO)
	return s
}
