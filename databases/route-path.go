package databases

import "time"

type RoutePath struct {
	ID               int        `gorm:"column:id;primaryKey;autoIncrement"`
	RoutePartitionID int        `gorm:"column:route_partition_id"`
	Order            int        `gorm:"column:order"`
	PathType         int        `gorm:"column:path_type"`
	PathName         string     `gorm:"column:path_name"`
	Clock            *time.Time `gorm:"column:clock"`
}
