// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAppVersion = "app_version"

// AppVersion mapped from table <app_version>
type AppVersion struct {
	AndVersion string `gorm:"column:and_version;not null" json:"and_version"`
	AndRegDt   string `gorm:"column:and_reg_dt;not null" json:"and_reg_dt"`
	IosVersion string `gorm:"column:ios_version;not null" json:"ios_version"`
	IosRegDt   string `gorm:"column:ios_reg_dt;not null" json:"ios_reg_dt"`
}

// TableName AppVersion's table name
func (*AppVersion) TableName() string {
	return TableNameAppVersion
}
