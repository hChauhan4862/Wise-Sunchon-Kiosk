// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVariSeatFix = "vari_seat_fix"

// VariSeatFix mapped from table <vari_seat_fix>
type VariSeatFix struct {
	ExamGubun       string  `gorm:"column:exam_gubun;not null" json:"exam_gubun"`
	RoomNo          int64   `gorm:"column:room_no;not null" json:"room_no"`
	StartDate       *string `gorm:"column:start_date" json:"start_date"`
	StartTime       *string `gorm:"column:start_time" json:"start_time"`
	EndDate         *string `gorm:"column:end_date" json:"end_date"`
	EndTime         *string `gorm:"column:end_time" json:"end_time"`
	SeatFix         *string `gorm:"column:seat_fix" json:"seat_fix"`
	OriSeatFix      *string `gorm:"column:ori_seat_fix" json:"ori_seat_fix"`
	RegDate         string  `gorm:"column:reg_date;not null" json:"reg_date"`
	SeatRoomSeatFix *string `gorm:"column:seat_room_seat_fix" json:"seat_room_seat_fix"`
}

// TableName VariSeatFix's table name
func (*VariSeatFix) TableName() string {
	return TableNameVariSeatFix
}
