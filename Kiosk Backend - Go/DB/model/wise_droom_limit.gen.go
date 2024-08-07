// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseDroomLimit = "wise_droom_limit"

// WiseDroomLimit mapped from table <wise_droom_limit>
type WiseDroomLimit struct {
	DroomNo       *int64  `gorm:"column:droom_no" json:"droom_no"`
	DroomName     *string `gorm:"column:droom_name" json:"droom_name"`
	StartDate     *string `gorm:"column:start_date" json:"start_date"`
	EndDate       *string `gorm:"column:end_date" json:"end_date"`
	ComStime      *string `gorm:"column:com_stime" json:"com_stime"`
	ComEtime      *string `gorm:"column:com_etime" json:"com_etime"`
	SatStime      *string `gorm:"column:sat_stime" json:"sat_stime"`
	SatEtime      *string `gorm:"column:sat_etime" json:"sat_etime"`
	SatUseYn      *string `gorm:"column:sat_use_yn" json:"sat_use_yn"`
	ReserveDayCnt *int64  `gorm:"column:reserve_day_cnt" json:"reserve_day_cnt"`
	InputTime     *string `gorm:"column:input_time" json:"input_time"`
	Bigo          *string `gorm:"column:bigo" json:"bigo"`
	SunStime      *string `gorm:"column:sun_stime" json:"sun_stime"`
	SunEtime      *string `gorm:"column:sun_etime" json:"sun_etime"`
	SunUseYn      *string `gorm:"column:sun_use_yn" json:"sun_use_yn"`
	PatLimit      *string `gorm:"column:pat_limit" json:"pat_limit"`
}

// TableName WiseDroomLimit's table name
func (*WiseDroomLimit) TableName() string {
	return TableNameWiseDroomLimit
}
