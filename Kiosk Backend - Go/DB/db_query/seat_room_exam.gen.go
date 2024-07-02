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

func newSeatRoomExam(db *gorm.DB, opts ...gen.DOOption) seatRoomExam {
	_seatRoomExam := seatRoomExam{}

	_seatRoomExam.seatRoomExamDo.UseDB(db, opts...)
	_seatRoomExam.seatRoomExamDo.UseModel(&model.SeatRoomExam{})

	tableName := _seatRoomExam.seatRoomExamDo.TableName()
	_seatRoomExam.ALL = field.NewAsterisk(tableName)
	_seatRoomExam.ExamGubun = field.NewString(tableName, "exam_gubun")
	_seatRoomExam.RoomNo = field.NewInt64(tableName, "room_no")
	_seatRoomExam.RoomName = field.NewString(tableName, "room_name")
	_seatRoomExam.SeatNoS = field.NewInt64(tableName, "seat_no_s")
	_seatRoomExam.SeatNoE = field.NewInt64(tableName, "seat_no_e")
	_seatRoomExam.SeatNoCnt = field.NewInt64(tableName, "seat_no_cnt")
	_seatRoomExam.ColorUse = field.NewInt64(tableName, "color_use")
	_seatRoomExam.ColorNuse = field.NewInt64(tableName, "color_nuse")
	_seatRoomExam.TimeStart = field.NewString(tableName, "time_start")
	_seatRoomExam.TimeEnd = field.NewString(tableName, "time_end")
	_seatRoomExam.UseHour = field.NewInt64(tableName, "use_hour")
	_seatRoomExam.ContMin = field.NewInt64(tableName, "cont_min")
	_seatRoomExam.ContCnt = field.NewInt64(tableName, "cont_cnt")
	_seatRoomExam.SeatFix = field.NewString(tableName, "seat_fix")
	_seatRoomExam.LimitPat = field.NewString(tableName, "limit_pat")
	_seatRoomExam.LimitPatGu = field.NewString(tableName, "limit_pat_gu")
	_seatRoomExam.LimitDept = field.NewString(tableName, "limit_dept")
	_seatRoomExam.LimitDeptGu = field.NewString(tableName, "limit_dept_gu")
	_seatRoomExam.LimitStat = field.NewString(tableName, "limit_stat")
	_seatRoomExam.LimitStatGu = field.NewString(tableName, "limit_stat_gu")
	_seatRoomExam.RoomBigo = field.NewString(tableName, "room_bigo")
	_seatRoomExam.ColorUsem = field.NewInt64(tableName, "color_usem")
	_seatRoomExam.ColorUsef = field.NewInt64(tableName, "color_usef")
	_seatRoomExam.LimitYearS = field.NewInt64(tableName, "limit_year_s")
	_seatRoomExam.LimitYearE = field.NewInt64(tableName, "limit_year_e")
	_seatRoomExam.LimitSeatpat = field.NewString(tableName, "limit_seatpat")
	_seatRoomExam.LimitSeatpatGu = field.NewInt64(tableName, "limit_seatpat_gu")
	_seatRoomExam.LimitPatcnt = field.NewString(tableName, "limit_patcnt")
	_seatRoomExam.LimitPatcntGu = field.NewInt64(tableName, "limit_patcnt_gu")
	_seatRoomExam.LimitGenderm = field.NewString(tableName, "limit_genderm")
	_seatRoomExam.LimitGendermGu = field.NewInt64(tableName, "limit_genderm_gu")
	_seatRoomExam.LimitGenderf = field.NewString(tableName, "limit_genderf")
	_seatRoomExam.LimitGenderfGu = field.NewInt64(tableName, "limit_genderf_gu")

	_seatRoomExam.fillFieldMap()

	return _seatRoomExam
}

