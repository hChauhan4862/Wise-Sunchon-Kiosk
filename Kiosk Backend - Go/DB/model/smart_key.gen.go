// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmartKey = "smart_key"

// SmartKey mapped from table <smart_key>
type SmartKey struct {
	Gubun    string  `gorm:"column:gubun;not null" json:"gubun"`
	RoomNo   int64   `gorm:"column:room_no;not null" json:"room_no"`
	RoomName *string `gorm:"column:room_name" json:"room_name"`
	SeatNo   int64   `gorm:"column:seat_no;not null" json:"seat_no"`
	CurrKey  *string `gorm:"column:curr_key" json:"curr_key"`
	PrevKey  *string `gorm:"column:prev_key" json:"prev_key"`
	RegDt    string  `gorm:"column:reg_dt;not null" json:"reg_dt"`
	UpdDt    *string `gorm:"column:upd_dt" json:"upd_dt"`
}

// TableName SmartKey's table name
func (*SmartKey) TableName() string {
	return TableNameSmartKey
}