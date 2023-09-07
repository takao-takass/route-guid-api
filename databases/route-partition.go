package databases

import "time"

type RoutePartition struct {
	ID      int        `gorm:"column:id;primaryKey;autoIncrement"`
	RouteID int        `gorm:"column:route_id"`
	Title   string     `gorm:"column:title"`
	Date    *time.Time `gorm:"column:date"`
}
