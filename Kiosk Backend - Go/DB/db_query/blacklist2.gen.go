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

func newBLACKLIST2(db *gorm.DB, opts ...gen.DOOption) bLACKLIST2 {
	_bLACKLIST2 := bLACKLIST2{}

	_bLACKLIST2.bLACKLIST2Do.UseDB(db, opts...)
	_bLACKLIST2.bLACKLIST2Do.UseModel(&model.BLACKLIST2{})

	tableName := _bLACKLIST2.bLACKLIST2Do.TableName()
	_bLACKLIST2.ALL = field.NewAsterisk(tableName)
	_bLACKLIST2.BLSEQNO = field.NewInt64(tableName, "BLSEQNO")
	_bLACKLIST2.SCHOOLNO = field.NewString(tableName, "SCHOOLNO")
	_bLACKLIST2.USERNAME = field.NewString(tableName, "USER_NAME")
	_bLACKLIST2.REASON = field.NewString(tableName, "REASON")
	_bLACKLIST2.ADMINID = field.NewString(tableName, "ADMINID")
	_bLACKLIST2.ISSUEFROM = field.NewInt64(tableName, "ISSUEFROM")
	_bLACKLIST2.ISSUEDETAIL = field.NewString(tableName, "ISSUEDETAIL")
	_bLACKLIST2.REGTIME = field.NewTime(tableName, "REGTIME")
	_bLACKLIST2.STATUS = field.NewInt64(tableName, "STATUS")
	_bLACKLIST2.BLOCKSTART = field.NewTime(tableName, "BLOCKSTART")
	_bLACKLIST2.BLOCKEXPIRE = field.NewTime(tableName, "BLOCKEXPIRE")

	_bLACKLIST2.fillFieldMap()

	return _bLACKLIST2
}

type bLACKLIST2 struct {
	bLACKLIST2Do

	ALL         field.Asterisk
	BLSEQNO     field.Int64
	SCHOOLNO    field.String
	USERNAME    field.String
	REASON      field.String
	ADMINID     field.String
	ISSUEFROM   field.Int64
	ISSUEDETAIL field.String
	REGTIME     field.Time
	STATUS      field.Int64
	BLOCKSTART  field.Time
	BLOCKEXPIRE field.Time

	fieldMap map[string]field.Expr
}

func (b bLACKLIST2) Table(newTableName string) *bLACKLIST2 {
	b.bLACKLIST2Do.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bLACKLIST2) As(alias string) *bLACKLIST2 {
	b.bLACKLIST2Do.DO = *(b.bLACKLIST2Do.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bLACKLIST2) updateTableName(table string) *bLACKLIST2 {
	b.ALL = field.NewAsterisk(table)
	b.BLSEQNO = field.NewInt64(table, "BLSEQNO")
	b.SCHOOLNO = field.NewString(table, "SCHOOLNO")
	b.USERNAME = field.NewString(table, "USER_NAME")
	b.REASON = field.NewString(table, "REASON")
	b.ADMINID = field.NewString(table, "ADMINID")
	b.ISSUEFROM = field.NewInt64(table, "ISSUEFROM")
	b.ISSUEDETAIL = field.NewString(table, "ISSUEDETAIL")
	b.REGTIME = field.NewTime(table, "REGTIME")
	b.STATUS = field.NewInt64(table, "STATUS")
	b.BLOCKSTART = field.NewTime(table, "BLOCKSTART")
	b.BLOCKEXPIRE = field.NewTime(table, "BLOCKEXPIRE")

	b.fillFieldMap()

	return b
}

func (b *bLACKLIST2) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bLACKLIST2) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 11)
	b.fieldMap["BLSEQNO"] = b.BLSEQNO
	b.fieldMap["SCHOOLNO"] = b.SCHOOLNO
	b.fieldMap["USER_NAME"] = b.USERNAME
	b.fieldMap["REASON"] = b.REASON
	b.fieldMap["ADMINID"] = b.ADMINID
	b.fieldMap["ISSUEFROM"] = b.ISSUEFROM
	b.fieldMap["ISSUEDETAIL"] = b.ISSUEDETAIL
	b.fieldMap["REGTIME"] = b.REGTIME
	b.fieldMap["STATUS"] = b.STATUS
	b.fieldMap["BLOCKSTART"] = b.BLOCKSTART
	b.fieldMap["BLOCKEXPIRE"] = b.BLOCKEXPIRE
}

