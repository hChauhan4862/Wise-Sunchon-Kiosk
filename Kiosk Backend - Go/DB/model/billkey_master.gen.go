// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBILLKEYMASTER = "BILLKEY_MASTER"

// BILLKEYMASTER mapped from table <BILLKEY_MASTER>
type BILLKEYMASTER struct {
	USERID  *string `gorm:"column:USER_ID" json:"USER_ID"`
	BILLKEY *string `gorm:"column:BILLKEY" json:"BILLKEY"`
	REGDT   *string `gorm:"column:REG_DT" json:"REG_DT"`
}

// TableName BILLKEYMASTER's table name
func (*BILLKEYMASTER) TableName() string {
	return TableNameBILLKEYMASTER
}