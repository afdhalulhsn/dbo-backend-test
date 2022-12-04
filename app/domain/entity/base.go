package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID string `gorm:"primaryKey;size:100" json:"id"`
}

func (a *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	a.ID = uuid.String()
	return nil
}
