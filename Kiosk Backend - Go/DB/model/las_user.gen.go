// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLASUSER = "LAS_USER"

// LASUSER mapped from table <LAS_USER>
type LASUSER struct {
	PID      *string `gorm:"column:PID" json:"PID"`
	RID      string  `gorm:"column:RID;not null" json:"RID"`
	NAME     *string `gorm:"column:NAME" json:"NAME"`
	PATCODE  *string `gorm:"column:PATCODE" json:"PATCODE"`
	PATNAME  *string `gorm:"column:PATNAME" json:"PATNAME"`
	DEPTCODE *string `gorm:"column:DEPTCODE" json:"DEPTCODE"`
	DEPTNAME *string `gorm:"column:DEPTNAME" json:"DEPTNAME"`
}

// TableName LASUSER's table name
func (*LASUSER) TableName() string {
	return TableNameLASUSER
}