package entities

import "time"

type TeamPesca struct {
	Id        uint       `gorm:"primary_key:auto_increment"`
	Name      string     `gorm:"type:text"`
	Active    bool       `gorm:"type:boolean"`
	CreatedAt time.Time  `gorm:"<-:created_at" `
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" `
}
