// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatSectorView = "seat_sector_view"

// SeatSectorView mapped from table <seat_sector_view>
type SeatSectorView struct {
	VIEWSECTORNO int64   `gorm:"column:VIEWSECTORNO;primaryKey" json:"VIEWSECTORNO"`
	NAME         *string `gorm:"column:NAME" json:"NAME"`
	ENNAME       *string `gorm:"column:EN_NAME" json:"EN_NAME"`
	ROOMNO       *int64  `gorm:"column:ROOMNO" json:"ROOMNO"`
	IMAGE        *string `gorm:"column:IMAGE" json:"IMAGE"`
	LINKPAGE     *string `gorm:"column:LINKPAGE" json:"LINKPAGE"`
	MAPPOINT     *string `gorm:"column:MAPPOINT" json:"MAPPOINT"`
	DEFSECTORNO  *int64  `gorm:"column:DEFSECTORNO" json:"DEFSECTORNO"`
	MAPLABEL     *string `gorm:"column:MAPLABEL" json:"MAPLABEL"`
	ENMAPLABEL   *string `gorm:"column:EN_MAPLABEL" json:"EN_MAPLABEL"`
	MAPPOINT2    *string `gorm:"column:MAPPOINT2" json:"MAPPOINT2"`
	MAPLABEL2    *string `gorm:"column:MAPLABEL2" json:"MAPLABEL2"`
	ENMAPLABEL2  *string `gorm:"column:EN_MAPLABEL2" json:"EN_MAPLABEL2"`
}

// TableName SeatSectorView's table name
func (*SeatSectorView) TableName() string {
	return TableNameSeatSectorView
}