type seatRoomExam struct {
	seatRoomExamDo

	ALL            field.Asterisk
	ExamGubun      field.String
	RoomNo         field.Int64
	RoomName       field.String
	SeatNoS        field.Int64
	SeatNoE        field.Int64
	SeatNoCnt      field.Int64
	ColorUse       field.Int64
	ColorNuse      field.Int64
	TimeStart      field.String
	TimeEnd        field.String
	UseHour        field.Int64
	ContMin        field.Int64
	ContCnt        field.Int64
	SeatFix        field.String
	LimitPat       field.String
	LimitPatGu     field.String
	LimitDept      field.String
	LimitDeptGu    field.String
	LimitStat      field.String
	LimitStatGu    field.String
	RoomBigo       field.String
	ColorUsem      field.Int64
	ColorUsef      field.Int64
	LimitYearS     field.Int64
	LimitYearE     field.Int64
	LimitSeatpat   field.String
	LimitSeatpatGu field.Int64
	LimitPatcnt    field.String
	LimitPatcntGu  field.Int64
	LimitGenderm   field.String
	LimitGendermGu field.Int64
	LimitGenderf   field.String
	LimitGenderfGu field.Int64

	fieldMap map[string]field.Expr
}

func (s seatRoomExam) Table(newTableName string) *seatRoomExam {
	s.seatRoomExamDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatRoomExam) As(alias string) *seatRoomExam {
	s.seatRoomExamDo.DO = *(s.seatRoomExamDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatRoomExam) updateTableName(table string) *seatRoomExam {
	s.ALL = field.NewAsterisk(table)
	s.ExamGubun = field.NewString(table, "exam_gubun")
	s.RoomNo = field.NewInt64(table, "room_no")
	s.RoomName = field.NewString(table, "room_name")
	s.SeatNoS = field.NewInt64(table, "seat_no_s")
	s.SeatNoE = field.NewInt64(table, "seat_no_e")
	s.SeatNoCnt = field.NewInt64(table, "seat_no_cnt")
	s.ColorUse = field.NewInt64(table, "color_use")
	s.ColorNuse = field.NewInt64(table, "color_nuse")
	s.TimeStart = field.NewString(table, "time_start")
	s.TimeEnd = field.NewString(table, "time_end")
	s.UseHour = field.NewInt64(table, "use_hour")
	s.ContMin = field.NewInt64(table, "cont_min")
	s.ContCnt = field.NewInt64(table, "cont_cnt")
	s.SeatFix = field.NewString(table, "seat_fix")
	s.LimitPat = field.NewString(table, "limit_pat")
	s.LimitPatGu = field.NewString(table, "limit_pat_gu")
	s.LimitDept = field.NewString(table, "limit_dept")
	s.LimitDeptGu = field.NewString(table, "limit_dept_gu")
	s.LimitStat = field.NewString(table, "limit_stat")
	s.LimitStatGu = field.NewString(table, "limit_stat_gu")
	s.RoomBigo = field.NewString(table, "room_bigo")
	s.ColorUsem = field.NewInt64(table, "color_usem")
	s.ColorUsef = field.NewInt64(table, "color_usef")
	s.LimitYearS = field.NewInt64(table, "limit_year_s")
	s.LimitYearE = field.NewInt64(table, "limit_year_e")
	s.LimitSeatpat = field.NewString(table, "limit_seatpat")
	s.LimitSeatpatGu = field.NewInt64(table, "limit_seatpat_gu")
	s.LimitPatcnt = field.NewString(table, "limit_patcnt")
	s.LimitPatcntGu = field.NewInt64(table, "limit_patcnt_gu")
	s.LimitGenderm = field.NewString(table, "limit_genderm")
	s.LimitGendermGu = field.NewInt64(table, "limit_genderm_gu")
	s.LimitGenderf = field.NewString(table, "limit_genderf")
	s.LimitGenderfGu = field.NewInt64(table, "limit_genderf_gu")

	s.fillFieldMap()

	return s
}

func (s *seatRoomExam) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatRoomExam) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 33)
	s.fieldMap["exam_gubun"] = s.ExamGubun
	s.fieldMap["room_no"] = s.RoomNo
	s.fieldMap["room_name"] = s.RoomName
	s.fieldMap["seat_no_s"] = s.SeatNoS
	s.fieldMap["seat_no_e"] = s.SeatNoE
	s.fieldMap["seat_no_cnt"] = s.SeatNoCnt
	s.fieldMap["color_use"] = s.ColorUse
	s.fieldMap["color_nuse"] = s.ColorNuse
	s.fieldMap["time_start"] = s.TimeStart
	s.fieldMap["time_end"] = s.TimeEnd
	s.fieldMap["use_hour"] = s.UseHour
	s.fieldMap["cont_min"] = s.ContMin
	s.fieldMap["cont_cnt"] = s.ContCnt
	s.fieldMap["seat_fix"] = s.SeatFix
	s.fieldMap["limit_pat"] = s.LimitPat
	s.fieldMap["limit_pat_gu"] = s.LimitPatGu
	s.fieldMap["limit_dept"] = s.LimitDept
	s.fieldMap["limit_dept_gu"] = s.LimitDeptGu
	s.fieldMap["limit_stat"] = s.LimitStat
	s.fieldMap["limit_stat_gu"] = s.LimitStatGu
	s.fieldMap["room_bigo"] = s.RoomBigo
	s.fieldMap["color_usem"] = s.ColorUsem
	s.fieldMap["color_usef"] = s.ColorUsef
	s.fieldMap["limit_year_s"] = s.LimitYearS
	s.fieldMap["limit_year_e"] = s.LimitYearE
	s.fieldMap["limit_seatpat"] = s.LimitSeatpat
	s.fieldMap["limit_seatpat_gu"] = s.LimitSeatpatGu
	s.fieldMap["limit_patcnt"] = s.LimitPatcnt
	s.fieldMap["limit_patcnt_gu"] = s.LimitPatcntGu
	s.fieldMap["limit_genderm"] = s.LimitGenderm
	s.fieldMap["limit_genderm_gu"] = s.LimitGendermGu
	s.fieldMap["limit_genderf"] = s.LimitGenderf
	s.fieldMap["limit_genderf_gu"] = s.LimitGenderfGu
}

