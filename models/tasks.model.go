package models

import "gorm.io/gorm"

type TaskModel struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"type:varchar(4);not null"`
	Description string `gorm:"type:varchar(10);not null"`
	IsCompleted bool   `gorm:"type:boolean;default:false"`
}

func (t *TaskModel) TableName() string {
	return "tasks"
}
