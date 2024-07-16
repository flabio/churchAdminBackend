package repositories

import (
	"sync"

	"member_church.com/core/interfaces"
	"member_church.com/infrastructure/database"
	"member_church.com/infrastructure/entities"
)

func GetMinisterialInstance() interfaces.IMinisterial {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}
func (db *OpenConnection) GetMinisterialFindAll() ([]entities.Ministerial, error) {
	var ministerial []entities.Ministerial
	db.mux.Lock()
	db.connection.Order("id desc").Find(&ministerial)
	defer db.mux.Unlock()
	defer database.Closedb()
	return ministerial, nil
}
func (db *OpenConnection) GetMinisterialFindById(id uint) (entities.Ministerial, error) {
	var ministerial entities.Ministerial
	db.mux.Lock()
	err := db.connection.Where("id=?", id).First(&ministerial).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return ministerial, err
}
func (db *OpenConnection) CreateMinisterial(ministerial entities.Ministerial) (entities.Ministerial, error) {
	db.mux.Lock()
	err := db.connection.Create(&ministerial).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return ministerial, err
}
func (db *OpenConnection) UpdateMinisterial(id uint, ministerial entities.Ministerial) (entities.Ministerial, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", id).Updates(&ministerial).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return ministerial, err
}
func (db *OpenConnection) DeleteMinisterial(id uint) (bool, error) {
	db.mux.Lock()
	var min entities.Ministerial
	err := db.connection.Where("id=?", id).Delete(&min).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}

func (db *OpenConnection) GetMinisterialFindByName(id uint, name string) (bool, error) {
	db.mux.Lock()
	var min entities.Ministerial

	query := db.connection.Where("name=?", name)
	if id > 0 {
		query = query.Where("id<>?", id).First(&min)
	} else {
		query = query.First(&min)
	}
	err := query.Error
	defer db.mux.Unlock()
	defer database.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}
