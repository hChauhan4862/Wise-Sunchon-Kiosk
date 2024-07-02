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

func newLIBCONFIG(db *gorm.DB, opts ...gen.DOOption) lIBCONFIG {
	_lIBCONFIG := lIBCONFIG{}

	_lIBCONFIG.lIBCONFIGDo.UseDB(db, opts...)
	_lIBCONFIG.lIBCONFIGDo.UseModel(&model.LIBCONFIG{})

	tableName := _lIBCONFIG.lIBCONFIGDo.TableName()
	_lIBCONFIG.ALL = field.NewAsterisk(tableName)
	_lIBCONFIG.STATUS = field.NewInt64(tableName, "STATUS")
	_lIBCONFIG.MINSEATMIN = field.NewInt64(tableName, "MIN_SEAT_MIN")
	_lIBCONFIG.MAXSEATBKCNT = field.NewInt64(tableName, "MAX_SEAT_BK_CNT")
	_lIBCONFIG.MAXSEATBKMIN = field.NewInt64(tableName, "MAX_SEAT_BK_MIN")
	_lIBCONFIG.MAXSEATEACHMIN = field.NewInt64(tableName, "MAX_SEAT_EACH_MIN")
	_lIBCONFIG.MAXEQBKCNT = field.NewInt64(tableName, "MAX_EQ_BK_CNT")
	_lIBCONFIG.SEATAUTOCANCELMIN = field.NewInt64(tableName, "SEAT_AUTO_CANCEL_MIN")
	_lIBCONFIG.BOOKAUTOCANCELMIN = field.NewInt64(tableName, "BOOK_AUTO_CANCEL_MIN")
	_lIBCONFIG.EXTENDBEFORE = field.NewInt64(tableName, "EXTEND_BEFORE")
	_lIBCONFIG.MAXEQCNTPERBOOK = field.NewInt64(tableName, "MAX_EQ_CNT_PER_BOOK")
	_lIBCONFIG.BLDAYBEFORE = field.NewInt64(tableName, "BL_DAYBEFORE")
	_lIBCONFIG.BLLIMITCNT = field.NewInt64(tableName, "BL_LIMITCNT")
	_lIBCONFIG.BLNOTRETURNCNT = field.NewInt64(tableName, "BL_NOTRETURNCNT")
	_lIBCONFIG.BLBOOKINGDAY = field.NewInt64(tableName, "BL_BOOKINGDAY")
	_lIBCONFIG.BLASSIGNDAY = field.NewInt64(tableName, "BL_ASSIGNDAY")
	_lIBCONFIG.CKQUOTASEATTYPE = field.NewString(tableName, "CK_QUOTA_SEATTYPE")
	_lIBCONFIG.CKMONTHQUOTA = field.NewInt64(tableName, "CK_MONTH_QUOTA")
	_lIBCONFIG.TIMEOUTGATEINUSERMIN = field.NewInt64(tableName, "TIMEOUT_GATEIN_USER_MIN")
	_lIBCONFIG.LIBNO = field.NewInt64(tableName, "LIBNO")
	_lIBCONFIG.ISPCLOGINCHECK = field.NewInt64(tableName, "IS_PCLOGIN_CHECK")
	_lIBCONFIG.PCLOGINCHECKMIN = field.NewInt64(tableName, "PCLOGIN_CHECK_MIN")
	_lIBCONFIG.SMSBEFOREMIN = field.NewInt64(tableName, "SMS_BEFORE_MIN")
	_lIBCONFIG.SMSFROMPHONE = field.NewString(tableName, "SMS_FROMPHONE")
	_lIBCONFIG.MAXMEDIABKCNT = field.NewInt64(tableName, "MAX_MEDIA_BK_CNT")
	_lIBCONFIG.MAXMEDIACNTPERBOOK = field.NewInt64(tableName, "MAX_MEDIA_CNT_PER_BOOK")
	_lIBCONFIG.ISUSEACPOLICY = field.NewInt64(tableName, "IS_USE_ACPOLICY")
	_lIBCONFIG.ACBEFORMINON = field.NewInt64(tableName, "AC_BEFOR_MIN_ON")
	_lIBCONFIG.ISUSEPUSHMSG = field.NewInt64(tableName, "IS_USE_PUSHMSG")
	_lIBCONFIG.BEFORMINPUSHMSG = field.NewInt64(tableName, "BEFOR_MIN_PUSHMSG")
	_lIBCONFIG.BLBASEDATE = field.NewTime(tableName, "BL_BASEDATE")

	_lIBCONFIG.fillFieldMap()

	return _lIBCONFIG
}

