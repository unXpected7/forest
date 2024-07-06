// overrides/swagger_overrides.go
package overrides

import (
	"time"
)

// This file is used by swag to recognize gorm.Model

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
