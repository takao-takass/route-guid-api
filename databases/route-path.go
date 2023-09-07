package databases

import "time"

type RoutePath struct {
	ID               int `gorm:"primaryKey"`
	RoutePartitionID int
	Clock            time.Time
	IsStation        bool
	IsStay           bool
	LocationName     *string
	IsTrain          bool
	TrainName        *string
	TrainType        *string
	TrainFor         *string
}
