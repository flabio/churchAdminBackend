package repositories

import (
	"sync"

	"member_church.com/core/interfaces"
	"member_church.com/infrastructure/database"
	"member_church.com/infrastructure/entities"
)

func GetChurchInstance() interfaces.IChurch {
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

func (db *OpenConnection) GetChurchFindAll() ([]entities.Church, error) {
	var entities []entities.Church
	db.mux.Lock()
	err := db.connection.Order("id desc").Find(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, err
}
func (db *OpenConnection) GetChurchFindById(id uint) (entities.Church, error) {
	var entities entities.Church
	db.mux.Lock()
	err := db.connection.Where("id=?",id).Find(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, err
}
func (db *OpenConnection) GetChurchFindByName(id uint, name string) (bool, error) {
	var entities entities.Church
	db.mux.Lock()
	query := db.connection.Where("name=?", name)
	if id > 0 {
		query = query.Where("id<>?", id)
	}
	err := query.First(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return false, err
	}
	return true, err
}

func (db *OpenConnection) GetChurchFindByEmail(id uint, email string) (bool, error) {
	var entities entities.Church
	db.mux.Lock()
	query := db.connection.Where("email=?", email)
	if id > 0 {
		query = query.Where("id<>?", id)
	}
	err := query.First(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return false, err
	}
	return true, err
}


func (db *OpenConnection) CreateChurch(entities entities.Church) (entities.Church, error) {
	db.mux.Lock()
	err := db.connection.Create(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, err
}
func (db *OpenConnection) UpdateChurch(id uint, entities entities.Church) (entities.Church, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", id).Updates(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, err
}
func (db *OpenConnection) DeleteChurch(id uint) (bool, error) {
	db.mux.Lock()
	var entities entities.Church
	err := db.connection.Where("id=?", id).Delete(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
