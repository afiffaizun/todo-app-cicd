package model

import "time"

type Todo struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Task        string    `json:"task" gorm:"not null"`
    IsCompleted bool      `json:"is_completed" gorm:"default:false"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
