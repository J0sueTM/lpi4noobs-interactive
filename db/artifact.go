package db

import "gorm.io/gorm"

type Artifact interface {
	FetchContent(bool) error
	Exists(*gorm.DB) bool
	Write(*gorm.DB) error
	Read(*gorm.DB) error
}
