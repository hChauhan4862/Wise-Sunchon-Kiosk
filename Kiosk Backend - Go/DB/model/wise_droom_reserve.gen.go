// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseDroomReserve = "wise_droom_reserve"

// WiseDroomReserve mapped from table <wise_droom_reserve>
type WiseDroomReserve struct {
	ReserveNo   *string `gorm:"column:reserve_no" json:"reserve_no"`
	DroomNo     *int64  `gorm:"column:droom_no" json:"droom_no"`
	DroomName   *string `gorm:"column:droom_name" json:"droom_name"`
	InputTime   *string `gorm:"column:input_time" json:"input_time"`
	ReserveDate *string `gorm:"column:reserve_date" json:"reserve_date"`
	StartTime   *string `gorm:"column:start_time" json:"start_time"`
	EndTime     *string `gorm:"column:end_time" json:"end_time"`
	UserID      *string `gorm:"column:user_id" json:"user_id"`
	UserName    *string `gorm:"column:user_name" json:"user_name"`
	UserComment *string `gorm:"column:user_comment" json:"user_comment"`
	UseMan      *int64  `gorm:"column:use_man" json:"use_man"`
	CouserIds   *string `gorm:"column:couser_ids" json:"couser_ids"`
	ReserveStat *int64  `gorm:"column:reserve_stat" json:"reserve_stat"`
	CanBigo     *string `gorm:"column:can_bigo" json:"can_bigo"`
	Bigo        *string `gorm:"column:bigo" json:"bigo"`
}

// TableName WiseDroomReserve's table name
func (*WiseDroomReserve) TableName() string {
	return TableNameWiseDroomReserve
}
