// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseSroomReserve = "wise_sroom_reserve"

// WiseSroomReserve mapped from table <wise_sroom_reserve>
type WiseSroomReserve struct {
	ReserveNo   *string `gorm:"column:reserve_no" json:"reserve_no"`
	SroomNo     *int64  `gorm:"column:sroom_no" json:"sroom_no"`
	SroomName   *string `gorm:"column:sroom_name" json:"sroom_name"`
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
	ReserveSeq  int64   `gorm:"column:reserve_seq;not null" json:"reserve_seq"`
}

// TableName WiseSroomReserve's table name
func (*WiseSroomReserve) TableName() string {
	return TableNameWiseSroomReserve
}
