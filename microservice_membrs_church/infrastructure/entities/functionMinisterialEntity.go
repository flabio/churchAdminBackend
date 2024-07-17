package entities

import "time"

type FunctionMinisterial struct {
	Id        uint       `gorm:"primary_key:auto_increment"`
	Name      string     `gorm:"type:varchar(100);not null"`
	UserId    uint       `gorm:"null" json:"user_id"`
	User      User       `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Active    bool       `gorm:"type:boolean"`
	CreatedAt time.Time  `gorm:"<-:created_at" `
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" `
}
