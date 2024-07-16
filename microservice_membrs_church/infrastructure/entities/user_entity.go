package entities

import "time"

type User struct {
	Id                 uint      `gorm:"primary key:auto_increment"`
	FirstName          string    `gorm:"type:varchar(255);not null"`
	LastName           string    `gorm:"type:varchar(255);not null"`
	Identication       string    `gorm:"type:varchar(255);not null"`
	TypeIdentification string    `gorm:"type:varchar(6);"`
	Birthdate          string    `gorm:"type:varchar(10)"`
	PlaceOfBirth       string    `gorm:"type:text"`
	Address            string    `gorm:"type:varchar(100)"`
	Phone              string    `gorm:"type:varchar(20)"`
	Cellphone          string    `gorm:"type:varchar(20)"`
	Email              string    `gorm:"type:text"`
	Neighborhood       string    `gorm:"type:text"`
	Locality           string    `gorm:"type:text"`
	Socioeconomic      uint      `gorm:"type:integer"`
	Rh                 string    `gorm:"type:varchar(6)"`
	Profession         string    `gorm:"type:text"`
	Company            string    `gorm:"type:text"`
	Position           string    `gorm:"type:text"`
	CivilStatus        string    `gorm:"type:varchar(50)"`
	LeadersName        string    `gorm:"type:text"`
	ConversionDate     string    `gorm:"type:varchar(10)"`
	WhoInvitedHim      string    `gorm:"type:text"`
	ChurchAttended     bool      `gorm:"type:boolean"`
	ChurchName         string    `gorm:"type:varchar(250)"`
	MeetingDate        string    `gorm:"type:varchar(10)"`
	BaptismDate        string    `gorm:"type:varchar(10)"`
	YourSpouseName     string    `gorm:"type:varchar(250)"`
	SpouseMemberChurch bool      `gorm:"type:boolean"`
	DateMarriage       string    `gorm:"type:varchar(10)"`
	ParentsName        string    `gorm:"type:varchar(250)"`
	IsMemberParent     bool      `gorm:"type:boolean"`
	ChildrenaName      string    `gorm:"type:varchar(250)"`
	Allergy            bool      `gorm:"type:boolean"`
	Authorize          bool      `gorm:"type:boolean"`
	Username           string    `gorm:"type:string"`
	Password           string    `gorm:"type:string"`
	CreatedAt          time.Time `gorm:"<-:created_at" `
	ChurchId           uint64    `gorm:"null" json:"church_id"`
	Church             Church    `gorm:"foreignKey:ChurchId  constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"church"`
	RolId              uint      `gorm:"null" json:"rol_id"`
	Rol                Rol       `gorm:"foreignKey:RolId  constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"rol"`
	TeamPescaId        uint      `gorm:"null" json:"team_pesca_id"`
	TeamPesca          TeamPesca `gorm:"foreignKey:TeamPescaId  constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"team_pesca"`
}