func (b bLACKLIST2) clone(db *gorm.DB) bLACKLIST2 {
	b.bLACKLIST2Do.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bLACKLIST2) replaceDB(db *gorm.DB) bLACKLIST2 {
	b.bLACKLIST2Do.ReplaceDB(db)
	return b
}

type bLACKLIST2Do struct{ gen.DO }

type IBLACKLIST2Do interface {
	gen.SubQuery
	Debug() IBLACKLIST2Do
	WithContext(ctx context.Context) IBLACKLIST2Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBLACKLIST2Do
	WriteDB() IBLACKLIST2Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBLACKLIST2Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBLACKLIST2Do
	Not(conds ...gen.Condition) IBLACKLIST2Do
	Or(conds ...gen.Condition) IBLACKLIST2Do
	Select(conds ...field.Expr) IBLACKLIST2Do
	Where(conds ...gen.Condition) IBLACKLIST2Do
	Order(conds ...field.Expr) IBLACKLIST2Do
	Distinct(cols ...field.Expr) IBLACKLIST2Do
	Omit(cols ...field.Expr) IBLACKLIST2Do
	Join(table schema.Tabler, on ...field.Expr) IBLACKLIST2Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IBLACKLIST2Do
	RightJoin(table schema.Tabler, on ...field.Expr) IBLACKLIST2Do
	Group(cols ...field.Expr) IBLACKLIST2Do
	Having(conds ...gen.Condition) IBLACKLIST2Do
	Limit(limit int) IBLACKLIST2Do
	Offset(offset int) IBLACKLIST2Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBLACKLIST2Do
	Unscoped() IBLACKLIST2Do
	Create(values ...*model.BLACKLIST2) error
	CreateInBatches(values []*model.BLACKLIST2, batchSize int) error
	Save(values ...*model.BLACKLIST2) error
	First() (*model.BLACKLIST2, error)
	Take() (*model.BLACKLIST2, error)
	Last() (*model.BLACKLIST2, error)
	Find() ([]*model.BLACKLIST2, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BLACKLIST2, err error)
	FindInBatches(result *[]*model.BLACKLIST2, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BLACKLIST2) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBLACKLIST2Do
	Assign(attrs ...field.AssignExpr) IBLACKLIST2Do
	Joins(fields ...field.RelationField) IBLACKLIST2Do
	Preload(fields ...field.RelationField) IBLACKLIST2Do
	FirstOrInit() (*model.BLACKLIST2, error)
	FirstOrCreate() (*model.BLACKLIST2, error)
	FindByPage(offset int, limit int) (result []*model.BLACKLIST2, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBLACKLIST2Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bLACKLIST2Do) Debug() IBLACKLIST2Do {
	return b.withDO(b.DO.Debug())
}

func (b bLACKLIST2Do) WithContext(ctx context.Context) IBLACKLIST2Do {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bLACKLIST2Do) ReadDB() IBLACKLIST2Do {
	return b.Clauses(dbresolver.Read)
}

func (b bLACKLIST2Do) WriteDB() IBLACKLIST2Do {
	return b.Clauses(dbresolver.Write)
}

func (b bLACKLIST2Do) Session(config *gorm.Session) IBLACKLIST2Do {
	return b.withDO(b.DO.Session(config))
}

func (b bLACKLIST2Do) Clauses(conds ...clause.Expression) IBLACKLIST2Do {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bLACKLIST2Do) Returning(value interface{}, columns ...string) IBLACKLIST2Do {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bLACKLIST2Do) Not(conds ...gen.Condition) IBLACKLIST2Do {
	return b.withDO(b.DO.Not(conds...))
}

func (b bLACKLIST2Do) Or(conds ...gen.Condition) IBLACKLIST2Do {
	return b.withDO(b.DO.Or(conds...))
}

func (b bLACKLIST2Do) Select(conds ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.Select(conds...))
}

