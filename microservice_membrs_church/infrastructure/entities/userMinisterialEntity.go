package entities

type UserMinisterial struct {
	UserId        uint        `gorm:"null" json:"user_id"`
	User          User        `gorm:"foreignKey:UserId  constraint:onUpdate:CASCADE,onDelete" json:"user"`
	MinisterialId uint        `gorm:"null" json:"ministerial_id"`
	Ministerial   Ministerial `gorm:"foreignKey:MinisterialId  constraint:onUpdate:CASCADE,onDelete" json:"ministerial"`
}
