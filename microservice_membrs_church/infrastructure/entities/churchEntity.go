package entities

import "time"

type Church struct {
	Id        uint       `gorm:"primary_key:auto_increment"`
	Name      string     `gorm:"type:varchar(250);not null"`
	Phone     string     `gorm:"type:varchar(250);not null"`
	Email     string     `gorm:"type:varchar(250);not null"`
	Address   string     `gorm:"type:varchar(100);not null"`
	Active    bool       `gorm:"type:boolean"`
	CreatedAt time.Time  `gorm:"<-:created_at" `
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" `
}
