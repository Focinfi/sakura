package models

// ActivityNoteTag for activity_note_tag
type ActivityNoteTag struct {
	ActivityNoteID uint `gorm:"unique_index:idx_activity_note_id_tag_id"`
	TagID          uint `gorm:"unique_index:idx_activity_note_id_tag_id"`
}
