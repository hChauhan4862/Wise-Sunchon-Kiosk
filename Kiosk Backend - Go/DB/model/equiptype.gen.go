// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEQUIPTYPE = "EQUIPTYPE"

// EQUIPTYPE mapped from table <EQUIPTYPE>
type EQUIPTYPE struct {
	EQTYPE   int64   `gorm:"column:EQTYPE;primaryKey" json:"EQTYPE"`
	NAME     *string `gorm:"column:NAME" json:"NAME"`
	ENNAME   *string `gorm:"column:EN_NAME" json:"EN_NAME"`
	DETAIL   *string `gorm:"column:DETAIL" json:"DETAIL"`
	ISDELETE *int64  `gorm:"column:ISDELETE" json:"ISDELETE"`
}

// TableName EQUIPTYPE's table name
func (*EQUIPTYPE) TableName() string {
	return TableNameEQUIPTYPE
}
