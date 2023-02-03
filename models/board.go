package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Text string
}
