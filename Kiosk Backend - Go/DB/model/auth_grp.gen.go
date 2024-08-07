// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAUTHGRP = "AUTH_GRP"

// AUTHGRP mapped from table <AUTH_GRP>
type AUTHGRP struct {
	AUTHGRPCD  int64   `gorm:"column:AUTH_GRP_CD;primaryKey" json:"AUTH_GRP_CD"`
	AUTHGRPNM  string  `gorm:"column:AUTH_GRP_NM;not null" json:"AUTH_GRP_NM"`
	APPGUBUN   string  `gorm:"column:APP_GUBUN;not null" json:"APP_GUBUN"`
	INOUTGUBUN *string `gorm:"column:IN_OUT_GUBUN" json:"IN_OUT_GUBUN"`
	AUTHTYPE   *string `gorm:"column:AUTH_TYPE;not null;default:R" json:"AUTH_TYPE"`
	USEYN      *string `gorm:"column:USE_YN;not null;default:Y" json:"USE_YN"`
	REGDT      string  `gorm:"column:REG_DT;not null" json:"REG_DT"`
	REGID      string  `gorm:"column:REG_ID;not null" json:"REG_ID"`
	MODDT      string  `gorm:"column:MOD_DT;not null" json:"MOD_DT"`
	MODID      string  `gorm:"column:MOD_ID;not null" json:"MOD_ID"`
}

// TableName AUTHGRP's table name
func (*AUTHGRP) TableName() string {
	return TableNameAUTHGRP
}
