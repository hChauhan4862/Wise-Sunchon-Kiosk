// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseDroomRessub = "wise_droom_ressub"

// WiseDroomRessub mapped from table <wise_droom_ressub>
type WiseDroomRessub struct {
	ReserveNo   *string `gorm:"column:reserve_no" json:"reserve_no"`
	UserID      *string `gorm:"column:user_id" json:"user_id"`
	UserName    *string `gorm:"column:user_name" json:"user_name"`
	CheckYn     *string `gorm:"column:check_yn" json:"check_yn"`
	CheckTime   *string `gorm:"column:check_time" json:"check_time"`
	ReserveStat *int64  `gorm:"column:reserve_stat" json:"reserve_stat"`
	Bigo        *string `gorm:"column:bigo" json:"bigo"`
}

// TableName WiseDroomRessub's table name
func (*WiseDroomRessub) TableName() string {
	return TableNameWiseDroomRessub
}