type lIBCONFIG struct {
	lIBCONFIGDo

	ALL                  field.Asterisk
	STATUS               field.Int64
	MINSEATMIN           field.Int64
	MAXSEATBKCNT         field.Int64
	MAXSEATBKMIN         field.Int64
	MAXSEATEACHMIN       field.Int64
	MAXEQBKCNT           field.Int64
	SEATAUTOCANCELMIN    field.Int64
	BOOKAUTOCANCELMIN    field.Int64
	EXTENDBEFORE         field.Int64
	MAXEQCNTPERBOOK      field.Int64
	BLDAYBEFORE          field.Int64
	BLLIMITCNT           field.Int64
	BLNOTRETURNCNT       field.Int64
	BLBOOKINGDAY         field.Int64
	BLASSIGNDAY          field.Int64
	CKQUOTASEATTYPE      field.String
	CKMONTHQUOTA         field.Int64
	TIMEOUTGATEINUSERMIN field.Int64
	LIBNO                field.Int64
	ISPCLOGINCHECK       field.Int64
	PCLOGINCHECKMIN      field.Int64
	SMSBEFOREMIN         field.Int64
	SMSFROMPHONE         field.String
	MAXMEDIABKCNT        field.Int64
	MAXMEDIACNTPERBOOK   field.Int64
	ISUSEACPOLICY        field.Int64
	ACBEFORMINON         field.Int64
	ISUSEPUSHMSG         field.Int64
	BEFORMINPUSHMSG      field.Int64
	BLBASEDATE           field.Time

	fieldMap map[string]field.Expr
}

func (l lIBCONFIG) Table(newTableName string) *lIBCONFIG {
	l.lIBCONFIGDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lIBCONFIG) As(alias string) *lIBCONFIG {
	l.lIBCONFIGDo.DO = *(l.lIBCONFIGDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lIBCONFIG) updateTableName(table string) *lIBCONFIG {
	l.ALL = field.NewAsterisk(table)
	l.STATUS = field.NewInt64(table, "STATUS")
	l.MINSEATMIN = field.NewInt64(table, "MIN_SEAT_MIN")
	l.MAXSEATBKCNT = field.NewInt64(table, "MAX_SEAT_BK_CNT")
	l.MAXSEATBKMIN = field.NewInt64(table, "MAX_SEAT_BK_MIN")
	l.MAXSEATEACHMIN = field.NewInt64(table, "MAX_SEAT_EACH_MIN")
	l.MAXEQBKCNT = field.NewInt64(table, "MAX_EQ_BK_CNT")
	l.SEATAUTOCANCELMIN = field.NewInt64(table, "SEAT_AUTO_CANCEL_MIN")
	l.BOOKAUTOCANCELMIN = field.NewInt64(table, "BOOK_AUTO_CANCEL_MIN")
	l.EXTENDBEFORE = field.NewInt64(table, "EXTEND_BEFORE")
	l.MAXEQCNTPERBOOK = field.NewInt64(table, "MAX_EQ_CNT_PER_BOOK")
	l.BLDAYBEFORE = field.NewInt64(table, "BL_DAYBEFORE")
	l.BLLIMITCNT = field.NewInt64(table, "BL_LIMITCNT")
	l.BLNOTRETURNCNT = field.NewInt64(table, "BL_NOTRETURNCNT")
	l.BLBOOKINGDAY = field.NewInt64(table, "BL_BOOKINGDAY")
	l.BLASSIGNDAY = field.NewInt64(table, "BL_ASSIGNDAY")
	l.CKQUOTASEATTYPE = field.NewString(table, "CK_QUOTA_SEATTYPE")
	l.CKMONTHQUOTA = field.NewInt64(table, "CK_MONTH_QUOTA")
	l.TIMEOUTGATEINUSERMIN = field.NewInt64(table, "TIMEOUT_GATEIN_USER_MIN")
	l.LIBNO = field.NewInt64(table, "LIBNO")
	l.ISPCLOGINCHECK = field.NewInt64(table, "IS_PCLOGIN_CHECK")
	l.PCLOGINCHECKMIN = field.NewInt64(table, "PCLOGIN_CHECK_MIN")
	l.SMSBEFOREMIN = field.NewInt64(table, "SMS_BEFORE_MIN")
	l.SMSFROMPHONE = field.NewString(table, "SMS_FROMPHONE")
	l.MAXMEDIABKCNT = field.NewInt64(table, "MAX_MEDIA_BK_CNT")
	l.MAXMEDIACNTPERBOOK = field.NewInt64(table, "MAX_MEDIA_CNT_PER_BOOK")
	l.ISUSEACPOLICY = field.NewInt64(table, "IS_USE_ACPOLICY")
	l.ACBEFORMINON = field.NewInt64(table, "AC_BEFOR_MIN_ON")
	l.ISUSEPUSHMSG = field.NewInt64(table, "IS_USE_PUSHMSG")
	l.BEFORMINPUSHMSG = field.NewInt64(table, "BEFOR_MIN_PUSHMSG")
	l.BLBASEDATE = field.NewTime(table, "BL_BASEDATE")

	l.fillFieldMap()

	return l
}

func (l *lIBCONFIG) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lIBCONFIG) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 30)
	l.fieldMap["STATUS"] = l.STATUS
	l.fieldMap["MIN_SEAT_MIN"] = l.MINSEATMIN
	l.fieldMap["MAX_SEAT_BK_CNT"] = l.MAXSEATBKCNT
	l.fieldMap["MAX_SEAT_BK_MIN"] = l.MAXSEATBKMIN
	l.fieldMap["MAX_SEAT_EACH_MIN"] = l.MAXSEATEACHMIN
	l.fieldMap["MAX_EQ_BK_CNT"] = l.MAXEQBKCNT
	l.fieldMap["SEAT_AUTO_CANCEL_MIN"] = l.SEATAUTOCANCELMIN
	l.fieldMap["BOOK_AUTO_CANCEL_MIN"] = l.BOOKAUTOCANCELMIN
	l.fieldMap["EXTEND_BEFORE"] = l.EXTENDBEFORE
	l.fieldMap["MAX_EQ_CNT_PER_BOOK"] = l.MAXEQCNTPERBOOK
	l.fieldMap["BL_DAYBEFORE"] = l.BLDAYBEFORE
	l.fieldMap["BL_LIMITCNT"] = l.BLLIMITCNT
	l.fieldMap["BL_NOTRETURNCNT"] = l.BLNOTRETURNCNT
	l.fieldMap["BL_BOOKINGDAY"] = l.BLBOOKINGDAY
	l.fieldMap["BL_ASSIGNDAY"] = l.BLASSIGNDAY
	l.fieldMap["CK_QUOTA_SEATTYPE"] = l.CKQUOTASEATTYPE
	l.fieldMap["CK_MONTH_QUOTA"] = l.CKMONTHQUOTA
	l.fieldMap["TIMEOUT_GATEIN_USER_MIN"] = l.TIMEOUTGATEINUSERMIN
	l.fieldMap["LIBNO"] = l.LIBNO
	l.fieldMap["IS_PCLOGIN_CHECK"] = l.ISPCLOGINCHECK
	l.fieldMap["PCLOGIN_CHECK_MIN"] = l.PCLOGINCHECKMIN
	l.fieldMap["SMS_BEFORE_MIN"] = l.SMSBEFOREMIN
	l.fieldMap["SMS_FROMPHONE"] = l.SMSFROMPHONE
	l.fieldMap["MAX_MEDIA_BK_CNT"] = l.MAXMEDIABKCNT
	l.fieldMap["MAX_MEDIA_CNT_PER_BOOK"] = l.MAXMEDIACNTPERBOOK
	l.fieldMap["IS_USE_ACPOLICY"] = l.ISUSEACPOLICY
	l.fieldMap["AC_BEFOR_MIN_ON"] = l.ACBEFORMINON
	l.fieldMap["IS_USE_PUSHMSG"] = l.ISUSEPUSHMSG
	l.fieldMap["BEFOR_MIN_PUSHMSG"] = l.BEFORMINPUSHMSG
	l.fieldMap["BL_BASEDATE"] = l.BLBASEDATE
}

