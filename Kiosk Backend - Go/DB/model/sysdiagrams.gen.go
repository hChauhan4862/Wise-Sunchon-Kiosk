// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysdiagram = "sysdiagrams"

// Sysdiagram mapped from table <sysdiagrams>
type Sysdiagram struct {
	Name        string   `gorm:"column:name;not null" json:"name"`
	PrincipalID int64    `gorm:"column:principal_id;not null" json:"principal_id"`
	DiagramID   int64    `gorm:"column:diagram_id;primaryKey" json:"diagram_id"`
	Version     *int64   `gorm:"column:version" json:"version"`
	Definition  *[]uint8 `gorm:"column:definition" json:"definition"`
}

// TableName Sysdiagram's table name
func (*Sysdiagram) TableName() string {
	return TableNameSysdiagram
}