func (s seatRoomExam) clone(db *gorm.DB) seatRoomExam {
	s.seatRoomExamDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatRoomExam) replaceDB(db *gorm.DB) seatRoomExam {
	s.seatRoomExamDo.ReplaceDB(db)
	return s
}

type seatRoomExamDo struct{ gen.DO }

type ISeatRoomExamDo interface {
	gen.SubQuery
	Debug() ISeatRoomExamDo
	WithContext(ctx context.Context) ISeatRoomExamDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatRoomExamDo
	WriteDB() ISeatRoomExamDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatRoomExamDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatRoomExamDo
	Not(conds ...gen.Condition) ISeatRoomExamDo
	Or(conds ...gen.Condition) ISeatRoomExamDo
	Select(conds ...field.Expr) ISeatRoomExamDo
	Where(conds ...gen.Condition) ISeatRoomExamDo
	Order(conds ...field.Expr) ISeatRoomExamDo
	Distinct(cols ...field.Expr) ISeatRoomExamDo
	Omit(cols ...field.Expr) ISeatRoomExamDo
	Join(table schema.Tabler, on ...field.Expr) ISeatRoomExamDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatRoomExamDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatRoomExamDo
	Group(cols ...field.Expr) ISeatRoomExamDo
	Having(conds ...gen.Condition) ISeatRoomExamDo
	Limit(limit int) ISeatRoomExamDo
	Offset(offset int) ISeatRoomExamDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatRoomExamDo
	Unscoped() ISeatRoomExamDo
	Create(values ...*model.SeatRoomExam) error
	CreateInBatches(values []*model.SeatRoomExam, batchSize int) error
	Save(values ...*model.SeatRoomExam) error
	First() (*model.SeatRoomExam, error)
	Take() (*model.SeatRoomExam, error)
	Last() (*model.SeatRoomExam, error)
	Find() ([]*model.SeatRoomExam, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatRoomExam, err error)
	FindInBatches(result *[]*model.SeatRoomExam, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatRoomExam) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatRoomExamDo
	Assign(attrs ...field.AssignExpr) ISeatRoomExamDo
	Joins(fields ...field.RelationField) ISeatRoomExamDo
	Preload(fields ...field.RelationField) ISeatRoomExamDo
	FirstOrInit() (*model.SeatRoomExam, error)
	FirstOrCreate() (*model.SeatRoomExam, error)
	FindByPage(offset int, limit int) (result []*model.SeatRoomExam, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatRoomExamDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatRoomExamDo) Debug() ISeatRoomExamDo {
	return s.withDO(s.DO.Debug())
}

func (s seatRoomExamDo) WithContext(ctx context.Context) ISeatRoomExamDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatRoomExamDo) ReadDB() ISeatRoomExamDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatRoomExamDo) WriteDB() ISeatRoomExamDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatRoomExamDo) Session(config *gorm.Session) ISeatRoomExamDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatRoomExamDo) Clauses(conds ...clause.Expression) ISeatRoomExamDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatRoomExamDo) Returning(value interface{}, columns ...string) ISeatRoomExamDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatRoomExamDo) Not(conds ...gen.Condition) ISeatRoomExamDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatRoomExamDo) Or(conds ...gen.Condition) ISeatRoomExamDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatRoomExamDo) Select(conds ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatRoomExamDo) Where(conds ...gen.Condition) ISeatRoomExamDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatRoomExamDo) Order(conds ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatRoomExamDo) Distinct(cols ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatRoomExamDo) Omit(cols ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatRoomExamDo) Join(table schema.Tabler, on ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatRoomExamDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatRoomExamDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatRoomExamDo) Group(cols ...field.Expr) ISeatRoomExamDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatRoomExamDo) Having(conds ...gen.Condition) ISeatRoomExamDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatRoomExamDo) Limit(limit int) ISeatRoomExamDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatRoomExamDo) Offset(offset int) ISeatRoomExamDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatRoomExamDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatRoomExamDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatRoomExamDo) Unscoped() ISeatRoomExamDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatRoomExamDo) Create(values ...*model.SeatRoomExam) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatRoomExamDo) CreateInBatches(values []*model.SeatRoomExam, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatRoomExamDo) Save(values ...*model.SeatRoomExam) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatRoomExamDo) First() (*model.SeatRoomExam, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoomExam), nil
	}
}