func (l lIBCONFIG) clone(db *gorm.DB) lIBCONFIG {
	l.lIBCONFIGDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l lIBCONFIG) replaceDB(db *gorm.DB) lIBCONFIG {
	l.lIBCONFIGDo.ReplaceDB(db)
	return l
}

type lIBCONFIGDo struct{ gen.DO }

type ILIBCONFIGDo interface {
	gen.SubQuery
	Debug() ILIBCONFIGDo
	WithContext(ctx context.Context) ILIBCONFIGDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILIBCONFIGDo
	WriteDB() ILIBCONFIGDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILIBCONFIGDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILIBCONFIGDo
	Not(conds ...gen.Condition) ILIBCONFIGDo
	Or(conds ...gen.Condition) ILIBCONFIGDo
	Select(conds ...field.Expr) ILIBCONFIGDo
	Where(conds ...gen.Condition) ILIBCONFIGDo
	Order(conds ...field.Expr) ILIBCONFIGDo
	Distinct(cols ...field.Expr) ILIBCONFIGDo
	Omit(cols ...field.Expr) ILIBCONFIGDo
	Join(table schema.Tabler, on ...field.Expr) ILIBCONFIGDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILIBCONFIGDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILIBCONFIGDo
	Group(cols ...field.Expr) ILIBCONFIGDo
	Having(conds ...gen.Condition) ILIBCONFIGDo
	Limit(limit int) ILIBCONFIGDo
	Offset(offset int) ILIBCONFIGDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILIBCONFIGDo
	Unscoped() ILIBCONFIGDo
	Create(values ...*model.LIBCONFIG) error
	CreateInBatches(values []*model.LIBCONFIG, batchSize int) error
	Save(values ...*model.LIBCONFIG) error
	First() (*model.LIBCONFIG, error)
	Take() (*model.LIBCONFIG, error)
	Last() (*model.LIBCONFIG, error)
	Find() ([]*model.LIBCONFIG, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LIBCONFIG, err error)
	FindInBatches(result *[]*model.LIBCONFIG, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LIBCONFIG) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILIBCONFIGDo
	Assign(attrs ...field.AssignExpr) ILIBCONFIGDo
	Joins(fields ...field.RelationField) ILIBCONFIGDo
	Preload(fields ...field.RelationField) ILIBCONFIGDo
	FirstOrInit() (*model.LIBCONFIG, error)
	FirstOrCreate() (*model.LIBCONFIG, error)
	FindByPage(offset int, limit int) (result []*model.LIBCONFIG, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILIBCONFIGDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l lIBCONFIGDo) Debug() ILIBCONFIGDo {
	return l.withDO(l.DO.Debug())
}

func (l lIBCONFIGDo) WithContext(ctx context.Context) ILIBCONFIGDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lIBCONFIGDo) ReadDB() ILIBCONFIGDo {
	return l.Clauses(dbresolver.Read)
}

