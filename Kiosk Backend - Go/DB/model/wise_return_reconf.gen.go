// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseReturnReconf = "wise_return_reconf"

// WiseReturnReconf mapped from table <wise_return_reconf>
type WiseReturnReconf struct {
	PcNo       *int64  `gorm:"column:pc_no" json:"pc_no"`
	ReconfTime *string `gorm:"column:reconf_time" json:"reconf_time"`
}

// TableName WiseReturnReconf's table name
func (*WiseReturnReconf) TableName() string {
	return TableNameWiseReturnReconf
}
