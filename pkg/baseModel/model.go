package baseModel

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time      `gorm:"type:timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type UidModel struct {
	Uid string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Model
}
type IdModel struct {
	ID uint `gorm:"primarykey"`
	Model
}