func (s seatRoomExamDo) Take() (*model.SeatRoomExam, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoomExam), nil
	}
}

func (s seatRoomExamDo) Last() (*model.SeatRoomExam, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoomExam), nil
	}
}

func (s seatRoomExamDo) Find() ([]*model.SeatRoomExam, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatRoomExam), err
}

func (s seatRoomExamDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatRoomExam, err error) {
	buf := make([]*model.SeatRoomExam, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatRoomExamDo) FindInBatches(result *[]*model.SeatRoomExam, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatRoomExamDo) Attrs(attrs ...field.AssignExpr) ISeatRoomExamDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatRoomExamDo) Assign(attrs ...field.AssignExpr) ISeatRoomExamDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatRoomExamDo) Joins(fields ...field.RelationField) ISeatRoomExamDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatRoomExamDo) Preload(fields ...field.RelationField) ISeatRoomExamDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatRoomExamDo) FirstOrInit() (*model.SeatRoomExam, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoomExam), nil
	}
}

func (s seatRoomExamDo) FirstOrCreate() (*model.SeatRoomExam, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoomExam), nil
	}
}

func (s seatRoomExamDo) FindByPage(offset int, limit int) (result []*model.SeatRoomExam, count int64, err error) {
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

func (s seatRoomExamDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatRoomExamDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatRoomExamDo) Delete(models ...*model.SeatRoomExam) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatRoomExamDo) withDO(do gen.Dao) *seatRoomExamDo {
	s.DO = *do.(*gen.DO)
	return s
}