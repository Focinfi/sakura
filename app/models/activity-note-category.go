package models

// ActivityNoteCategory for activity_note_category
type ActivityNoteCategory struct {
	ActivityNoteID     uint `gorm:"unique_index:idx_activity_note_id_activity_category_id"`
	ActivityCategoryID uint `gorm:"unique_index:idx_activity_note_id_activity_category_id"`
}
