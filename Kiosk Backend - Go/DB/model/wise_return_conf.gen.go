// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseReturnConf = "wise_return_conf"

// WiseReturnConf mapped from table <wise_return_conf>
type WiseReturnConf struct {
	PcNo       *int64  `gorm:"column:pc_no" json:"pc_no"`
	Gubun      *string `gorm:"column:gubun" json:"gubun"`
	GubunValue *string `gorm:"column:gubun_value" json:"gubun_value"`
}

// TableName WiseReturnConf's table name
func (*WiseReturnConf) TableName() string {
	return TableNameWiseReturnConf
}
