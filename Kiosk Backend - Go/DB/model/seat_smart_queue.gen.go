// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatSmartQueue = "seat_smart_queue"

// SeatSmartQueue mapped from table <seat_smart_queue>
type SeatSmartQueue struct {
	InsDate   *string `gorm:"column:ins_date" json:"ins_date"`
	RoomNo    *int64  `gorm:"column:room_no" json:"room_no"`
	SeatNo    *int64  `gorm:"column:seat_no" json:"seat_no"`
	CmdType   *string `gorm:"column:cmd_type" json:"cmd_type"`
	SetCmd    *string `gorm:"column:set_cmd" json:"set_cmd"`
	SetDate   *string `gorm:"column:set_date" json:"set_date"`
	CheckDate *string `gorm:"column:check_date" json:"check_date"`
}

// TableName SeatSmartQueue's table name
func (*SeatSmartQueue) TableName() string {
	return TableNameSeatSmartQueue
}
