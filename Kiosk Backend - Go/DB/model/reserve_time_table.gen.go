// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_reserve_time_table = "reserve_time_table"

// HCV_reserve_time_table mapped from table <reserve_time_table>
type HCV_reserve_time_table struct {
	SroomNo     *int64  `gorm:"column:sroom_no" json:"sroom_no"`
	SroomName   *string `gorm:"column:sroom_name" json:"sroom_name"`
	ReserveDate *string `gorm:"column:reserve_date" json:"reserve_date"`
	Time0       *int64  `gorm:"column:time0" json:"time0"`
	Time1       *int64  `gorm:"column:time1" json:"time1"`
	Time2       *int64  `gorm:"column:time2" json:"time2"`
	Time3       *int64  `gorm:"column:time3" json:"time3"`
	Time4       *int64  `gorm:"column:time4" json:"time4"`
	Time5       *int64  `gorm:"column:time5" json:"time5"`
	Time6       *int64  `gorm:"column:time6" json:"time6"`
	Time7       *int64  `gorm:"column:time7" json:"time7"`
	Time8       *int64  `gorm:"column:time8" json:"time8"`
	Time9       *int64  `gorm:"column:time9" json:"time9"`
	Time10      *int64  `gorm:"column:time10" json:"time10"`
	Time11      *int64  `gorm:"column:time11" json:"time11"`
	Time12      *int64  `gorm:"column:time12" json:"time12"`
	Time13      *int64  `gorm:"column:time13" json:"time13"`
	Time14      *int64  `gorm:"column:time14" json:"time14"`
	Time15      *int64  `gorm:"column:time15" json:"time15"`
	Time16      *int64  `gorm:"column:time16" json:"time16"`
	Time17      *int64  `gorm:"column:time17" json:"time17"`
	Time18      *int64  `gorm:"column:time18" json:"time18"`
	Time19      *int64  `gorm:"column:time19" json:"time19"`
	Time20      *int64  `gorm:"column:time20" json:"time20"`
	Time21      *int64  `gorm:"column:time21" json:"time21"`
	Time22      *int64  `gorm:"column:time22" json:"time22"`
	Time23      *int64  `gorm:"column:time23" json:"time23"`
	Time24      *int64  `gorm:"column:time24" json:"time24"`
	Time25      *int64  `gorm:"column:time25" json:"time25"`
	Time26      *int64  `gorm:"column:time26" json:"time26"`
	Time27      *int64  `gorm:"column:time27" json:"time27"`
	Time28      *int64  `gorm:"column:time28" json:"time28"`
	Time29      *int64  `gorm:"column:time29" json:"time29"`
	Time30      *int64  `gorm:"column:time30" json:"time30"`
	Time31      *int64  `gorm:"column:time31" json:"time31"`
	Time32      *int64  `gorm:"column:time32" json:"time32"`
	Time33      *int64  `gorm:"column:time33" json:"time33"`
}

// TableName HCV_reserve_time_table's table name
func (*HCV_reserve_time_table) TableName() string {
	return TableNameHCV_reserve_time_table
}
