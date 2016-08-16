package models

// UserFeeling for user_feeling
type UserFeeling struct {
	UserID    uint `gorm:"unique_index:idx_user_id_feeling_id"`
	FeelingID uint `gorm:"unique_index:idx_user_id_feeling_id"`
}
