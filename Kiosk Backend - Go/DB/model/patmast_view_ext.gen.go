// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_patmast_view_ext = "patmast_view_ext"

// HCV_patmast_view_ext mapped from table <patmast_view_ext>
type HCV_patmast_view_ext struct {
	AltPid       *string `gorm:"column:alt_pid" json:"alt_pid"`
	PatType      *string `gorm:"column:pat_type" json:"pat_type"`
	DeptCode     *string `gorm:"column:dept_code" json:"dept_code"`
	CardIssueCnt int64   `gorm:"column:card_issue_cnt;not null" json:"card_issue_cnt"`
	DateExprd    *string `gorm:"column:date_exprd" json:"date_exprd"`
	Username     *string `gorm:"column:username" json:"username"`
	Phone3       *string `gorm:"column:phone3" json:"phone3"`
}

// TableName HCV_patmast_view_ext's table name
func (*HCV_patmast_view_ext) TableName() string {
	return TableNameHCV_patmast_view_ext
}