func (l lIBCONFIGDo) WriteDB() ILIBCONFIGDo {
	return l.Clauses(dbresolver.Write)
}

func (l lIBCONFIGDo) Session(config *gorm.Session) ILIBCONFIGDo {
	return l.withDO(l.DO.Session(config))
}

func (l lIBCONFIGDo) Clauses(conds ...clause.Expression) ILIBCONFIGDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lIBCONFIGDo) Returning(value interface{}, columns ...string) ILIBCONFIGDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lIBCONFIGDo) Not(conds ...gen.Condition) ILIBCONFIGDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lIBCONFIGDo) Or(conds ...gen.Condition) ILIBCONFIGDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lIBCONFIGDo) Select(conds ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lIBCONFIGDo) Where(conds ...gen.Condition) ILIBCONFIGDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lIBCONFIGDo) Order(conds ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lIBCONFIGDo) Distinct(cols ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lIBCONFIGDo) Omit(cols ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lIBCONFIGDo) Join(table schema.Tabler, on ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lIBCONFIGDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lIBCONFIGDo) RightJoin(table schema.Tabler, on ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lIBCONFIGDo) Group(cols ...field.Expr) ILIBCONFIGDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lIBCONFIGDo) Having(conds ...gen.Condition) ILIBCONFIGDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lIBCONFIGDo) Limit(limit int) ILIBCONFIGDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lIBCONFIGDo) Offset(offset int) ILIBCONFIGDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lIBCONFIGDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILIBCONFIGDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lIBCONFIGDo) Unscoped() ILIBCONFIGDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lIBCONFIGDo) Create(values ...*model.LIBCONFIG) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lIBCONFIGDo) CreateInBatches(values []*model.LIBCONFIG, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lIBCONFIGDo) Save(values ...*model.LIBCONFIG) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lIBCONFIGDo) First() (*model.LIBCONFIG, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LIBCONFIG), nil
	}
}

func (l lIBCONFIGDo) Take() (*model.LIBCONFIG, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LIBCONFIG), nil
	}
}

func (l lIBCONFIGDo) Last() (*model.LIBCONFIG, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LIBCONFIG), nil
	}
}

func (l lIBCONFIGDo) Find() ([]*model.LIBCONFIG, error) {
	result, err := l.DO.Find()
	return result.([]*model.LIBCONFIG), err
}

func (l lIBCONFIGDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LIBCONFIG, err error) {
	buf := make([]*model.LIBCONFIG, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lIBCONFIGDo) FindInBatches(result *[]*model.LIBCONFIG, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lIBCONFIGDo) Attrs(attrs ...field.AssignExpr) ILIBCONFIGDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lIBCONFIGDo) Assign(attrs ...field.AssignExpr) ILIBCONFIGDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lIBCONFIGDo) Joins(fields ...field.RelationField) ILIBCONFIGDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l lIBCONFIGDo) Preload(fields ...field.RelationField) ILIBCONFIGDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l lIBCONFIGDo) FirstOrInit() (*model.LIBCONFIG, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LIBCONFIG), nil
	}
}

func (l lIBCONFIGDo) FirstOrCreate() (*model.LIBCONFIG, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LIBCONFIG), nil
	}
}

func (l lIBCONFIGDo) FindByPage(offset int, limit int) (result []*model.LIBCONFIG, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l lIBCONFIGDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l lIBCONFIGDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l lIBCONFIGDo) Delete(models ...*model.LIBCONFIG) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *lIBCONFIGDo) withDO(do gen.Dao) *lIBCONFIGDo {
	l.DO = *do.(*gen.DO)
	return l
}
