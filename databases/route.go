package databases

type Route struct {
	ID     int    `gorm:"primary_key"`
	UserID int    `gorm:"column:user_id"`
	Title  string `gorm:"column:title"`
}
