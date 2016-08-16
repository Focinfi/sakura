package models

import "github.com/jinzhu/gorm"

// Feeling for user feeling
type Feeling struct {
	gorm.Model
	Name uint
	// 喜
	Xi uint
	// 怒
	Nu uint
	// 思
	Si uint
	// 悲
	Bei uint
	// 恐
	Kong uint
}
