package interfaces

import "member_church.com/infrastructure/entities"

type IChurch interface {
	GetChurchFindAll() ([]entities.Church, error)
	GetChurchFindById(id uint) (entities.Church, error)
	GetChurchFindByName(id uint, name string) (bool, error)
	GetChurchFindByEmail(id uint, email string) (bool, error)
	CreateChurch(church entities.Church) (entities.Church, error)
	UpdateChurch(id uint, church entities.Church) (entities.Church, error)
	DeleteChurch(id uint) (bool, error)
}
