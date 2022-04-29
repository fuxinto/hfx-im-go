package baseModel

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type UidModel struct {
	Uid string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Model
}
