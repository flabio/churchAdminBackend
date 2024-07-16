package interfaces

import "member_church.com/infrastructure/entities"

type IMinisterial interface {
	GetMinisterialFindAll() ([]entities.Ministerial, error)
	GetMinisterialFindById(id uint) (entities.Ministerial, error)
	GetMinisterialFindByName(id uint, name string)(bool,error)
	CreateMinisterial(ministerial entities.Ministerial) (entities.Ministerial, error)
	UpdateMinisterial(id uint, ministerial entities.Ministerial) (entities.Ministerial, error)
	DeleteMinisterial(id uint) (bool, error)
}
