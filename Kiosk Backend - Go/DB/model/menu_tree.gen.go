// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMENUTREE = "MENU_TREE"

// MENUTREE mapped from table <MENU_TREE>
type MENUTREE struct {
	MENUID      string `gorm:"column:MENU_ID;primaryKey" json:"MENU_ID"`
	MENUNM      string `gorm:"column:MENU_NM;not null" json:"MENU_NM"`
	DISPORDER   int64  `gorm:"column:DISP_ORDER;not null" json:"DISP_ORDER"`
	LVL         string `gorm:"column:LVL;not null" json:"LVL"`
	UPPERMENUID string `gorm:"column:UPPER_MENU_ID;not null" json:"UPPER_MENU_ID"`
	MENUURL     string `gorm:"column:MENU_URL;not null" json:"MENU_URL"`
	USEYN       string `gorm:"column:USE_YN;not null" json:"USE_YN"`
	REGDT       string `gorm:"column:REG_DT;not null" json:"REG_DT"`
	REGID       string `gorm:"column:REG_ID;not null" json:"REG_ID"`
	MODDT       string `gorm:"column:MOD_DT;not null" json:"MOD_DT"`
	MODID       string `gorm:"column:MOD_ID;not null" json:"MOD_ID"`
}

// TableName MENUTREE's table name
func (*MENUTREE) TableName() string {
	return TableNameMENUTREE
}