func (b bLACKLIST2Do) Where(conds ...gen.Condition) IBLACKLIST2Do {
	return b.withDO(b.DO.Where(conds...))
}

func (b bLACKLIST2Do) Order(conds ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.Order(conds...))
}

func (b bLACKLIST2Do) Distinct(cols ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bLACKLIST2Do) Omit(cols ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bLACKLIST2Do) Join(table schema.Tabler, on ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bLACKLIST2Do) LeftJoin(table schema.Tabler, on ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bLACKLIST2Do) RightJoin(table schema.Tabler, on ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bLACKLIST2Do) Group(cols ...field.Expr) IBLACKLIST2Do {
	return b.withDO(b.DO.Group(cols...))
}

func (b bLACKLIST2Do) Having(conds ...gen.Condition) IBLACKLIST2Do {
	return b.withDO(b.DO.Having(conds...))
}

func (b bLACKLIST2Do) Limit(limit int) IBLACKLIST2Do {
	return b.withDO(b.DO.Limit(limit))
}

func (b bLACKLIST2Do) Offset(offset int) IBLACKLIST2Do {
	return b.withDO(b.DO.Offset(offset))
}

func (b bLACKLIST2Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IBLACKLIST2Do {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bLACKLIST2Do) Unscoped() IBLACKLIST2Do {
	return b.withDO(b.DO.Unscoped())
}

func (b bLACKLIST2Do) Create(values ...*model.BLACKLIST2) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bLACKLIST2Do) CreateInBatches(values []*model.BLACKLIST2, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bLACKLIST2Do) Save(values ...*model.BLACKLIST2) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bLACKLIST2Do) First() (*model.BLACKLIST2, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BLACKLIST2), nil
	}
}

func (b bLACKLIST2Do) Take() (*model.BLACKLIST2, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BLACKLIST2), nil
	}
}

func (b bLACKLIST2Do) Last() (*model.BLACKLIST2, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BLACKLIST2), nil
	}
}

func (b bLACKLIST2Do) Find() ([]*model.BLACKLIST2, error) {
	result, err := b.DO.Find()
	return result.([]*model.BLACKLIST2), err
}

func (b bLACKLIST2Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BLACKLIST2, err error) {
	buf := make([]*model.BLACKLIST2, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bLACKLIST2Do) FindInBatches(result *[]*model.BLACKLIST2, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bLACKLIST2Do) Attrs(attrs ...field.AssignExpr) IBLACKLIST2Do {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bLACKLIST2Do) Assign(attrs ...field.AssignExpr) IBLACKLIST2Do {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bLACKLIST2Do) Joins(fields ...field.RelationField) IBLACKLIST2Do {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bLACKLIST2Do) Preload(fields ...field.RelationField) IBLACKLIST2Do {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bLACKLIST2Do) FirstOrInit() (*model.BLACKLIST2, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BLACKLIST2), nil
	}
}

func (b bLACKLIST2Do) FirstOrCreate() (*model.BLACKLIST2, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BLACKLIST2), nil
	}
}

func (b bLACKLIST2Do) FindByPage(offset int, limit int) (result []*model.BLACKLIST2, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bLACKLIST2Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bLACKLIST2Do) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bLACKLIST2Do) Delete(models ...*model.BLACKLIST2) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bLACKLIST2Do) withDO(do gen.Dao) *bLACKLIST2Do {
	b.DO = *do.(*gen.DO)
	return b
}
