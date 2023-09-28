package db

import (
	"encoding/json"
	"os"

	"gorm.io/gorm"
)

type Artifact interface {
	FetchContent(bool) error
	Exists(*gorm.DB) bool
	Write(*gorm.DB) error
	Read(*gorm.DB) error
}

func PopulateFromResource(artifact interface{}, resource string) error {
	data, err := os.ReadFile(resource)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, artifact)
	if err != nil {
		return err
	}

	return nil
}
