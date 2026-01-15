package model

import "time"

type Todo struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Task        string    `json:"task" gorm:"not null"`
    IsCompleted bool      `json:"is_completed" gorm:"default:false"`
    UserID      uint      `json:"user_id" gorm:"not null;index"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
