// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatUserdatum = "seat_userdata"

// SeatUserdatum mapped from table <seat_userdata>
type SeatUserdatum struct {
	Pid          *string `gorm:"column:pid" json:"pid"`
	JuminBunho   *string `gorm:"column:jumin_bunho" json:"jumin_bunho"`
	Password     *string `gorm:"column:password" json:"password"`
	PatType      *string `gorm:"column:pat_type" json:"pat_type"`
	PatTypeDesc  *string `gorm:"column:pat_type_desc" json:"pat_type_desc"`
	Name         *string `gorm:"column:name" json:"name"`
	DeptCode     *string `gorm:"column:dept_code" json:"dept_code"`
	DeptName     *string `gorm:"column:dept_name" json:"dept_name"`
	CardIssueCnt *int64  `gorm:"column:card_issue_cnt" json:"card_issue_cnt"`
	DateExprd    *string `gorm:"column:date_exprd" json:"date_exprd"`
	Email        *string `gorm:"column:email" json:"email"`
	Phone1       *string `gorm:"column:phone1" json:"phone1"`
	Status       *string `gorm:"column:status" json:"status"`
	StatusName   *string `gorm:"column:status_name" json:"status_name"`
	PhotoURL     *string `gorm:"column:photo_url" json:"photo_url"`
	SchoolNo     *string `gorm:"column:school_no" json:"school_no"`
}

// TableName SeatUserdatum's table name
func (*SeatUserdatum) TableName() string {
	return TableNameSeatUserdatum
}
